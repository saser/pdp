package grpctest

import (
	"context"
	"net"
	"testing"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

func NewClientConnT(tb testing.TB, sd *grpc.ServiceDesc, ss interface{}) *grpc.ClientConn {
	tb.Helper()

	lis := bufconn.Listen(bufSize)
	tb.Cleanup(func() {
		if err := lis.Close(); err != nil && err != net.ErrClosed {
			tb.Errorf("lis.Close() = %v; want nil", err)
			return
		}
	})

	var g errgroup.Group
	tb.Cleanup(func() {
		if err := g.Wait(); err != nil {
			tb.Errorf("g.Wait() = %v; want nil", err)
		}
	})

	s := grpc.NewServer()
	s.RegisterService(sd, ss)
	g.Go(func() error {
		return s.Serve(lis)
	})
	tb.Cleanup(s.GracefulStop)

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}
	cc, err := grpc.Dial("bufconn", grpc.WithInsecure(), grpc.WithContextDialer(dialer), grpc.WithBlock())
	if err != nil {
		tb.Fatalf("grpc.Dial() err = %v; want nil", err)
	}
	tb.Cleanup(func() {
		if err := cc.Close(); err != nil {
			tb.Errorf("cc.Close() = %v; want nil", err)
		}
	})

	return cc
}
