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

func TestServer_RemoveDependency(t *testing.T) {
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
		Comment:    "AddDependency shows up for both `task` and `dependency`",
	})

	req := &taskspb.RemoveDependencyRequest{
		Task:       task.GetName(),
		Dependency: dependency.GetName(),
		Comment:    "RemoveDependency shows up for both `task` and `dependency`",
	}
	got, err := c.RemoveDependency(ctx, req)
	if err != nil {
		t.Errorf("RemoveDependency(%v) err = %v; want nil", req, err)
	}
	if diff := cmp.Diff(task, got, protocmp.Transform(), protocmp.IgnoreFields(task, "dependencies")); diff != "" {
		t.Errorf("unexpected result of RemoveDependency (-want +got)\n%s", diff)
	}

	gotDependencies := got.GetDependencies()
	wantDependencies := []string{}
	if diff := cmp.Diff(wantDependencies, gotDependencies, cmpopts.EquateEmpty()); diff != "" {
		t.Errorf("unexpected diff of task dependencies (-want +got)\n%s", diff)
	}
}

func TestServer_RemoveDependency_Errors(t *testing.T) {
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
		Comment:    "AddDependency shows up for both `task` and `dependency`",
	})

	independent := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "This task is independent of the others",
		},
	})

	for _, tt := range []struct {
		name string
		req  *taskspb.RemoveDependencyRequest
		tf   errtest.TestFunc
	}{
		{
			name: "EmptyTask",
			req: &taskspb.RemoveDependencyRequest{
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
			req: &taskspb.RemoveDependencyRequest{
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
			req: &taskspb.RemoveDependencyRequest{
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
			req: &taskspb.RemoveDependencyRequest{
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
			req: &taskspb.RemoveDependencyRequest{
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
			req: &taskspb.RemoveDependencyRequest{
				Task:       task.GetName(),
				Dependency: "tasks/999",
			},
			tf: errtest.All(
				grpctest.WantCode(codes.NotFound),
				errtest.ErrorContains("not found"),
				errtest.ErrorContains(`"tasks/999"`),
			),
		},
		{
			name: "FlippedArguments",
			req: &taskspb.RemoveDependencyRequest{
				Task:       dependency.GetName(),
				Dependency: task.GetName(),
			},
			tf: errtest.All(
				grpctest.WantCode(codes.FailedPrecondition),
				errtest.ErrorContains("no dependency exists"),
				errtest.ErrorContains(fmt.Sprintf("%q", task.GetName())),
				errtest.ErrorContains(fmt.Sprintf("%q", dependency.GetName())),
			),
		},
		{
			name: "IndependentTask",
			req: &taskspb.RemoveDependencyRequest{
				Task:       independent.GetName(),
				Dependency: dependency.GetName(),
			},
			tf: errtest.All(
				grpctest.WantCode(codes.FailedPrecondition),
				errtest.ErrorContains("no dependency exists"),
				errtest.ErrorContains(fmt.Sprintf("%q", independent.GetName())),
				errtest.ErrorContains(fmt.Sprintf("%q", dependency.GetName())),
			),
		},
		{
			name: "IndependentDependency",
			req: &taskspb.RemoveDependencyRequest{
				Task:       task.GetName(),
				Dependency: independent.GetName(),
			},
			tf: errtest.All(
				grpctest.WantCode(codes.FailedPrecondition),
				errtest.ErrorContains("no dependency exists"),
				errtest.ErrorContains(fmt.Sprintf("%q", task.GetName())),
				errtest.ErrorContains(fmt.Sprintf("%q", independent.GetName())),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.RemoveDependency(ctx, tt.req)
			if err == nil {
				t.Fatalf("RemoveDependency(%v) err = nil; want non-nil", tt.req)
			}
			tt.tf(t, err)
		})
	}
}

func TestServer_RemoveDependency_RemovedTwice(t *testing.T) {
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
		Comment:    "Added a dependency",
	})
	task = c.RemoveDependencyT(ctx, t, &taskspb.RemoveDependencyRequest{
		Task:       task.GetName(),
		Dependency: dependency.GetName(),
		Comment:    "Removed dependency",
	})

	req := &taskspb.RemoveDependencyRequest{
		Task:       task.GetName(),
		Dependency: dependency.GetName(),
		Comment:    "Removed dependency again",
	}
	_, err := c.RemoveDependency(ctx, req)
	if err == nil {
		t.Errorf("RemoveDependency(%v) err = nil; want non-nil", req)
	}
	tf := errtest.All(
		grpctest.WantCode(codes.FailedPrecondition),
		errtest.ErrorContains("no dependency exists"),
		errtest.ErrorContains(fmt.Sprintf("%q", task.GetName())),
		errtest.ErrorContains(fmt.Sprintf("%q", dependency.GetName())),
	)
	tf(t, err)
}
