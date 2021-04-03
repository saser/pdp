package server

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Saser/pdp/aip/pagetoken"
	taskspb "github.com/Saser/pdp/tasks2/tasks_go_proto"
)

const maxPageSize = 100

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
	name := req.GetName()
	if name == "" {
		return nil, status.Error(codes.InvalidArgument, "empty name")
	}
	if prefix := "tasks/"; !strings.HasPrefix(name, prefix) {
		return nil, status.Errorf(codes.InvalidArgument, `invalid name %q does not have format "tasks/{task}"`, name)
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
	task.Name = fmt.Sprintf("tasks/%d", idx+1)
	s.tasks = append(s.tasks, task)
	s.taskIndices[task.Name] = idx
	return task, nil
}
