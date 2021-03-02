package server

import (
	"context"
	"fmt"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

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
	task, ok := s.tasks[req.GetName()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "task not found: %q", req.GetName())
	}
	return task, nil
}

func (s *Server) CreateTask(ctx context.Context, req *taskspb.CreateTaskRequest) (*taskspb.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	name := fmt.Sprintf("tasks/%d", len(s.tasks)+1)
	task := proto.Clone(req.GetTask()).(*taskspb.Task)
	task.Name = name
	s.tasks[name] = task
	return task, nil
}

func (s *Server) ListTasks(ctx context.Context, req *taskspb.ListTasksRequest) (*taskspb.ListTasksResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var tasks []*taskspb.Task
	for _, t := range s.tasks {
		tasks = append(tasks, t)
	}
	return &taskspb.ListTasksResponse{
		Tasks: tasks,
	}, nil
}
