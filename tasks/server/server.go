package server

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	taskspb "github.com/Saser/pdp/tasks/tasks_go_proto"
)

type Server struct {
	taskspb.UnimplementedTasksServer

	mu    sync.Mutex
	tasks map[string]*taskspb.Task
}

func New() *Server {
	return &Server{
		tasks: make(map[string]*taskspb.Task),
	}
}

func (s *Server) GetTask(ctx context.Context, req *taskspb.GetTaskRequest) (*taskspb.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	name := req.GetName()
	if !strings.HasPrefix(name, "tasks/") {
		return nil, status.Errorf(codes.InvalidArgument, "invalid task name: %q", name)
	}
	task, ok := s.tasks[name]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "task not found: %q", name)
	}
	return task, nil
}

func (s *Server) CreateTask(ctx context.Context, req *taskspb.CreateTaskRequest) (*taskspb.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	task := req.GetTask()
	task.Name = fmt.Sprintf("tasks/%d", len(s.tasks)+1)
	task.Deleted = false
	task.Completed = false
	s.tasks[task.Name] = task
	return task, nil
}
