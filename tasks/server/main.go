package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Saser/pdp/tasks/service"

	taskspb "github.com/Saser/pdp/tasks/tasks_go_proto"
)

var (
	port = flag.Int("port", 8080, "The port to listen on.")
)

func emain() error {
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	svc := service.New()
	srv := grpc.NewServer()
	reflection.Register(srv)
	taskspb.RegisterTasksServer(srv, svc)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		return err
	}
	defer lis.Close()

	errc := make(chan error, 1)
	go func() {
		errc <- srv.Serve(lis)
	}()
	<-ctx.Done()
	srv.GracefulStop()
	return <-errc
}

func main() {
	if err := emain(); err != nil {
		log.Fatal(err)
	}
}
