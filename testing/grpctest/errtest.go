package grpctest

import (
	"testing"

	"github.com/Saser/pdp/testing/errtest"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func WantCode(code codes.Code) errtest.TestFunc {
	return func(tb testing.TB, err error) {
		tb.Helper()
		if got := status.Code(err); got != code {
			tb.Errorf("status.Code(%v) = %v; want %v", err, got, code)
		}
	}
}
