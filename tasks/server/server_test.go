package server

import (
	"context"
	"errors"
	"net"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/testing/protocmp"

	taskspb "github.com/Saser/pdp/tasks/tasks_go_proto"
)

const bufSize = 1024 * 1024

func setup(t *testing.T) taskspb.TasksClient {
	lis := bufconn.Listen(bufSize)
	t.Cleanup(func() {
		if err := lis.Close(); err != nil && errors.Is(err, net.ErrClosed) {
			t.Errorf("lis.Close() = %v; want nil", err)
		}
	})

	gs := grpc.NewServer()
	t.Cleanup(gs.GracefulStop)
	s := New()
	taskspb.RegisterTasksServer(gs, s)
	go func() {
		if err := gs.Serve(lis); err != nil {
			t.Errorf("gs.Serve() = %v; want nil", err)
		}
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}
	cc, err := grpc.Dial("bufconn", grpc.WithInsecure(), grpc.WithContextDialer(dialer))
	if err != nil {
		t.Fatalf("grpc.Dial() = %v; want nil", err)
	}
	t.Cleanup(func() {
		if err := cc.Close(); err != nil {
			t.Errorf("cc.Close() = %v; want nil", err)
		}
	})
	return taskspb.NewTasksClient(cc)
}

func createTasks(ctx context.Context, t *testing.T, c taskspb.TasksClient, tasks []*taskspb.Task) []*taskspb.Task {
	created := make([]*taskspb.Task, len(tasks))
	for i, task := range tasks {
		req := &taskspb.CreateTaskRequest{
			Task: task,
		}
		task2, err := c.CreateTask(ctx, req)
		if err != nil {
			t.Errorf("CreateTask(%v) err = %v", req, err)
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
