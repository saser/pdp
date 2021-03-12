package server

import (
	"context"
	"testing"

	"github.com/Saser/pdp/testing/errtest"
	"github.com/Saser/pdp/testing/grpctest"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/testing/protocmp"

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

func (c testTasksClient) DeleteTaskT(ctx context.Context, t *testing.T, req *taskspb.DeleteTaskRequest) *taskspb.Task {
	t.Helper()
	task, err := c.DeleteTask(ctx, req)
	if err != nil {
		t.Errorf("DeleteTask(%v) err = %v; want nil", req, err)
	}
	return task
}

func Test_setup(t *testing.T) {
	_ = setup(t)
}

func TestGetTask_OK(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	want := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "My task",
		},
	})
	req := &taskspb.GetTaskRequest{
		Name: want.GetName(),
	}
	got, err := c.GetTask(ctx, req)
	if err != nil {
		t.Errorf("GetTask(%v) err = %v; want nil", req, err)
	}
	if diff := cmp.Diff(want, got, protocmp.Transform()); diff != "" {
		t.Errorf("difference between CreateTask() and GetTask() (-want +got)\n%s", diff)
	}
}

func TestGetTask_Errors(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "Some task",
		},
	})
	for _, tt := range []struct {
		name string
		req  *taskspb.GetTaskRequest
		tf   errtest.TestFunc
	}{
		{
			name: "EmptyName",
			req: &taskspb.GetTaskRequest{
				Name: "",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("empty name"),
			),
		},
		{
			name: "InvalidName",
			req: &taskspb.GetTaskRequest{
				Name: "issues/1",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("issues/1"),
				errtest.ErrorContains("invalid task name"),
				errtest.ErrorContains("tasks/{task}"),
			),
		},
		{
			name: "NonExistentTask",
			req: &taskspb.GetTaskRequest{
				Name: "tasks/2",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.NotFound),
				errtest.ErrorContains("tasks/2"),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.GetTask(ctx, tt.req)
			tt.tf(t, err)
		})
	}
}

func TestGetTask_AfterDeletion(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	task := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "Some task",
		},
	})
	deleted := c.DeleteTaskT(ctx, t, &taskspb.DeleteTaskRequest{
		Name: task.GetName(),
	})
	req := &taskspb.GetTaskRequest{
		Name: task.GetName(),
	}
	got, err := c.GetTask(ctx, req)
	if err != nil {
		t.Errorf("GetTask(%v) err = %v; want nil", req, err)
	}
	if diff := cmp.Diff(deleted, got, protocmp.Transform()); diff != "" {
		t.Errorf("diff between DeleteTask() and GetTask() (-want +got)\n%s", diff)
	}
}
