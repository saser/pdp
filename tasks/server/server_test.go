package server

import (
	"context"
	"strings"
	"testing"

	"github.com/Saser/pdp/testing/grpctest"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/testing/protocmp"

	taskspb "github.com/Saser/pdp/tasks/tasks_go_proto"
)

const bufSize = 1024 * 1024

func setup(t *testing.T) taskspb.TasksClient {
	t.Helper()
	register := func(s *grpc.Server) { taskspb.RegisterTasksServer(s, New()) }
	cc := grpctest.NewClientConnT(t, register)
	return taskspb.NewTasksClient(cc)
}

func createTasks(ctx context.Context, t *testing.T, c taskspb.TasksClient, tasks []*taskspb.Task) []*taskspb.Task {
	t.Helper()
	created := make([]*taskspb.Task, len(tasks))
	for i, task := range tasks {
		req := &taskspb.CreateTaskRequest{
			Task: task,
		}
		task2, err := c.CreateTask(ctx, req)
		if err != nil {
			t.Errorf("CreateTask(%v) err = %v; want nil", req, err)
		}
		created[i] = task2
	}
	return created
}

func TestServer_CreateTask(t *testing.T) {
	c := setup(t)
	ctx := context.Background()
	createReq := &taskspb.CreateTaskRequest{
		Task: &taskspb.Task{
			Title: "My task",
		},
	}
	created, err := c.CreateTask(ctx, createReq)
	if err != nil {
		t.Errorf("c.CreateTask(%v) err = %v; want nil", createReq, err)
	}
	if got, prefix := created.GetName(), "tasks/"; !strings.HasPrefix(got, prefix) {
		t.Errorf("created.GetName() = %q; want prefix %q", got, prefix)
	}

	getReq := &taskspb.GetTaskRequest{
		Name: created.GetName(),
	}
	got, err := c.GetTask(ctx, getReq)
	if err != nil {
		t.Errorf("c.GetTask(%v) err = %v; want nil", getReq, err)
	}
	if diff := cmp.Diff(created, got, protocmp.Transform()); diff != "" {
		t.Errorf("diff between CreateTask() and GetTask() (-want +got)\n%s", diff)
	}
}

func TestServer_ListTasks(t *testing.T) {
	c := setup(t)
	ctx := context.Background()
	tasks := createTasks(ctx, t, c, []*taskspb.Task{
		{Title: "Some task"},
		{Title: "Some other task"},
	})
	req := &taskspb.ListTasksRequest{}
	res, err := c.ListTasks(ctx, req)
	if err != nil {
		t.Errorf("ListTasks(%v) err = %v; want nil", err)
	}
	less := func(t1, t2 *taskspb.Task) bool { return t1.GetName() < t2.GetName() }
	if diff := cmp.Diff(tasks, res.GetTasks(), protocmp.Transform(), cmpopts.SortSlices(less)); diff != "" {
		t.Errorf("diff between created tasks and ListTasks() (-want +got)\n%s", diff)
	}
}
