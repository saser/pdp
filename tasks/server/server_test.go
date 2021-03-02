package server

import (
	"context"
	"errors"
	"net"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
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
