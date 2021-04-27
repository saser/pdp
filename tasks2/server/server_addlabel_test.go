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

func TestServer_AddLabel(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	task := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "This task depends on another",
		},
	})
	label := c.CreateLabelT(ctx, t, &taskspb.CreateLabelRequest{
		Label: &taskspb.Label{
			DisplayName: "some label",
		},
	})

	req := &taskspb.AddLabelRequest{
		Task:    task.GetName(),
		Label:   label.GetName(),
		Comment: "`label` was added.",
	}
	got, err := c.AddLabel(ctx, req)
	if err != nil {
		t.Errorf("AddLabel(%v) err = %v; want nil", req, err)
	}
	if diff := cmp.Diff(task, got, protocmp.Transform(), protocmp.IgnoreFields(task, "labels")); diff != "" {
		t.Errorf("unexpected result of AddLabel (-want +got)\n%s", diff)
	}

	gotLabels := got.GetLabels()
	wantLabels := []string{label.GetName()}
	if diff := cmp.Diff(wantLabels, gotLabels); diff != "" {
		t.Errorf("unexpected diff of task labels (-want +got)\n%s", diff)
	}
}

func TestServer_AddLabel_Errors(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	task := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "This task depends on another",
		},
	})
	label := c.CreateLabelT(ctx, t, &taskspb.CreateLabelRequest{
		Label: &taskspb.Label{
			DisplayName: "some label",
		},
	})

	for _, tt := range []struct {
		name string
		req  *taskspb.AddLabelRequest
		tf   errtest.TestFunc
	}{
		{
			name: "EmptyTask",
			req: &taskspb.AddLabelRequest{
				Task:  "",
				Label: label.GetName(),
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("empty task"),
			),
		},
		{
			name: "InvalidTask",
			req: &taskspb.AddLabelRequest{
				Task:  "invalid/123",
				Label: label.GetName(),
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("invalid task"),
				errtest.ErrorContains(`"tasks/{task}"`),
			),
		},
		{
			name: "MissingTask",
			req: &taskspb.AddLabelRequest{
				Task:  "tasks/999",
				Label: label.GetName(),
			},
			tf: errtest.All(
				grpctest.WantCode(codes.NotFound),
				errtest.ErrorContains("not found"),
				errtest.ErrorContains(`"tasks/999"`),
			),
		},
		{
			name: "EmptyLabel",
			req: &taskspb.AddLabelRequest{
				Task:  task.GetName(),
				Label: "",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("empty label"),
			),
		},
		{
			name: "InvalidLabel",
			req: &taskspb.AddLabelRequest{
				Task:  task.GetName(),
				Label: "invalid/456",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("invalid label"),
				errtest.ErrorContains(`"labels/{label}"`),
			),
		},
		{
			name: "MissingLabel",
			req: &taskspb.AddLabelRequest{
				Task:  task.GetName(),
				Label: "labels/999",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.NotFound),
				errtest.ErrorContains("not found"),
				errtest.ErrorContains(`"labels/999"`),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.AddLabel(ctx, tt.req)
			if err == nil {
				t.Errorf("AddLabel(%v) err = nil; want non-nil", tt.req)
			}
			tt.tf(t, err)
		})
	}
}

func TestServer_AddLabel_AlreadyExists(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	task := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "This task depends on another",
		},
	})
	label := c.CreateLabelT(ctx, t, &taskspb.CreateLabelRequest{
		Label: &taskspb.Label{
			DisplayName: "some label",
		},
	})
	task = c.AddLabelT(ctx, t, &taskspb.AddLabelRequest{
		Task:    task.GetName(),
		Label:   label.GetName(),
		Comment: "First time adding the label.",
	})

	req := &taskspb.AddLabelRequest{
		Task:    task.GetName(),
		Label:   label.GetName(),
		Comment: "Second time adding the label.",
	}
	_, err := c.AddLabel(ctx, req)
	if err == nil {
		t.Fatalf("AddLabel(%v) err = nil; want non-nil", req)
	}
	tf := errtest.All(
		grpctest.WantCode(codes.FailedPrecondition),
		errtest.ErrorContains("already has label"),
		errtest.ErrorContains(fmt.Sprintf("%q", task.GetName())),
		errtest.ErrorContains(fmt.Sprintf("%q", label.GetName())),
	)
	tf(t, err)
}
