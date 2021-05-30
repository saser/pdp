package server

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/Saser/pdp/aip/resourcename"
	"github.com/Saser/pdp/wellknown/money"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	financepb "github.com/Saser/pdp/finance/finance_go_proto"
	moneypb "google.golang.org/genproto/googleapis/type/money"
)

var (
	internalError = status.Error(codes.Internal, "internal error")
)

type Server struct {
	financepb.UnimplementedFinanceServer

	mu sync.Mutex

	accounts     []*financepb.Account
	accountIndex map[string]int    // name -> index into `accounts`
	accountName  map[string]string // display name -> account name
}

func New() *Server {
	return &Server{
		accountIndex: make(map[string]int),
		accountName:  make(map[string]string),
	}
}

func (s *Server) CreateAccount(ctx context.Context, req *financepb.CreateAccountRequest) (*financepb.Account, error) {
	account := req.GetAccount()
	displayName := account.GetDisplayName()
	if displayName == "" {
		return nil, status.Error(codes.InvalidArgument, "empty display name")
	}
	if account.GetKind() == financepb.Account_KIND_UNSPECIFIED {
		return nil, status.Errorf(codes.InvalidArgument, "account kind cannot be %v", financepb.Account_KIND_UNSPECIFIED)
	}
	if err := money.Validate(account.GetStartingBalance()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid starting balance: %v", err)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.accountName[displayName]; ok {
		return nil, status.Errorf(codes.AlreadyExists, "duplicate display name: %q", displayName)
	}

	index := len(s.accounts)
	name, err := accountPattern.Render(resourcename.Values{"account": fmt.Sprint(index + 1)})
	if err != nil {
		log.Printf("CreateAccount failed to render account name: %v", err)
		return nil, internalError
	}

	stored := proto.Clone(account).(*financepb.Account)
	stored.Name = name
	stored.CurrentBalance = proto.Clone(stored.GetStartingBalance()).(*moneypb.Money)
	s.accounts = append(s.accounts, stored)
	s.accountIndex[name] = index
	s.accountName[displayName] = name
	return stored, nil
}
