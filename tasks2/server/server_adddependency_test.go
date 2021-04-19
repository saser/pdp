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

func TestServer_AddDependency(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	task := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "This task depends on another",
		},
	})
	dependency := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "This task is depended on",
		},
	})

	req := &taskspb.AddDependencyRequest{
		Task:       task.GetName(),
		Dependency: dependency.GetName(),
		Comment:    "`dependency` needs to be done first.",
	}
	got, err := c.AddDependency(ctx, req)
	if err != nil {
		t.Errorf("AddDependency(%v) err = %v; want nil", req, err)
	}
	if diff := cmp.Diff(task, got, protocmp.Transform(), protocmp.IgnoreFields(task, "dependencies")); diff != "" {
		t.Errorf("unexpected result of AddDependency (-want +got)\n%s", diff)
	}

	gotDependencies := got.GetDependencies()
	wantDependencies := []string{dependency.GetName()}
	if diff := cmp.Diff(wantDependencies, gotDependencies); diff != "" {
		t.Errorf("unexpected diff of task dependencies (-want +got)\n%s", diff)
	}
}

func TestServer_AddDependency_Errors(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	task := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "This task depends on another",
		},
	})
	dependency := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "This task is depended on",
		},
	})

	for _, tt := range []struct {
		name string
		req  *taskspb.AddDependencyRequest
		tf   errtest.TestFunc
	}{
		{
			name: "EmptyTask",
			req: &taskspb.AddDependencyRequest{
				Task:       "",
				Dependency: dependency.GetName(),
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("empty task"),
			),
		},
		{
			name: "InvalidTask",
			req: &taskspb.AddDependencyRequest{
				Task:       "invalid/123",
				Dependency: dependency.GetName(),
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("invalid task"),
				errtest.ErrorContains(`"tasks/{task}"`),
			),
		},
		{
			name: "MissingTask",
			req: &taskspb.AddDependencyRequest{
				Task:       "tasks/999",
				Dependency: dependency.GetName(),
			},
			tf: errtest.All(
				grpctest.WantCode(codes.NotFound),
				errtest.ErrorContains("not found"),
				errtest.ErrorContains(`"tasks/999"`),
			),
		},
		{
			name: "EmptyDependency",
			req: &taskspb.AddDependencyRequest{
				Task:       task.GetName(),
				Dependency: "",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("empty dependency"),
			),
		},
		{
			name: "InvalidDependency",
			req: &taskspb.AddDependencyRequest{
				Task:       task.GetName(),
				Dependency: "invalid/456",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("invalid dependency"),
				errtest.ErrorContains(`"tasks/{task}"`),
			),
		},
		{
			name: "MissingDependency",
			req: &taskspb.AddDependencyRequest{
				Task:       task.GetName(),
				Dependency: "tasks/999",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.NotFound),
				errtest.ErrorContains("not found"),
				errtest.ErrorContains(`"tasks/999"`),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.AddDependency(ctx, tt.req)
			if err == nil {
				t.Errorf("AddDependency(%v) err = nil; want non-nil", tt.req)
			}
			tt.tf(t, err)
		})
	}
}

func TestServer_AddDependency_AlreadyExists(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	task := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "This task depends on another",
		},
	})
	dependency := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "This task is depended on",
		},
	})
	task = c.AddDependencyT(ctx, t, &taskspb.AddDependencyRequest{
		Task:       task.GetName(),
		Dependency: dependency.GetName(),
		Comment:    "First time adding the dependency.",
	})

	req := &taskspb.AddDependencyRequest{
		Task:       task.GetName(),
		Dependency: dependency.GetName(),
		Comment:    "Second time adding the dependency.",
	}
	_, err := c.AddDependency(ctx, req)
	if err == nil {
		t.Fatalf("AddDependency(%v) err = nil; want non-nil", req)
	}
	tf := errtest.All(
		grpctest.WantCode(codes.AlreadyExists),
		errtest.ErrorContains("already depends"),
		errtest.ErrorContains(fmt.Sprintf("%q", task.GetName())),
		errtest.ErrorContains(fmt.Sprintf("%q", dependency.GetName())),
	)
	tf(t, err)
}
