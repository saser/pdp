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

func TestServer_CreateLabel(t *testing.T) {
	ctx := context.Background()
	c := setup(t)

	req := &taskspb.CreateLabelRequest{
		Label: &taskspb.Label{
			DisplayName: "some label",
		},
	}
	label, err := c.CreateLabel(ctx, req)
	if err != nil {
		t.Errorf("CreateLabel(%v) err = %v; want nil", req, err)
	}
	if diff := cmp.Diff(req.GetLabel(), label, protocmp.Transform(), protocmp.IgnoreFields(&taskspb.Label{}, "name")); diff != "" {
		t.Errorf("diff between requested and returned label (-want +got)\n%s", diff)
	}
}

func TestServer_CreateLabel_Errors(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	for _, tt := range []struct {
		name string
		req  *taskspb.CreateLabelRequest
		tf   errtest.TestFunc
	}{
		{
			name: "EmptyDisplayName",
			req: &taskspb.CreateLabelRequest{
				Label: &taskspb.Label{
					DisplayName: "",
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("empty display name"),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.CreateLabel(ctx, tt.req)
			tt.tf(t, err)
		})
	}
}

func TestServer_CreateLabel_DuplicateDisplayName(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	label := c.CreateLabelT(ctx, t, &taskspb.CreateLabelRequest{
		Label: &taskspb.Label{
			DisplayName: "Duplicate display name",
		},
	})

	req := &taskspb.CreateLabelRequest{
		Label: &taskspb.Label{
			DisplayName: label.GetDisplayName(),
		},
	}
	_, err := c.CreateLabel(ctx, req)
	if err == nil {
		t.Fatalf("CreateLabel(%v) err = nil; want non-nil", req)
	}
	tf := errtest.All(
		grpctest.WantCode(codes.AlreadyExists),
		errtest.ErrorContains("label already exists"),
		errtest.ErrorContains(fmt.Sprintf("%q", label.GetDisplayName())),
	)
	tf(t, err)
}
