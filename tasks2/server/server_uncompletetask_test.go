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

func TestServer_UncompleteTask(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	task := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "some task",
		},
	})
	task = c.CompleteTaskT(ctx, t, &taskspb.CompleteTaskRequest{
		Task:    task.GetName(),
		Comment: "Completing so it can be uncompleted later",
	})

	req := &taskspb.UncompleteTaskRequest{
		Task:    task.GetName(),
		Comment: "Uncompleting previously completed task",
	}
	got, err := c.UncompleteTask(ctx, req)
	if err != nil {
		t.Errorf("UncompleteTask(%v) err = %v; want nil", req, err)
	}
	if diff := cmp.Diff(task, got, protocmp.Transform(), protocmp.IgnoreFields(&taskspb.Task{}, "completed")); diff != "" {
		t.Errorf("unexpected task diff after UncompleteTask (-want +got)\n%s", diff)
	}
	if got.GetCompleted() {
		t.Error("GetCompleted() = true; want false")
	}
}

func TestServer_UncompleteTask_Errors(t *testing.T) {
	ctx := context.Background()
	c := setup(t)

	for _, tt := range []struct {
		name string
		req  *taskspb.UncompleteTaskRequest
		tf   errtest.TestFunc
	}{
		{
			name: "EmptyTask",
			req: &taskspb.UncompleteTaskRequest{
				Task: "",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("empty task"),
			),
		},
		{
			name: "InvalidTask",
			req: &taskspb.UncompleteTaskRequest{
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
			req: &taskspb.UncompleteTaskRequest{
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
			_, err := c.UncompleteTask(ctx, tt.req)
			if err == nil {
				t.Fatalf("UncompleteTask(%v) err = nil; want non-nil", tt.req)
			}
			tt.tf(t, err)
		})
	}
}

func TestServer_UncompleteTask_NotCompleted(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	task := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "some task",
		},
	})

	req := &taskspb.UncompleteTaskRequest{
		Task:    task.GetName(),
		Comment: "Uncomplete already uncompleted task",
	}
	_, err := c.UncompleteTask(ctx, req)
	if err == nil {
		t.Fatalf("UncompleteTask(%v) err = nil; want non-nil", req)
	}
	tf := errtest.All(
		grpctest.WantCode(codes.FailedPrecondition),
		errtest.ErrorContains(fmt.Sprintf("%q", task.GetName())),
		errtest.ErrorContains("not completed"),
	)
	tf(t, err)
}
