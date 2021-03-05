package grpctest

import (
	"context"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

func NewClientConnT(t testing.TB, register func(*grpc.Server)) *grpc.ClientConn {
	lis := bufconn.Listen(bufSize)
	t.Cleanup(func() {
		if err := lis.Close(); err != nil && err != net.ErrClosed {
			t.Errorf("lis.Close() = %v; want nil", err)
			return
		}
	})

	s := grpc.NewServer()
	register(s)
	go func() {
		if err := s.Serve(lis); err != nil {
			t.Errorf("s.Serve() = %v; want nil", err)
			return
		}
	}()
	t.Cleanup(s.GracefulStop)
	dialer := func(context.Context, string) (net.Conn, error) { return lis.Dial() }
	cc, err := grpc.Dial("bufconn", grpc.WithInsecure(), grpc.WithContextDialer(dialer))
	if err != nil {
		t.Fatalf("grpc.Dial() err = %v; want nil", err)
	}
	return cc
}
