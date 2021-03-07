package server

import (
	"testing"

	"github.com/Saser/pdp/testing/grpctest"
	"google.golang.org/grpc"

	taskspb "github.com/Saser/pdp/tasks/tasks_go_proto"
)

func setup(t *testing.T) taskspb.TasksClient {
	t.Helper()
	register := func(s *grpc.Server) { taskspb.RegisterTasksServer(s, New()) }
	cc := grpctest.NewClientConnT(t, register)
	t.Cleanup(func() {
		if err := cc.Close(); err != nil {
			t.Errorf("cc.Close() = %v; want nil", err)
		}
	})
	return taskspb.NewTasksClient(cc)
}

func Test_setup(t *testing.T) {
	_ = setup(t)
}
