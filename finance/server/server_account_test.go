package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/Saser/pdp/testing/errtest"
	"github.com/Saser/pdp/testing/grpctest"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/genproto/googleapis/type/money"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"

	financepb "github.com/Saser/pdp/finance/finance_go_proto"
)

func TestServer_CreateAccount(t *testing.T) {
	ctx := context.Background()
	c := setup(t)

	req := &financepb.CreateAccountRequest{
		Account: &financepb.Account{
			DisplayName: "Bank account",
			Kind:        financepb.Account_DEBIT,
			StartingBalance: &money.Money{
				CurrencyCode: "SEK",
				Units:        100,
			},
		},
	}
	got, err := c.CreateAccount(ctx, req)
	if err != nil {
		t.Fatalf("CreateAccount(%v) err = %v; want nil", req, err)
	}
	if got.GetName() == "" {
		t.Errorf("CreateAccount(%v) account.GetName() is empty; want non-empty", req)
	}
	want := proto.Clone(req.GetAccount()).(*financepb.Account)
	want.CurrentBalance = proto.Clone(want.GetStartingBalance()).(*money.Money)
	if diff := cmp.Diff(want, got, protocmp.Transform(), protocmp.IgnoreFields(&financepb.Account{}, "name")); diff != "" {
		t.Errorf("unexpected diff in CreateAccount (-want +got)\n%s", diff)
	}
}

func TestServer_CreateAccount_Errors(t *testing.T) {
	ctx := context.Background()
	c := setup(t)
	existing := c.CreateAccountT(ctx, t, &financepb.CreateAccountRequest{
		Account: &financepb.Account{
			DisplayName: "Existing account",
			Kind:        financepb.Account_DEBIT,
			StartingBalance: &money.Money{
				CurrencyCode: "SEK",
				Units:        100,
			},
		},
	})
	for _, tt := range []struct {
		name string
		req  *financepb.CreateAccountRequest
		tf   errtest.TestFunc
	}{
		{
			name: "EmptyDisplayName",
			req: &financepb.CreateAccountRequest{
				Account: &financepb.Account{
					DisplayName: "",
					Kind:        financepb.Account_DEBIT,
					StartingBalance: &money.Money{
						CurrencyCode: "SEK",
						Units:        100,
					},
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("empty display name"),
			),
		},
		{
			name: "UnspecifiedKind",
			req: &financepb.CreateAccountRequest{
				Account: &financepb.Account{
					DisplayName: "Another account",
					Kind:        financepb.Account_KIND_UNSPECIFIED,
					StartingBalance: &money.Money{
						CurrencyCode: "SEK",
						Units:        100,
					},
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("kind cannot be KIND_UNSPECIFIED"),
			),
		},
		{
			name: "InvalidStartingBalance_CurrencyCode",
			req: &financepb.CreateAccountRequest{
				Account: &financepb.Account{
					DisplayName: "Another account",
					Kind:        financepb.Account_DEBIT,
					StartingBalance: &money.Money{
						CurrencyCode: "USD",
						Units:        100,
					},
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains(`invalid currency code "USD"`),
				errtest.ErrorContains(`want "SEK"`),
			),
		},
		{
			name: "InvalidStartingBalance_Amount",
			req: &financepb.CreateAccountRequest{
				Account: &financepb.Account{
					DisplayName: "Another account",
					Kind:        financepb.Account_DEBIT,
					StartingBalance: &money.Money{
						CurrencyCode: "SEK",
						Units:        +100,
						Nanos:        -100,
					},
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.InvalidArgument),
				errtest.ErrorContains("mismatched signs"),
			),
		},
		{
			name: "DuplicateDisplayName",
			req: &financepb.CreateAccountRequest{
				Account: &financepb.Account{
					DisplayName: existing.GetDisplayName(),
					Kind:        financepb.Account_DEBIT,
					StartingBalance: &money.Money{
						CurrencyCode: "SEK",
						Units:        100,
					},
				},
			},
			tf: errtest.All(
				grpctest.WantCode(codes.AlreadyExists),
				errtest.ErrorContains("duplicate display name"),
				errtest.ErrorContains(fmt.Sprintf("%q", existing.GetDisplayName())),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.CreateAccount(ctx, tt.req)
			if err == nil {
				t.Fatalf("CreateAccount(%v) err = nil; want non-nil", tt.req)
			}
			tt.tf(t, err)
		})
	}
}
