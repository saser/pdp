package server

import (
	"context"
	"testing"

	"github.com/Saser/pdp/testing/grpctest"
	"google.golang.org/grpc"

	taskspb "github.com/Saser/pdp/tasks/tasks_go_proto"
)

type testTasksClient struct {
	taskspb.TasksClient
}

func setup(t *testing.T) testTasksClient {
	t.Helper()
	register := func(s *grpc.Server) { taskspb.RegisterTasksServer(s, New()) }
	cc := grpctest.NewClientConnT(t, register)
	return testTasksClient{
		TasksClient: taskspb.NewTasksClient(cc),
	}
}

func (c testTasksClient) CreateTaskT(ctx context.Context, t *testing.T, req *taskspb.CreateTaskRequest) *taskspb.Task {
	t.Helper()
	task, err := c.CreateTask(ctx, req)
	if err != nil {
		t.Errorf("CreateTask(%v) err = %v; want nil", req, err)
	}
	return task
}

func Test_setup(t *testing.T) {
	_ = setup(t)
}
