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

func TestServer_UpdateLabel(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	label := c.CreateLabelT(ctx, t, &taskspb.CreateLabelRequest{
		Label: &taskspb.Label{
			DisplayName: "Before update",
		},
	})
	want := proto.Clone(label).(*taskspb.Label)
	want.DisplayName = "After update"

	req := &taskspb.UpdateLabelRequest{
		Label: want,
	}
	got, err := c.UpdateLabel(ctx, req)
	if err != nil {
		t.Fatalf("UpdateLabel(%v) err = %v; want nil", req, err)
	}
	if diff := cmp.Diff(want, got, protocmp.Transform()); diff != "" {
		t.Errorf("unexpected result of UpdateLabel (-want +got)\n%s", diff)
	}
}

func TestServer_UpdateLabel_EmptyDisplayName(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	want := c.CreateLabelT(ctx, t, &taskspb.CreateLabelRequest{
		Label: &taskspb.Label{
			DisplayName: "Before update",
		},
	})
	got := c.UpdateLabelT(ctx, t, &taskspb.UpdateLabelRequest{
		Label: func() *taskspb.Label {
			updated := proto.Clone(want).(*taskspb.Label)
			updated.DisplayName = ""
			return updated
		}(),
	})
	if diff := cmp.Diff(want, got, protocmp.Transform()); diff != "" {
		t.Errorf("unexpected result of UpdateLabel (-want +got)\n%s", diff)
	}
}

func TestServer_UpdateLabel_Errors(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	for _, tt := range []struct {
		name string
		req  *taskspb.UpdateLabelRequest
		tf   errtest.TestFunc
	}{
		{
			name: "EmptyName",
			req: &taskspb.UpdateLabelRequest{
				Label: &taskspb.Label{
					Name:        "",
					DisplayName: "This has no name",
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("empty name"),
			),
		},
		{
			name: "InvalidName",
			req: &taskspb.UpdateLabelRequest{
				Label: &taskspb.Label{
					Name:        "invalid/1",
					DisplayName: "Invalid name",
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("invalid name"),
				errtest.ErrorContains(`"labels/{label}"`),
			),
		},
		{
			name: "NotFound",
			req: &taskspb.UpdateLabelRequest{
				Label: &taskspb.Label{
					Name:        "labels/999",
					DisplayName: "This does not exist",
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.NotFound),
				errtest.ErrorContains(`"labels/999"`),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.UpdateLabel(ctx, tt.req)
			if err == nil {
				t.Fatalf("UpdateLabel(%v) err = nil; want non-nil", tt.req)
			}
			tt.tf(t, err)
		})
	}
}
