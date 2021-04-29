package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/Saser/pdp/testing/errtest"
	"github.com/Saser/pdp/testing/grpctest"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/testing/protocmp"

	taskspb "github.com/Saser/pdp/tasks2/tasks_go_proto"
)

func TestServer_CompleteTask(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	task := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "some task",
		},
	})

	req := &taskspb.CompleteTaskRequest{
		Task:    task.GetName(),
		Comment: "Task completed",
	}
	got, err := c.CompleteTask(ctx, req)
	if err != nil {
		t.Errorf("CompleteTask(%v) err = %v; want nil", req, err)
	}
	if diff := cmp.Diff(task, got, protocmp.Transform(), protocmp.IgnoreFields(&taskspb.Task{}, "completed")); diff != "" {
		t.Errorf("unexpected task diff after CompleteTask (-want +got)\n%s", diff)
	}
	if !got.GetCompleted() {
		t.Error("GetCompleted() = false; want true")
	}
}

func TestServer_CompleteTask_Errors(t *testing.T) {
	ctx := context.Background()
	c := setup(t)

	for _, tt := range []struct {
		name string
		req  *taskspb.CompleteTaskRequest
		tf   errtest.TestFunc
	}{
		{
			name: "EmptyTask",
			req: &taskspb.CompleteTaskRequest{
				Task: "",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("empty task"),
			),
		},
		{
			name: "InvalidTask",
			req: &taskspb.CompleteTaskRequest{
				Task: "invalid/123",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("invalid task"),
				errtest.ErrorContains(`"invalid/123"`),
				errtest.ErrorContains(`"tasks/{task}"`),
			),
		},
		{
			name: "NotFoundTask",
			req: &taskspb.CompleteTaskRequest{
				Task: "tasks/999",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.NotFound),
				errtest.ErrorContains("not found"),
				errtest.ErrorContains(`"tasks/999"`),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.CompleteTask(ctx, tt.req)
			if err == nil {
				t.Fatalf("CompleteTask(%v) err = nil; want non-nil", tt.req)
			}
			tt.tf(t, err)
		})
	}
}

func TestServer_CompleteTask_AlreadyCompleted(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	task := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "some task",
		},
	})
	task = c.CompleteTaskT(ctx, t, &taskspb.CompleteTaskRequest{
		Task:    task.GetName(),
		Comment: "First completion",
	})

	req := &taskspb.CompleteTaskRequest{
		Task:    task.GetName(),
		Comment: "Second completion",
	}
	_, err := c.CompleteTask(ctx, req)
	if err == nil {
		t.Fatalf("CompleteTask(%v) err = nil; want non-nil", req)
	}
	tf := errtest.All(
		grpctest.WantCode(codes.FailedPrecondition),
		errtest.ErrorContains(fmt.Sprintf("%q", task.GetName())),
		errtest.ErrorContains("already completed"),
	)
	tf(t, err)
}
