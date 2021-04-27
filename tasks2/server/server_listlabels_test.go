package server

import (
	"context"
	"testing"

	"github.com/Saser/pdp/testing/errtest"
	"github.com/Saser/pdp/testing/grpctest"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/testing/protocmp"

	taskspb "github.com/Saser/pdp/tasks2/tasks_go_proto"
)

func TestServer_ListLabels(t *testing.T) {
	ctx := context.Background()
	c := setup(t)

	var want []*taskspb.Label
	for _, label := range []*taskspb.Label{
		{DisplayName: "High Priority"},
		{DisplayName: "Really High Priority"},
		{DisplayName: "Actually High Priority"},
	} {
		want = append(want, c.CreateLabelT(ctx, t, &taskspb.CreateLabelRequest{
			Label: label,
		}))
	}

	req := &taskspb.ListLabelsRequest{}
	res, err := c.ListLabels(ctx, req)
	if err != nil {
		t.Errorf("ListLabels(%v) err = %v; want nil", req, err)
	}
	if token := res.GetNextPageToken(); token != "" {
		t.Errorf("res.GetNextPageToken() = %q; want an empty string", token)
	}
	got := res.GetLabels()
	if diff := cmp.Diff(want, got, protocmp.Transform(), cmpopts.SortSlices(taskLessFunc)); diff != "" {
		t.Errorf("diff between created and listed labels (-want +got)\n%s", diff)
	}
}

func TestServer_ListLabels_Pagination(t *testing.T) {
	ctx := context.Background()
	c := setup(t)

	var want []*taskspb.Label
	for _, label := range []*taskspb.Label{
		{DisplayName: "High Priority"},
		{DisplayName: "Really High Priority"},
		{DisplayName: "Actually High Priority"},
	} {
		want = append(want, c.CreateLabelT(ctx, t, &taskspb.CreateLabelRequest{
			Label: label,
		}))
	}

	var got []*taskspb.Label
	res1 := c.ListLabelsT(ctx, t, &taskspb.ListLabelsRequest{
		PageSize: 1,
	})
	if res1.GetNextPageToken() == "" {
		t.Fatal(`res1.GetNextPageToken() = ""; want non-empty`)
	}
	got = append(got, res1.GetLabels()...)

	res2 := c.ListLabelsT(ctx, t, &taskspb.ListLabelsRequest{
		PageToken: res1.GetNextPageToken(),
	})
	if token := res2.GetNextPageToken(); token != "" {
		t.Errorf("res2.GetNextPageToken() = %q; want an empty string", token)
	}
	got = append(got, res2.GetLabels()...)

	if diff := cmp.Diff(want, got, protocmp.Transform(), cmpopts.SortSlices(labelLessFunc)); diff != "" {
		t.Errorf("diff between created and listed labels (-want +got)\n%s", diff)
	}
}

func TestServer_ListLabels_Errors(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	for _, tt := range []struct {
		name string
		req  *taskspb.ListLabelsRequest
		tf   errtest.TestFunc
	}{
		{
			name: "NegativePageSize",
			req: &taskspb.ListLabelsRequest{
				PageSize: -1,
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("negative"),
			),
		},
		{
			name: "InvalidPageToken",
			req: &taskspb.ListLabelsRequest{
				PageToken: "invalid-page-token",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains(`"invalid-page-token"`),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.ListLabels(ctx, tt.req)
			tt.tf(t, err)
		})
	}
}
