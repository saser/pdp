// Package pagetoken implements an offset-based page token suitable for use in List methods that support
// pagination, as described by https://google.aip.dev/132 and https://google.aip.dev/158. Page
// tokens returned by this package encodes the set of query parameters used for the List requests,
// and verifies such parameters for subsequent requests.
package pagetoken

import (
	"encoding/base64"
	"errors"
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	pagetokenpb "github.com/Saser/pdp/aip/pagetoken/page_token_go_proto"
)

// ListRequest defines the common interface of proto messages for List requests.
type ListRequest interface {
	proto.Message

	GetPageSize() int32
	GetPageToken() string
}

// PageToken implements an offset-based page token.
type PageToken struct {
	m *pagetokenpb.PageToken
}

// Parse returns a PageToken suitable for the given request. If req has an empty `page_token` field,
// Parse returns a PageToken that encodes an offset of 0 as well as the query parameters of the
// request. If the request has a non-empty `page_token` field, Parse reconstructs a PageToken from
// the `page_token` field, and verifies that it is valid for the request (i.e., the PageToken and
// the request have the same query parameters).
func Parse(req ListRequest) (*PageToken, error) {
	cleared, err := clearedAny(req)
	if err != nil {
		return nil, fmt.Errorf("parse page token: %w", err)
	}
	tok := req.GetPageToken()
	if tok == "" {
		return &PageToken{
			m: &pagetokenpb.PageToken{
				Offset:  0,
				Request: cleared,
			},
		}, nil
	}
	given, err := unmarshalString(tok)
	if err != nil {
		return nil, fmt.Errorf("parse page token: %w", err)
	}
	if !proto.Equal(given.m.GetRequest(), cleared) {
		return nil, errors.New("parse page token: list request differs")
	}
	return given, nil
}

// Offset returns the offset this PageToken encodes.
func (pt *PageToken) Offset() int32 {
	return pt.m.GetOffset()
}

// Next returns a new PageToken pt2 such that pt2 encodes the same query parameters as pt, and where
// pt2.Offset() == pt.Offset() + pageSize.
func (pt *PageToken) Next(pageSize int32) *PageToken {
	m2 := proto.Clone(pt.m).(*pagetokenpb.PageToken)
	m2.Offset += pageSize
	return &PageToken{
		m: m2,
	}
}

// String marshals this PageToken into a string that can then be unmarshaled to an equivalent
// PageToken. It is intended to be used for the `next_page_token` field in List response methods.
func (pt *PageToken) String() string {
	if pt.m.GetOffset() == 0 {
		return ""
	}
	s, err := marshalString(pt)
	if err != nil {
		panic(err)
	}
	return s
}

func marshalString(pt *PageToken) (string, error) {
	b, err := proto.Marshal(pt.m)
	if err != nil {
		return "", fmt.Errorf("marshal page token: %w", err)
	}
	s := base64.URLEncoding.EncodeToString(b)
	return s, nil
}

func unmarshalString(s string) (*PageToken, error) {
	b, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("unmarshal page token: %w", err)
	}
	m := &pagetokenpb.PageToken{}
	if err := proto.Unmarshal(b, m); err != nil {
		return nil, fmt.Errorf("unmarshal page token: %w", err)
	}
	return &PageToken{
		m: m,
	}, nil
}

func clearedAny(req ListRequest) (*anypb.Any, error) {
	cleared := proto.Clone(req)
	cleared.ProtoReflect().Clear(cleared.ProtoReflect().Descriptor().Fields().ByName("page_size"))
	cleared.ProtoReflect().Clear(cleared.ProtoReflect().Descriptor().Fields().ByName("page_token"))
	return anypb.New(cleared)
}
