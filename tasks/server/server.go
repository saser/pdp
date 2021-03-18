package server

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/Saser/pdp/aip/pagetoken"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	taskspb "github.com/Saser/pdp/tasks/tasks_go_proto"
)

const maxPageSize = 50

type Server struct {
	taskspb.UnimplementedTasksServer

	mu          sync.Mutex
	tasks       []*taskspb.Task
	taskIndices map[string]int
}

func New() *Server {
	return &Server{
		taskIndices: make(map[string]int),
	}
}

func (s *Server) GetTask(ctx context.Context, req *taskspb.GetTaskRequest) (*taskspb.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	name := req.GetName()
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty name")
	}
	if !strings.HasPrefix(name, "tasks/") {
		return nil, status.Errorf(codes.InvalidArgument, "invalid task name %q, expected name of format %q", name, "tasks/{task}")
	}
	idx, ok := s.taskIndices[name]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "task not found: %q", name)
	}
	return s.tasks[idx], nil
}

func (s *Server) ListTasks(ctx context.Context, req *taskspb.ListTasksRequest) (*taskspb.ListTasksResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	pageSize := req.GetPageSize()
	if pageSize < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "negative page size: %d", pageSize)
	}
	if pageSize == 0 || pageSize > maxPageSize {
		pageSize = maxPageSize
	}

	pt, err := pagetoken.Parse(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid page token: %q", req.GetPageToken())
	}

	// allTasks is the list of tasks we are paginating over. The offset represented by the page
	// token applies to this list.
	var allTasks []*taskspb.Task
	for _, task := range s.tasks {
		if task.GetDeleted() && !req.GetShowDeleted() {
			continue
		}
		allTasks = append(allTasks, task)
	}

	// page is the list of tasks that will be returned by this request.
	var page []*taskspb.Task
	for i := int(pt.Offset()); i < len(allTasks); i++ {
		if len(page) >= int(pageSize) {
			break
		}
		page = append(page, allTasks[i])
	}

	nextPageToken := ""
	if next := pt.Next(pageSize); next.Offset() < int32(len(allTasks)) {
		nextPageToken = next.String()
	}

	return &taskspb.ListTasksResponse{
		Tasks:         page,
		NextPageToken: nextPageToken,
	}, nil
}

func (s *Server) CreateTask(ctx context.Context, req *taskspb.CreateTaskRequest) (*taskspb.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Validate the task that was passed in.
	task := req.GetTask()
	if name := task.GetName(); name != "" {
		return nil, status.Errorf(codes.InvalidArgument, `"name" must be empty, was %q`, name)
	}
	if task.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "empty title")
	}
	if task.GetDeleted() {
		return nil, status.Error(codes.InvalidArgument, `"deleted" must be false`)
	}
	if task.GetCompleted() {
		return nil, status.Error(codes.InvalidArgument, `"completed" must be false (use the SetCompleted method to set the "completed" field)`)
	}

	// Task is valid, so go ahead and store it.
	task.Name = fmt.Sprintf("tasks/%d", len(s.tasks)+1)
	s.taskIndices[task.Name] = len(s.tasks)
	s.tasks = append(s.tasks, task)
	return task, nil
}

func (s *Server) DeleteTask(ctx context.Context, req *taskspb.DeleteTaskRequest) (*taskspb.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	name := req.GetName()
	idx, ok := s.taskIndices[name]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "task not found: %q", name)
	}
	task := s.tasks[idx]
	task.Deleted = true
	return task, nil
}
