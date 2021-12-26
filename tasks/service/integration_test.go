package service

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/testing/protocmp"

	"github.com/Saser/pdp/testing/grpctest"

	taskspb "github.com/Saser/pdp/tasks/tasks_go_proto"
)

func TestCreateGetAndList(t *testing.T) {
	ctx := context.Background()
	cc := grpctest.NewClientConnT(t, &taskspb.Tasks_ServiceDesc, New())
	c := taskspb.NewTasksClient(cc)

	// Create a simple task.
	var task *taskspb.Task
	{
		want := &taskspb.Task{
			Title:     "This is my task",
			Completed: false,
		}
		req := &taskspb.CreateTaskRequest{
			Task: want,
		}
		got, err := c.CreateTask(ctx, req)
		if err != nil {
			t.Fatalf("CreateTask(%v) err = %v; want nil", req, err)
		}
		if diff := cmp.Diff(want, got, protocmp.Transform(), protocmp.IgnoreFields(got, "name")); diff != "" {
			t.Fatalf("CreateTask: unexpected result (-want +got)\n%s", diff)
		}
		task = got
	}

	// Get the newly created task.
	{
		req := &taskspb.GetTaskRequest{Name: task.GetName()}
		got, err := c.GetTask(ctx, req)
		if err != nil {
			t.Fatalf("GetTask(%v) err = %v; want nil", req, err)
		}
		if diff := cmp.Diff(task, got, protocmp.Transform()); diff != "" {
			t.Errorf("GetTask: unexpected result (-want +got)\n%s", diff)
		}
	}

	// Look for the newly created task among all tasks.
	{
		req := &taskspb.ListTasksRequest{}
		res, err := c.ListTasks(ctx, req)
		if err != nil {
			t.Fatalf("ListTasks(%v) err = %v; want nil", req, err)
		}
		// There should be only one task -- the one we created.
		want := []*taskspb.Task{task}
		got := res.GetTasks()
		if diff := cmp.Diff(want, got, protocmp.Transform()); diff != "" {
			t.Errorf("ListTasks: unexpected result (-want +got)\n%s", diff)
		}
	}
}

func TestCreateAndListManyTasks(t *testing.T) {
	ctx := context.Background()
	cc := grpctest.NewClientConnT(t, &taskspb.Tasks_ServiceDesc, New())
	c := taskspb.NewTasksClient(cc)

	// Create a bunch of tasks.
	tasks := []*taskspb.Task{
		{Title: "Buy groceries"},
		{Title: "Cook pancakes"},
		{Title: "Nap on the couch"},
	}
	want := make([]*taskspb.Task, len(tasks))
	for i := range tasks {
		req := &taskspb.CreateTaskRequest{Task: tasks[i]}
		tt, err := c.CreateTask(ctx, req)
		if err != nil {
			t.Errorf("[i = %d] CreateTask(%v) err = %v; want nil", i, req, err)
		}
		want[i] = tt
	}
	if t.Failed() {
		t.FailNow()
	}

	// List all tasks -- it should be the same list as the one we created.
	// However, there is no ordering guarantee, so the returned list is allowed
	// to be out of order.
	req := &taskspb.ListTasksRequest{}
	res, err := c.ListTasks(ctx, req)
	if err != nil {
		t.Fatalf("ListTasks(%v) err = %v; want nil", req, err)
	}
	got := res.GetTasks()
	lessFunc := func(t1, t2 *taskspb.Task) bool {
		return t1.GetName() < t2.GetName()
	}
	if diff := cmp.Diff(want, got, protocmp.Transform(), cmpopts.SortSlices(lessFunc)); diff != "" {
		t.Errorf("ListTasks: unexpected result (-want +got)\n%s", diff)
	}
}
