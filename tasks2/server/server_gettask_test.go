package server

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/testing/protocmp"

	taskspb "github.com/Saser/pdp/tasks2/tasks_go_proto"
	"github.com/Saser/pdp/testing/errtest"
	"github.com/Saser/pdp/testing/grpctest"
)

func TestServer_GetTask(t *testing.T) {
	ctx := context.Background()
	c := setup(t)

	want := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "some task",
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
		t.Errorf("diff between CreateTask and GetTask (-want +got)\n%s", diff)
	}
}

func TestServer_GetTask_Errors(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
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
				Name: "foobar/1",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("invalid name"),
				errtest.ErrorContains(`"foobar/1"`),
				errtest.ErrorContains(`"tasks/{task}"`),
			),
		},
		{
			name: "NotFound",
			req: &taskspb.GetTaskRequest{
				Name: "tasks/999",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.NotFound),
				errtest.ErrorContains(`"tasks/999"`),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.GetTask(ctx, tt.req)
			tt.tf(t, err)
		})
	}
}
