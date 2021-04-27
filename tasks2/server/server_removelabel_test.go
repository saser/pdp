package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/Saser/pdp/testing/errtest"
	"github.com/Saser/pdp/testing/grpctest"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/testing/protocmp"

	taskspb "github.com/Saser/pdp/tasks2/tasks_go_proto"
)

func TestServer_RemoveLabel(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	task := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "some task",
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
		Comment: fmt.Sprintf("AddLabel for label %q", label.GetName()),
	})

	req := &taskspb.RemoveLabelRequest{
		Task:    task.GetName(),
		Label:   label.GetName(),
		Comment: fmt.Sprintf("RemoveLabel for label %q", label.GetName()),
	}
	got, err := c.RemoveLabel(ctx, req)
	if err != nil {
		t.Errorf("RemoveLabel(%v) err = %v; want nil", req, err)
	}
	if diff := cmp.Diff(task, got, protocmp.Transform(), protocmp.IgnoreFields(task, "labels")); diff != "" {
		t.Errorf("unexpected result of RemoveLabel (-want +got)\n%s", diff)
	}

	gotLabels := got.GetLabels()
	wantLabels := []string{}
	if diff := cmp.Diff(wantLabels, gotLabels, cmpopts.EquateEmpty()); diff != "" {
		t.Errorf("unexpected diff of task labels (-want +got)\n%s", diff)
	}
}

func TestServer_RemoveLabel_Errors(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	task := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "Some task",
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
		Comment: fmt.Sprintf("AddLabel of label %q", label.GetName()),
	})

	unrelated := c.CreateLabelT(ctx, t, &taskspb.CreateLabelRequest{
		Label: &taskspb.Label{
			DisplayName: "an unrelated label",
		},
	})

	for _, tt := range []struct {
		name string
		req  *taskspb.RemoveLabelRequest
		tf   errtest.TestFunc
	}{
		{
			name: "EmptyTask",
			req: &taskspb.RemoveLabelRequest{
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
			req: &taskspb.RemoveLabelRequest{
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
			req: &taskspb.RemoveLabelRequest{
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
			req: &taskspb.RemoveLabelRequest{
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
			req: &taskspb.RemoveLabelRequest{
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
			req: &taskspb.RemoveLabelRequest{
				Task:  task.GetName(),
				Label: "labels/999",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.NotFound),
				errtest.ErrorContains("not found"),
				errtest.ErrorContains(`"labels/999"`),
			),
		},
		{
			name: "UnrelatedLabel",
			req: &taskspb.RemoveLabelRequest{
				Task:  task.GetName(),
				Label: unrelated.GetName(),
			},
			tf: errtest.All(
				grpctest.WantCode(codes.FailedPrecondition),
				errtest.ErrorContains("does not have label"),
				errtest.ErrorContains(fmt.Sprintf("%q", task.GetName())),
				errtest.ErrorContains(fmt.Sprintf("%q", unrelated.GetName())),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.RemoveLabel(ctx, tt.req)
			if err == nil {
				t.Fatalf("RemoveLabel(%v) err = nil; want non-nil", tt.req)
			}
			tt.tf(t, err)
		})
	}
}

func TestServer_RemoveLabel_RemovedTwice(t *testing.T) {
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
		Comment: "Added a label",
	})
	task = c.RemoveLabelT(ctx, t, &taskspb.RemoveLabelRequest{
		Task:    task.GetName(),
		Label:   label.GetName(),
		Comment: "Removed label",
	})

	req := &taskspb.RemoveLabelRequest{
		Task:    task.GetName(),
		Label:   label.GetName(),
		Comment: "Removed label again",
	}
	_, err := c.RemoveLabel(ctx, req)
	if err == nil {
		t.Errorf("RemoveLabel(%v) err = nil; want non-nil", req)
	}
	tf := errtest.All(
		grpctest.WantCode(codes.FailedPrecondition),
		errtest.ErrorContains("does not have label"),
		errtest.ErrorContains(fmt.Sprintf("%q", task.GetName())),
		errtest.ErrorContains(fmt.Sprintf("%q", label.GetName())),
	)
	tf(t, err)
}
