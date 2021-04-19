package server

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/Saser/pdp/aip/pagetoken"
	"github.com/Saser/pdp/aip/resourcename"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	taskspb "github.com/Saser/pdp/tasks2/tasks_go_proto"
)

const maxPageSize = 100

type Server struct {
	taskspb.UnimplementedTasksServer

	mu           sync.Mutex
	tasks        []*taskspb.Task
	taskIndices  map[string]int // task name -> index into `tasks`
	events       []*taskspb.Event
	eventIndices map[string]int      // event name -> index into `events`
	taskEvents   map[string][]string // task name -> event names
}

func New() *Server {
	return &Server{
		taskIndices:  make(map[string]int),
		eventIndices: make(map[string]int),
		taskEvents:   make(map[string][]string),
	}
}

func (s *Server) GetTask(ctx context.Context, req *taskspb.GetTaskRequest) (*taskspb.Task, error) {
	name := req.GetName()
	if name == "" {
		return nil, status.Error(codes.InvalidArgument, "empty name")
	}
	if !taskPattern.Matches(name) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid name %q does not have format %q", name, taskPattern)
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	idx, ok := s.taskIndices[name]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "task %q not found", name)
	}
	return s.tasks[idx], nil
}

func (s *Server) ListTasks(ctx context.Context, req *taskspb.ListTasksRequest) (*taskspb.ListTasksResponse, error) {
	pageSize := req.GetPageSize()
	if pageSize < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "negative page size %d", pageSize)
	}
	if pageSize == 0 || pageSize > maxPageSize {
		pageSize = maxPageSize
	}
	pt, err := pagetoken.Parse(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid page token %q: %v", req.GetPageToken(), err)
	}

	start := int(pt.Offset())
	next := pt.Next(pageSize)
	end := int(next.Offset())

	s.mu.Lock()
	defer s.mu.Unlock()
	res := &taskspb.ListTasksResponse{}
	if end >= len(s.tasks) {
		res.Tasks = s.tasks[start:]
		res.NextPageToken = ""
	} else {
		res.Tasks = s.tasks[start:end]
		res.NextPageToken = next.String()
	}
	return res, nil
}

func (s *Server) CreateTask(ctx context.Context, req *taskspb.CreateTaskRequest) (*taskspb.Task, error) {
	task := req.GetTask()
	if task.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "empty title")
	}
	if task.GetCompleted() == true {
		return nil, status.Error(codes.InvalidArgument, `"completed" is true; must be false when creating task`)
	}
	if len(task.GetDependencies()) > 0 {
		return nil, status.Error(codes.InvalidArgument, `"dependencies" is non-empty; must be empty when creating task`)
	}
	if len(task.GetLabels()) > 0 {
		return nil, status.Error(codes.InvalidArgument, `"labels" is non-empty; must be empty when creating task`)
	}
	if task.GetDeferral() != nil {
		return nil, status.Error(codes.InvalidArgument, `"deferral" is non-empty; must be empty when creating task`)
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	idx := len(s.tasks)
	v := resourcename.Values{
		"task": fmt.Sprint(idx + 1),
	}
	name, err := taskPattern.Render(v)
	if err != nil {
		log.Printf("CreateTask failed to render task name: %v", err)
		return nil, status.Error(codes.Internal, "internal error")
	}
	task.Name = name
	s.tasks = append(s.tasks, task)
	s.taskIndices[task.Name] = idx
	return task, nil
}

func (s *Server) AddDependency(ctx context.Context, req *taskspb.AddDependencyRequest) (*taskspb.Task, error) {
	taskName := req.GetTask()
	if taskName == "" {
		return nil, status.Error(codes.InvalidArgument, "empty task")
	}
	if !taskPattern.Matches(taskName) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid task name %q; want format %q", taskName, taskPattern)
	}
	dependencyName := req.GetDependency()
	if dependencyName == "" {
		return nil, status.Error(codes.InvalidArgument, "empty dependency")
	}
	if !taskPattern.Matches(dependencyName) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid dependency name %q; want format %q", dependencyName, taskPattern)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	taskIdx, ok := s.taskIndices[taskName]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "task %q not found", taskName)
	}
	task := s.tasks[taskIdx]

	if _, ok := s.taskIndices[dependencyName]; !ok {
		return nil, status.Errorf(codes.NotFound, "dependency %q not found", dependencyName)
	}

	for _, existing := range task.GetDependencies() {
		if existing == dependencyName {
			return nil, status.Errorf(codes.AlreadyExists, "%q already depends on %q", taskName, dependencyName)
		}
	}

	task.Dependencies = append(task.GetDependencies(), dependencyName)

	for _, parent := range []string{
		taskName,
		dependencyName,
	} {
		event := &taskspb.Event{
			CreateTime: timestamppb.Now(),
			Comment:    req.GetComment(),
			Kind: &taskspb.Event_AddDependency_{AddDependency: &taskspb.Event_AddDependency{
				Task:       taskName,
				Dependency: dependencyName,
			}},
		}
		if _, err := s.createEvent(ctx, parent, event); err != nil {
			return nil, err
		}
	}

	return task, nil
}

func (s *Server) ListEvents(ctx context.Context, req *taskspb.ListEventsRequest) (*taskspb.ListEventsResponse, error) {
	parent := req.GetParent()
	if parent == "" {
		return nil, status.Error(codes.InvalidArgument, "empty parent")
	}
	if !taskPattern.Matches(parent) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid parent %q; want format %q", parent, taskPattern)
	}

	pageSize := req.GetPageSize()
	if pageSize < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "negative page size %d", pageSize)
	}
	if pageSize == 0 || pageSize > maxPageSize {
		pageSize = maxPageSize
	}

	pt, err := pagetoken.Parse(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid page token %q: %v", req.GetPageToken(), err)
	}
	start := int(pt.Offset())
	next := pt.Next(pageSize)
	end := int(next.Offset())

	s.mu.Lock()
	defer s.mu.Unlock()

	eventNames, ok := s.taskEvents[parent]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "no events found for parent %q", parent)
	}

	// eventNames is the collection we are paginating over.
	var lastPage bool
	if end > len(eventNames) {
		eventNames = eventNames[start:]
		lastPage = true
	} else {
		eventNames = eventNames[start:end]
		lastPage = false
	}

	var indices []int
	for _, event := range eventNames {
		index, ok := s.eventIndices[event]
		if !ok {
			log.Printf("no index for event %q (parent %q)", event, parent)
			return nil, status.Error(codes.Internal, "internal error")
		}
		indices = append(indices, index)
	}

	var events []*taskspb.Event
	for _, index := range indices {
		events = append(events, s.events[index])
	}

	res := &taskspb.ListEventsResponse{
		Events: events,
	}
	if !lastPage {
		res.NextPageToken = next.String()
	}
	return res, nil
}

// createEvent creates the given event under the given parent. The event's name and parent fields
// will be overwritten, and the updated event is returned. This method takes care of the internal
// bookkeeping. Any error returned is created by the status package, and can be returned directly
// from an RPC method.
//
// This method is not thread-safe. A caller must hold the server's mutex before calling this method.
func (s *Server) createEvent(ctx context.Context, parent string, event *taskspb.Event) (*taskspb.Event, error) {
	v, err := eventPattern.Match(parent)
	if err != nil {
		log.Printf("parent task %q didn't match event pattern %q", parent, eventPattern)
		return nil, status.Error(codes.Internal, "internal error")
	}
	taskEvents := s.taskEvents[parent]
	defer func() {
		s.taskEvents[parent] = taskEvents
	}()
	id := strconv.Itoa(len(taskEvents) + 1)
	v["event"] = id
	name, err := eventPattern.Render(v)
	if err != nil {
		log.Printf("failed to render name for new event: %v", err)
		return nil, status.Error(codes.Internal, "internal error")
	}
	taskEvents = append(taskEvents, name)

	event.Name = name
	event.Parent = parent

	index := len(s.events)
	s.eventIndices[name] = index
	s.events = append(s.events, event)
	return event, nil
}
