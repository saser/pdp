package server

import (
	"context"
	"testing"

	"github.com/Saser/pdp/testing/grpctest"
	"google.golang.org/grpc"

	financepb "github.com/Saser/pdp/finance/finance_go_proto"
)

type testFinanceClient struct {
	financepb.FinanceClient
}

func setup(t *testing.T) testFinanceClient {
	t.Helper()
	cc := grpctest.NewClientConnT(t, &financepb.Finance_ServiceDesc, New())
	return testFinanceClient{FinanceClient: financepb.NewFinanceClient(cc)}
}

func (c testFinanceClient) CreateAccountT(ctx context.Context, t *testing.T, in *financepb.CreateAccountRequest, opts ...grpc.CallOption) *financepb.Account {
	t.Helper()
	account, err := c.CreateAccount(ctx, in, opts...)
	if err != nil {
		t.Fatalf("CreateAccount(%v) err = %v; want nil", in, err)
	}
	return account
}
