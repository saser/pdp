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

func NewClientConnT(t testing.TB, register func(*grpc.Server)) *grpc.ClientConn {
	t.Helper()

	lis := bufconn.Listen(bufSize)
	t.Cleanup(func() {
		if err := lis.Close(); err != nil && err != net.ErrClosed {
			t.Errorf("lis.Close() = %v; want nil", err)
			return
		}
	})

	var g errgroup.Group
	t.Cleanup(func() {
		if err := g.Wait(); err != nil {
			t.Errorf("g.Wait() = %v; want nil", err)
		}
	})

	s := grpc.NewServer()
	register(s)
	g.Go(func() error {
		return s.Serve(lis)
	})
	t.Cleanup(s.GracefulStop)

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}
	cc, err := grpc.Dial("bufconn", grpc.WithInsecure(), grpc.WithContextDialer(dialer), grpc.WithBlock())
	if err != nil {
		t.Fatalf("grpc.Dial() err = %v; want nil", err)
	}
	t.Cleanup(func() {
		if err := cc.Close(); err != nil {
			t.Errorf("cc.Close() = %v; want nil", err)
		}
	})

	return cc
}
