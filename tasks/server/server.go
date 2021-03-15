package server

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"

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

	pageSize := int(req.GetPageSize())
	if pageSize < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "negative page size: %d", pageSize)
	}
	if pageSize == 0 || pageSize > maxPageSize {
		pageSize = maxPageSize
	}

	offset := 0
	if tok := req.GetPageToken(); tok != "" {
		var err error
		offset, err = strconv.Atoi(tok)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid page token: %q", tok)
		}
	}

	var tasks []*taskspb.Task
	for i := offset; i < len(s.tasks) && i-offset < pageSize; i++ {
		task := s.tasks[i]
		if task.GetDeleted() && !req.GetShowDeleted() {
			continue
		}
		tasks = append(tasks, s.tasks[i])
	}

	nextPageToken := ""
	if nextOffset := offset + len(tasks); nextOffset < len(s.tasks) {
		nextPageToken = strconv.Itoa(nextOffset)
	}

	return &taskspb.ListTasksResponse{
		Tasks:         tasks,
		NextPageToken: nextPageToken,
	}, nil
}

func (s *Server) CreateTask(ctx context.Context, req *taskspb.CreateTaskRequest) (*taskspb.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	task := req.GetTask()
	task.Name = fmt.Sprintf("tasks/%d", len(s.tasks)+1)
	task.Deleted = false
	task.Completed = false
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
