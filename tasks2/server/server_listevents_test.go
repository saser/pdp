package server

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"

	taskspb "github.com/Saser/pdp/tasks2/tasks_go_proto"
)

func TestServer_ListEvents_AddDependency(t *testing.T) {
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
	c.AddDependencyT(ctx, t, &taskspb.AddDependencyRequest{
		Task:       task.GetName(),
		Dependency: dependency.GetName(),
		Comment:    "This should show up in both `task`s events and `dependency`s events.",
	})

	for _, parent := range []string{
		task.GetName(),
		dependency.GetName(),
	} {
		req := &taskspb.ListEventsRequest{
			Parent: parent,
		}
		res, err := c.ListEvents(ctx, req)
		if err != nil {
			t.Errorf("ListEvents(%v) err = %v; want nil", req, err)
		}
		got := res.GetEvents()
		want := []*taskspb.Event{
			{
				Parent:  parent,
				Comment: "This should show up in both `task`s events and `dependency`s events.",
				Kind: &taskspb.Event_AddDependency_{AddDependency: &taskspb.Event_AddDependency{
					Task:       task.GetName(),
					Dependency: dependency.GetName(),
				}},
			},
		}
		if diff := cmp.Diff(want, got, protocmp.Transform(), protocmp.IgnoreFields(&taskspb.Event{}, "name", "create_time")); diff != "" {
			t.Errorf("unexpected diff between events (-want +got)\n%s", diff)
		}
	}
}
