package pagetoken

import (
	"testing"

	pagetokenpb "github.com/Saser/pdp/aip/pagetoken/page_token_go_proto"
)

func TestParse_EmptyPageToken(t *testing.T) {
	for _, tt := range []struct {
		name string
		req  ListRequest
	}{
		{
			name: "DefaultFields",
			req:  &pagetokenpb.ExampleListRequest{},
		},
		{
			name: "NonDefaultPageSize",
			req: &pagetokenpb.ExampleListRequest{
				PageSize: 25,
			},
		},
		{
			name: "CustomParameter",
			req: &pagetokenpb.ExampleListRequest{
				Foo: "bar",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := Parse(tt.req); err != nil {
				t.Errorf("Parse(%v) err = %v; want nil", tt.req, err)
			}
		})
	}
}

func TestParse_NonEmptyPageToken_OK(t *testing.T) {
	req1 := &pagetokenpb.ExampleListRequest{
		PageSize: 10,
		Foo:      "bar",
	}
	pt1, err := Parse(req1)
	if err != nil {
		t.Fatalf("Parse(%v) err = %v; want nil", req1, err)
	}

	req2 := &pagetokenpb.ExampleListRequest{
		PageSize:  50,
		PageToken: pt1.Next(req1.GetPageSize()).String(),
		Foo:       "bar",
	}
	if _, err := Parse(req2); err != nil {
		t.Errorf("Parse(%v) err = %v; want nil", req2, err)
	}
}

func TestParse_NonEmptyPageToken_Errors(t *testing.T) {
	req1 := &pagetokenpb.ExampleListRequest{
		PageSize: 10,
		Foo:      "bar",
	}
	pt1, err := Parse(req1)
	if err != nil {
		t.Fatalf("Parse(%v) err = %v; want nil", req1, err)
	}

	req2 := &pagetokenpb.ExampleListRequest{
		PageSize:  50,
		PageToken: pt1.Next(req1.GetPageSize()).String(),
		Foo:       "barbarbar",
	}
	if _, err := Parse(req2); err == nil {
		t.Errorf("Parse(%v) err = nil; want non-nil", req2)
	}
}

func TestPageToken_Next_Offset(t *testing.T) {
	req := &pagetokenpb.ExampleListRequest{
		Foo: "bar",
	}
	pt1, err := Parse(req)
	if err != nil {
		t.Fatalf("Parse(%v) err = %v; want nil", req, err)
	}
	if got, want := pt1.Offset(), int32(0); got != want {
		t.Errorf("pt1.Offset() = %v; want %v", got, want)
	}
	pt2 := pt1.Next(10)
	if got, want := pt2.Offset(), int32(10); got != want {
		t.Errorf("pt2.Offset() = %v; want %v", got, want)
	}
}

func TestPageToken_Next_String(t *testing.T) {
	req := &pagetokenpb.ExampleListRequest{
		Foo: "bar",
	}
	pt1, err := Parse(req)
	if err != nil {
		t.Fatalf("Parse(%v) err = %v; want nil", req, err)
	}
	if got, want := pt1.String(), ""; got != want {
		t.Errorf("pt1.String() = %v; want %v", got, want)
	}
	pt2 := pt1.Next(10)
	if pt2.String() == "" {
		t.Errorf("pt2.String() = %q; want non-empty", "")
	}
}
