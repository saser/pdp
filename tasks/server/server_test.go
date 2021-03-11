package server

import (
	"context"
	"testing"

	"github.com/Saser/pdp/testing/grpctest"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		want codes.Code
	}{
		{
			name: "Empty name",
			req: &taskspb.GetTaskRequest{
				Name: "",
			},
			want: codes.InvalidArgument,
		},
		{
			name: "Invalid name",
			req: &taskspb.GetTaskRequest{
				Name: "issues/1",
			},
			want: codes.InvalidArgument,
		},
		{
			name: "Valid but non-existent name",
			req: &taskspb.GetTaskRequest{
				Name: "tasks/2",
			},
			want: codes.NotFound,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.GetTask(ctx, tt.req)
			if err == nil {
				t.Errorf("GetTask(%v) err = nil; want non-nil", tt.req)
			}
			if got := status.Code(err); got != tt.want {
				t.Errorf("status.Code(%v) = %v; want %v", err, got, tt.want)
			}
		})
	}
}
