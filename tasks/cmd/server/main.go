package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/Saser/pdp/tasks/server"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	taskspb "github.com/Saser/pdp/tasks/tasks_go_proto"
)

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

	s := server.New()
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
