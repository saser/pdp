package service

import (
	"context"
	"fmt"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	tasks3pb "github.com/Saser/pdp/tasks3/tasks3_go_proto"
)

type Server struct {
	tasks3pb.UnimplementedTasksServer

	mu    sync.Mutex
	tasks []*tasks3pb.Task
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetTask(ctx context.Context, req *tasks3pb.GetTaskRequest) (*tasks3pb.Task, error) {
	name := req.GetName()
	if name == "" {
		return nil, status.Error(codes.InvalidArgument, "empty name")
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	for _, task := range s.tasks {
		if task.GetName() == name {
			return proto.Clone(task).(*tasks3pb.Task), nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "task %q not found", name)
}

func (s *Server) ListTasks(ctx context.Context, req *tasks3pb.ListTasksRequest) (*tasks3pb.ListTasksResponse, error) {
	if ps := req.GetPageSize(); ps != 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid page size: %d", ps)
	}
	if tok := req.GetPageToken(); tok != "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid page token: %q", tok)
	}
	res := &tasks3pb.ListTasksResponse{}
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, task := range s.tasks {
		res.Tasks = append(res.Tasks, proto.Clone(task).(*tasks3pb.Task))
	}
	return res, nil
}

func (s *Server) CreateTask(ctx context.Context, req *tasks3pb.CreateTaskRequest) (*tasks3pb.Task, error) {
	newTask := proto.Clone(req.GetTask()).(*tasks3pb.Task)
	if newTask.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "empty title")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	newTask.Name = fmt.Sprintf("tasks/%d", len(s.tasks)+1)
	s.tasks = append(s.tasks, newTask)
	return newTask, nil
}
