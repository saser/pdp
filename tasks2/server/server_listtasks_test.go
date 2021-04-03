package server

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/testing/protocmp"

	taskspb "github.com/Saser/pdp/tasks2/tasks_go_proto"
	"github.com/Saser/pdp/testing/errtest"
	"github.com/Saser/pdp/testing/grpctest"
)

func TestServer_ListTasks(t *testing.T) {
	ctx := context.Background()
	c := setup(t)

	var want []*taskspb.Task
	for _, task := range []*taskspb.Task{
		{Title: "Buy milk"},
		{Title: "Get a haircut"},
		{Title: "Find a job"},
	} {
		want = append(want, c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
			Task: task,
		}))
	}

	req := &taskspb.ListTasksRequest{}
	res, err := c.ListTasks(ctx, req)
	if err != nil {
		t.Errorf("ListTasks(%v) err = %v; want nil", req, err)
	}
	if token := res.GetNextPageToken(); token != "" {
		t.Errorf("res.GetNextPageToken() = %q; want an empty string", token)
	}
	got := res.GetTasks()
	if diff := cmp.Diff(want, got, protocmp.Transform(), cmpopts.SortSlices(taskLessFunc)); diff != "" {
		t.Errorf("diff between created and listed tasks (-want +got)\n%s", diff)
	}
}

func TestServer_ListTasks_Pagination(t *testing.T) {
	ctx := context.Background()
	c := setup(t)

	var want []*taskspb.Task
	for _, task := range []*taskspb.Task{
		{Title: "Buy milk"},
		{Title: "Get a haircut"},
		{Title: "Find a job"},
	} {
		want = append(want, c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
			Task: task,
		}))
	}

	var got []*taskspb.Task
	res1 := c.ListTasksT(ctx, t, &taskspb.ListTasksRequest{
		PageSize: 1,
	})
	if res1.GetNextPageToken() == "" {
		t.Fatal(`res1.GetNextPageToken() = ""; want non-empty`)
	}
	got = append(got, res1.GetTasks()...)

	res2 := c.ListTasksT(ctx, t, &taskspb.ListTasksRequest{
		PageToken: res1.GetNextPageToken(),
	})
	if token := res2.GetNextPageToken(); token != "" {
		t.Errorf("res2.GetNextPageToken() = %q; want an empty string", token)
	}
	got = append(got, res2.GetTasks()...)

	if diff := cmp.Diff(want, got, protocmp.Transform(), cmpopts.SortSlices(taskLessFunc)); diff != "" {
		t.Errorf("diff between created and listed tasks (-want +got)\n%s", diff)
	}
}

func TestListTasks_Errors(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	for _, tt := range []struct {
		name string
		req  *taskspb.ListTasksRequest
		tf   errtest.TestFunc
	}{
		{
			name: "NegativePageSize",
			req: &taskspb.ListTasksRequest{
				PageSize: -1,
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("negative"),
			),
		},
		{
			name: "InvalidPageToken",
			req: &taskspb.ListTasksRequest{
				PageToken: "invalid-page-token",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains(`"invalid-page-token"`),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.ListTasks(ctx, tt.req)
			tt.tf(t, err)
		})
	}
}
