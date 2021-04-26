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

func TestServer_GetLabel(t *testing.T) {
	ctx := context.Background()
	c := setup(t)

	want := c.CreateLabelT(ctx, t, &taskspb.CreateLabelRequest{
		Label: &taskspb.Label{
			DisplayName: "some label",
		},
	})

	req := &taskspb.GetLabelRequest{
		Name: want.GetName(),
	}
	got, err := c.GetLabel(ctx, req)
	if err != nil {
		t.Errorf("GetLabel(%v) err = %v; want nil", req, err)
	}
	if diff := cmp.Diff(want, got, protocmp.Transform()); diff != "" {
		t.Errorf("diff between CreateLabel and GetLabel (-want +got)\n%s", diff)
	}
}

func TestServer_GetLabel_Errors(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	for _, tt := range []struct {
		name string
		req  *taskspb.GetLabelRequest
		tf   errtest.TestFunc
	}{
		{
			name: "EmptyName",
			req: &taskspb.GetLabelRequest{
				Name: "",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("empty name"),
			),
		},
		{
			name: "InvalidName",
			req: &taskspb.GetLabelRequest{
				Name: "foobar/1",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("invalid name"),
				errtest.ErrorContains(`"foobar/1"`),
				errtest.ErrorContains(`"labels/{label}"`),
			),
		},
		{
			name: "NotFound",
			req: &taskspb.GetLabelRequest{
				Name: "labels/999",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.NotFound),
				errtest.ErrorContains(`"labels/999"`),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.GetLabel(ctx, tt.req)
			tt.tf(t, err)
		})
	}
}
