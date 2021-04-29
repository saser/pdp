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

var internalError = status.Error(codes.Internal, "internal error")

type Server struct {
	taskspb.UnimplementedTasksServer

	// mu protects all fields in this struct.
	mu sync.Mutex

	tasks       []*taskspb.Task
	taskIndices map[string]int // task name -> index into `tasks`

	events       []*taskspb.Event
	eventIndices map[string]int      // event name -> index into `events`
	taskEvents   map[string][]string // task name -> event names

	labels       []*taskspb.Label
	labelIndices map[string]int // label name -> index into `labels`
}

func New() *Server {
	return &Server{
		taskIndices:  make(map[string]int),
		eventIndices: make(map[string]int),
		taskEvents:   make(map[string][]string),
		labelIndices: make(map[string]int),
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
		return nil, internalError
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
			return nil, status.Errorf(codes.FailedPrecondition, "%q already depends on %q", taskName, dependencyName)
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

func (s *Server) RemoveDependency(ctx context.Context, req *taskspb.RemoveDependencyRequest) (*taskspb.Task, error) {
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

	depIdx := -1
	dependencies := task.GetDependencies()
	for i, existing := range dependencies {
		if existing == dependencyName {
			depIdx = i
			break
		}
	}
	if depIdx == -1 {
		return nil, status.Errorf(codes.FailedPrecondition, "no dependency exists from %q on %q", taskName, dependencyName)
	}
	task.Dependencies = append(dependencies[:depIdx], dependencies[depIdx+1:]...)

	for _, parent := range []string{
		taskName,
		dependencyName,
	} {
		event := &taskspb.Event{
			CreateTime: timestamppb.Now(),
			Comment:    req.GetComment(),
			Kind: &taskspb.Event_RemoveDependency_{RemoveDependency: &taskspb.Event_RemoveDependency{
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

func (s *Server) AddLabel(ctx context.Context, req *taskspb.AddLabelRequest) (*taskspb.Task, error) {
	taskName := req.GetTask()
	if taskName == "" {
		return nil, status.Error(codes.InvalidArgument, "empty task")
	}
	if !taskPattern.Matches(taskName) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid task name %q; want format %q", taskName, taskPattern)
	}
	labelName := req.GetLabel()
	if labelName == "" {
		return nil, status.Error(codes.InvalidArgument, "empty label")
	}
	if !labelPattern.Matches(labelName) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid label name %q; want format %q", labelName, labelPattern)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	taskIdx, ok := s.taskIndices[taskName]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "task %q not found", taskName)
	}
	task := s.tasks[taskIdx]

	if _, ok := s.labelIndices[labelName]; !ok {
		return nil, status.Errorf(codes.NotFound, "label %q not found", labelName)
	}

	labels := task.GetLabels()
	for _, existing := range labels {
		if existing == labelName {
			return nil, status.Errorf(codes.FailedPrecondition, "task %q already has label %q", taskName, labelName)
		}
	}

	task.Labels = append(labels, labelName)

	event := &taskspb.Event{
		CreateTime: timestamppb.Now(),
		Comment:    req.GetComment(),
		Kind: &taskspb.Event_AddLabel_{AddLabel: &taskspb.Event_AddLabel{
			Label: labelName,
		}},
	}
	if _, err := s.createEvent(ctx, taskName, event); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *Server) RemoveLabel(ctx context.Context, req *taskspb.RemoveLabelRequest) (*taskspb.Task, error) {
	taskName := req.GetTask()
	if taskName == "" {
		return nil, status.Error(codes.InvalidArgument, "empty task")
	}
	if !taskPattern.Matches(taskName) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid task name %q; want format %q", taskName, taskPattern)
	}
	labelName := req.GetLabel()
	if labelName == "" {
		return nil, status.Error(codes.InvalidArgument, "empty label")
	}
	if !labelPattern.Matches(labelName) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid label name %q; want format %q", labelName, labelPattern)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	taskIdx, ok := s.taskIndices[taskName]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "task %q not found", taskName)
	}
	task := s.tasks[taskIdx]

	if _, ok := s.labelIndices[labelName]; !ok {
		return nil, status.Errorf(codes.NotFound, "label %q not found", labelName)
	}

	depIdx := -1
	labels := task.GetLabels()
	for i, existing := range labels {
		if existing == labelName {
			depIdx = i
			break
		}
	}
	if depIdx == -1 {
		return nil, status.Errorf(codes.FailedPrecondition, "task %q does not have label %q", taskName, labelName)
	}
	task.Labels = append(labels[:depIdx], labels[depIdx+1:]...)

	event := &taskspb.Event{
		CreateTime: timestamppb.Now(),
		Comment:    req.GetComment(),
		Kind: &taskspb.Event_RemoveLabel_{RemoveLabel: &taskspb.Event_RemoveLabel{
			Label: labelName,
		}},
	}
	if _, err := s.createEvent(ctx, taskName, event); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *Server) CompleteTask(ctx context.Context, req *taskspb.CompleteTaskRequest) (*taskspb.Task, error) {
	taskName := req.GetTask()
	if taskName == "" {
		return nil, status.Error(codes.InvalidArgument, "empty task")
	}
	if !taskPattern.Matches(taskName) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid task %q doesn't match format %q", taskName, taskPattern)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	taskIdx, ok := s.taskIndices[taskName]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "task %q not found", taskName)
	}
	task := s.tasks[taskIdx]

	if task.GetCompleted() {
		return nil, status.Errorf(codes.FailedPrecondition, "task %q is already completed", taskName)
	}
	task.Completed = true

	event := &taskspb.Event{
		CreateTime: timestamppb.Now(),
		Comment:    req.GetComment(),
		Kind:       &taskspb.Event_Complete_{Complete: &taskspb.Event_Complete{}},
	}
	if _, err := s.createEvent(ctx, taskName, event); err != nil {
		return nil, err
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
			return nil, internalError
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

func (s *Server) ListLabels(ctx context.Context, req *taskspb.ListLabelsRequest) (*taskspb.ListLabelsResponse, error) {
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
	res := &taskspb.ListLabelsResponse{}
	if end >= len(s.labels) {
		res.Labels = s.labels[start:]
		res.NextPageToken = ""
	} else {
		res.Labels = s.labels[start:end]
		res.NextPageToken = next.String()
	}
	return res, nil
}

func (s *Server) CreateLabel(ctx context.Context, req *taskspb.CreateLabelRequest) (*taskspb.Label, error) {
	label := req.GetLabel()
	if label.GetDisplayName() == "" {
		return nil, status.Error(codes.InvalidArgument, "empty display name")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for _, existing := range s.labels {
		if label.GetDisplayName() == existing.GetDisplayName() {
			return nil, status.Errorf(codes.AlreadyExists, "label already exists: %q has display name %q", existing.GetName(), existing.GetDisplayName())
		}
	}

	idx := len(s.labels)
	v := resourcename.Values{
		"label": fmt.Sprint(idx + 1),
	}
	name, err := labelPattern.Render(v)
	if err != nil {
		log.Printf("CreateLabel failed to render label name: %v", err)
		return nil, internalError
	}
	label.Name = name
	s.labels = append(s.labels, label)
	s.labelIndices[label.Name] = idx
	return label, nil
}

func (s *Server) GetLabel(ctx context.Context, req *taskspb.GetLabelRequest) (*taskspb.Label, error) {
	name := req.GetName()
	if name == "" {
		return nil, status.Error(codes.InvalidArgument, "empty name")
	}
	if !labelPattern.Matches(name) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid name %q does not have format %q", name, labelPattern)
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	idx, ok := s.labelIndices[name]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "label %q not found", name)
	}
	return s.labels[idx], nil
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
		return nil, internalError
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
		return nil, internalError
	}
	taskEvents = append(taskEvents, name)

	event.Name = name
	event.Parent = parent

	index := len(s.events)
	s.eventIndices[name] = index
	s.events = append(s.events, event)
	return event, nil
}
