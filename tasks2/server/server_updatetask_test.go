package server

import (
	"context"
	"testing"

	"github.com/Saser/pdp/testing/errtest"
	"github.com/Saser/pdp/testing/grpctest"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"

	taskspb "github.com/Saser/pdp/tasks2/tasks_go_proto"
)

func TestServer_UpdateTask(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	task := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "Before update",
		},
	})
	want := proto.Clone(task).(*taskspb.Task)
	want.Title = "After update"

	req := &taskspb.UpdateTaskRequest{
		Task: want,
	}
	got, err := c.UpdateTask(ctx, req)
	if err != nil {
		t.Fatalf("UpdateTask(%v) err = %v; want nil", req, err)
	}
	if diff := cmp.Diff(want, got, protocmp.Transform()); diff != "" {
		t.Errorf("unexpected result of UpdateTask (-want +got)\n%s", diff)
	}
}

func TestServer_UpdateTask_EmptyTitle(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	want := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "Before update",
		},
	})
	got := c.UpdateTaskT(ctx, t, &taskspb.UpdateTaskRequest{
		Task: func() *taskspb.Task {
			updated := proto.Clone(want).(*taskspb.Task)
			updated.Title = ""
			return updated
		}(),
	})
	if diff := cmp.Diff(want, got, protocmp.Transform()); diff != "" {
		t.Errorf("unexpected result of UpdateTask (-want +got)\n%s", diff)
	}
}

func TestServer_UpdateTask_Errors(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	task := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "Some task",
		},
	})
	dependency := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "A dependency",
		},
	})
	label := c.CreateLabelT(ctx, t, &taskspb.CreateLabelRequest{
		Label: &taskspb.Label{
			DisplayName: "A label",
		},
	})
	clone := func(f func(task *taskspb.Task)) *taskspb.Task {
		modified := proto.Clone(task).(*taskspb.Task)
		f(modified)
		return modified
	}
	for _, tt := range []struct {
		name string
		req  *taskspb.UpdateTaskRequest
		tf   errtest.TestFunc
	}{
		{
			name: "EmptyName",
			req: &taskspb.UpdateTaskRequest{
				Task: &taskspb.Task{
					Name:  "",
					Title: "This has no name",
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("empty name"),
			),
		},
		{
			name: "InvalidName",
			req: &taskspb.UpdateTaskRequest{
				Task: &taskspb.Task{
					Name:  "invalid/1",
					Title: "Invalid name",
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("invalid name"),
				errtest.ErrorContains(`"tasks/{task}"`),
			),
		},
		{
			name: "NotFound",
			req: &taskspb.UpdateTaskRequest{
				Task: &taskspb.Task{
					Name:  "tasks/999",
					Title: "This does not exist",
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.NotFound),
				errtest.ErrorContains(`"tasks/999"`),
			),
		},
		{
			name: "UpdatesCompleted",
			req: &taskspb.UpdateTaskRequest{
				Task: clone(func(task *taskspb.Task) {
					task.Completed = true
				}),
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains(`"completed"`),
				errtest.ErrorContains("CompleteTask"),
				errtest.ErrorContains("UncompleteTask"),
			),
		},
		{
			name: "UpdatesDependencies",
			req: &taskspb.UpdateTaskRequest{
				Task: clone(func(task *taskspb.Task) {
					task.Dependencies = []string{
						dependency.GetName(),
					}
				}),
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains(`"dependencies"`),
				errtest.ErrorContains("AddDependency"),
				errtest.ErrorContains("RemoveDependency"),
			),
		},
		{
			name: "UpdatesLabels",
			req: &taskspb.UpdateTaskRequest{
				Task: clone(func(task *taskspb.Task) {
					task.Labels = []string{
						label.GetName(),
					}
				}),
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains(`"labels"`),
				errtest.ErrorContains("AddLabel"),
				errtest.ErrorContains("RemoveLabel"),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.UpdateTask(ctx, tt.req)
			if err == nil {
				t.Fatalf("UpdateTask(%v) err = nil; want non-nil", tt.req)
			}
			tt.tf(t, err)
		})
	}
}
