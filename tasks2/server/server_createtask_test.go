package server

import (
	"context"
	"testing"

	"github.com/Saser/pdp/testing/errtest"
	"github.com/Saser/pdp/testing/grpctest"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/testing/protocmp"

	taskspb "github.com/Saser/pdp/tasks2/tasks_go_proto"
)

func TestServer_CreateTask(t *testing.T) {
	ctx := context.Background()
	c := setup(t)

	req := &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "some task",
		},
	}
	task, err := c.CreateTask(ctx, req)
	if err != nil {
		t.Errorf("CreateTask(%v) err = %v; want nil", req, err)
	}
	if diff := cmp.Diff(req.GetTask(), task, protocmp.Transform(), protocmp.IgnoreFields(&taskspb.Task{}, "name")); diff != "" {
		t.Errorf("diff between requested and returned task (-want +got)\n%s", diff)
	}
}

func TestServer_CreateTask_Errors(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	for _, tt := range []struct {
		name string
		req  *taskspb.CreateTaskRequest
		tf   errtest.TestFunc
	}{
		{
			name: "EmptyTitle",
			req: &taskspb.CreateTaskRequest{
				Task: &taskspb.Task{
					Title: "",
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("empty title"),
			),
		},
		{
			name: "WithCompleted",
			req: &taskspb.CreateTaskRequest{
				Task: &taskspb.Task{
					Title:     "some task",
					Completed: true,
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains(`"completed"`),
			),
		},
		{
			name: "WithDependencies",
			req: &taskspb.CreateTaskRequest{
				Task: &taskspb.Task{
					Title: "some task",
					Dependencies: []string{
						"tasks/123",
						"tasks/456",
					},
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains(`"dependencies"`),
			),
		},
		{
			name: "WithLabels",
			req: &taskspb.CreateTaskRequest{
				Task: &taskspb.Task{
					Title: "some task",
					Labels: []string{
						"labels/123",
						"labels/456",
					},
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains(`"labels"`),
			),
		},
		{
			name: "WithDeferral",
			req: &taskspb.CreateTaskRequest{
				Task: &taskspb.Task{
					Title: "some task",
					Deferral: &taskspb.Deferral{
						Constraints: []*taskspb.Deferral_Constraint{
							{Kind: &taskspb.Deferral_Constraint_Dependency{Dependency: "tasks/123"}},
						},
					},
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains(`"deferral"`),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.CreateTask(ctx, tt.req)
			tt.tf(t, err)
		})
	}
}
