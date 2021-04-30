package server

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"

	taskspb "github.com/Saser/pdp/tasks2/tasks_go_proto"
)

func checkEvents(t *testing.T, events []*taskspb.Event) {
	t.Helper()
	for _, event := range events {
		name := event.GetName()
		if _, err := eventPattern.Match(name); err != nil {
			t.Errorf("eventPattern.Match(%q) err = %v; want nil", name, err)
		}
		createTime := event.GetCreateTime()
		if createTime == nil {
			t.Error("GetCreateTime() = nil; want non-nil")
		}
		if err := createTime.CheckValid(); err != nil {
			t.Errorf("GetCreateTime().CheckValid() = %v; want nil", err)
		}
	}
}

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
		checkEvents(t, got)

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

func TestServer_ListEvents_RemoveDependency(t *testing.T) {
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
		Comment:    "[AddDependency] This should show up in both `task`s events and `dependency`s events.",
	})
	c.RemoveDependencyT(ctx, t, &taskspb.RemoveDependencyRequest{
		Task:       task.GetName(),
		Dependency: dependency.GetName(),
		Comment:    "[RemoveDependency] This should show up in both `task`s events and `dependency`s events.",
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
		checkEvents(t, got)

		want := []*taskspb.Event{
			{
				Parent:  parent,
				Comment: "[AddDependency] This should show up in both `task`s events and `dependency`s events.",
				Kind: &taskspb.Event_AddDependency_{AddDependency: &taskspb.Event_AddDependency{
					Task:       task.GetName(),
					Dependency: dependency.GetName(),
				}},
			},
			{
				Parent:  parent,
				Comment: "[RemoveDependency] This should show up in both `task`s events and `dependency`s events.",
				Kind: &taskspb.Event_RemoveDependency_{RemoveDependency: &taskspb.Event_RemoveDependency{
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

func TestServer_ListEvents_AddLabel(t *testing.T) {
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
	c.AddLabelT(ctx, t, &taskspb.AddLabelRequest{
		Task:    task.GetName(),
		Label:   label.GetName(),
		Comment: "This should show up in `task`s events.",
	})

	req := &taskspb.ListEventsRequest{
		Parent: task.GetName(),
	}
	res, err := c.ListEvents(ctx, req)
	if err != nil {
		t.Errorf("ListEvents(%v) err = %v; want nil", req, err)
	}

	got := res.GetEvents()
	checkEvents(t, got)

	want := []*taskspb.Event{
		{
			Parent:  task.GetName(),
			Comment: "This should show up in `task`s events.",
			Kind: &taskspb.Event_AddLabel_{AddLabel: &taskspb.Event_AddLabel{
				Label: label.GetName(),
			}},
		},
	}
	if diff := cmp.Diff(want, got, protocmp.Transform(), protocmp.IgnoreFields(&taskspb.Event{}, "name", "create_time")); diff != "" {
		t.Errorf("unexpected diff between events (-want +got)\n%s", diff)
	}
}

func TestServer_ListEvents_RemoveLabel(t *testing.T) {
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
	c.AddLabelT(ctx, t, &taskspb.AddLabelRequest{
		Task:    task.GetName(),
		Label:   label.GetName(),
		Comment: "[AddLabel] This should show up in `task`s events.",
	})
	c.RemoveLabelT(ctx, t, &taskspb.RemoveLabelRequest{
		Task:    task.GetName(),
		Label:   label.GetName(),
		Comment: "[RemoveLabel] This should show up in `task`s events.",
	})

	req := &taskspb.ListEventsRequest{
		Parent: task.GetName(),
	}
	res, err := c.ListEvents(ctx, req)
	if err != nil {
		t.Errorf("ListEvents(%v) err = %v; want nil", req, err)
	}

	got := res.GetEvents()
	checkEvents(t, got)

	want := []*taskspb.Event{
		{
			Parent:  task.GetName(),
			Comment: "[AddLabel] This should show up in `task`s events.",
			Kind: &taskspb.Event_AddLabel_{AddLabel: &taskspb.Event_AddLabel{
				Label: label.GetName(),
			}},
		},
		{
			Parent:  task.GetName(),
			Comment: "[RemoveLabel] This should show up in `task`s events.",
			Kind: &taskspb.Event_RemoveLabel_{RemoveLabel: &taskspb.Event_RemoveLabel{
				Label: label.GetName(),
			}},
		},
	}
	if diff := cmp.Diff(want, got, protocmp.Transform(), protocmp.IgnoreFields(&taskspb.Event{}, "name", "create_time")); diff != "" {
		t.Errorf("unexpected diff between events (-want +got)\n%s", diff)
	}
}

func TestServer_ListEvents_CompleteTask(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	task := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "Some task",
		},
	})
	task = c.CompleteTaskT(ctx, t, &taskspb.CompleteTaskRequest{
		Task:    task.GetName(),
		Comment: "Completed task",
	})

	req := &taskspb.ListEventsRequest{
		Parent: task.GetName(),
	}
	res, err := c.ListEvents(ctx, req)
	if err != nil {
		t.Errorf("ListEvents(%v) err = %v; want nil", req, err)
	}

	got := res.GetEvents()
	checkEvents(t, got)

	want := []*taskspb.Event{
		{
			Parent:  task.GetName(),
			Comment: "Completed task",
			Kind:    &taskspb.Event_Complete_{Complete: &taskspb.Event_Complete{}},
		},
	}
	if diff := cmp.Diff(want, got, protocmp.Transform(), protocmp.IgnoreFields(&taskspb.Event{}, "name", "create_time")); diff != "" {
		t.Errorf("unexpected diff between events (-want +got)\n%s", diff)
	}
}

func TestServer_ListEvents_UncompleteTask(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	task := c.CreateTaskT(ctx, t, &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "Some task",
		},
	})
	task = c.CompleteTaskT(ctx, t, &taskspb.CompleteTaskRequest{
		Task:    task.GetName(),
		Comment: "Completed task",
	})
	task = c.UncompleteTaskT(ctx, t, &taskspb.UncompleteTaskRequest{
		Task:    task.GetName(),
		Comment: "Uncompleted task",
	})

	req := &taskspb.ListEventsRequest{
		Parent: task.GetName(),
	}
	res, err := c.ListEvents(ctx, req)
	if err != nil {
		t.Errorf("ListEvents(%v) err = %v; want nil", req, err)
	}

	got := res.GetEvents()
	checkEvents(t, got)

	want := []*taskspb.Event{
		{
			Parent:  task.GetName(),
			Comment: "Completed task",
			Kind:    &taskspb.Event_Complete_{Complete: &taskspb.Event_Complete{}},
		},
		{
			Parent:  task.GetName(),
			Comment: "Uncompleted task",
			Kind:    &taskspb.Event_Uncomplete_{Uncomplete: &taskspb.Event_Uncomplete{}},
		},
	}
	if diff := cmp.Diff(want, got, protocmp.Transform(), protocmp.IgnoreFields(&taskspb.Event{}, "name", "create_time")); diff != "" {
		t.Errorf("unexpected diff between events (-want +got)\n%s", diff)
	}
}
