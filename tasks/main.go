package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	taskspb "github.com/Saser/pdp/genproto/tasks/v1"
)

type server struct {
	taskspb.UnimplementedTasksServer

	mu    sync.Mutex
	tasks map[string]*taskspb.Task
}

func newServer() *server {
	return &server{
		tasks: make(map[string]*taskspb.Task),
	}
}

func (s *server) GetTask(ctx context.Context, req *taskspb.GetTaskRequest) (*taskspb.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	task, ok := s.tasks[req.GetName()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "task not found: %q", req.GetName())
	}
	return task, nil
}

func (s *server) CreateTask(ctx context.Context, req *taskspb.CreateTaskRequest) (*taskspb.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	name := fmt.Sprintf("tasks/%d", len(s.tasks)+1)
	task := proto.Clone(req.GetTask()).(*taskspb.Task)
	task.Name = name
	s.tasks[name] = task
	return task, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Print(err)
		return
	}
	defer func() {
		if err := lis.Close(); err != nil {
			if errors.Is(err, net.ErrClosed) {
				// The listener was closed by the gRPC server.
				return
			}
			log.Print(err)
			return
		}
	}()

	s := newServer()
	gs := grpc.NewServer()
	taskspb.RegisterTasksServer(gs, s)
	reflection.Register(gs)

	ctx, reset := signal.NotifyContext(context.Background(), os.Interrupt)
	defer reset()

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return gs.Serve(lis)
	})

	<-ctx.Done()
	gs.GracefulStop()

	if err := g.Wait(); err != nil {
		log.Print(err)
		return
	}
}
