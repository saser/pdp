package resource

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"

	testingpb "github.com/Saser/pdp/aip/resource/internal/testing/testing_go_proto"
)

func TestDescriptor(t *testing.T) {
	for _, tt := range []struct {
		name string
		m    proto.Message
		want *annotations.ResourceDescriptor
	}{
		{
			name: "Publisher",
			m:    &testingpb.Publisher{},
			want: &annotations.ResourceDescriptor{
				Type:    "testing.internal.resource.aip.api.saser.se/Publisher",
				Pattern: []string{"publishers/{publisher}"},
			},
		},
		{
			name: "Book",
			m:    &testingpb.Book{},
			want: &annotations.ResourceDescriptor{
				Type:    "testing.internal.resource.aip.api.saser.se/Book",
				Pattern: []string{"publishers/{publisher}/books/{book}"},
			},
		},
		{
			name: "Invalid",
			m:    &testingpb.Invalid{},
			want: nil,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := Descriptor(tt.m)
			if diff := cmp.Diff(tt.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("diff between resource descriptors (-want +got)\n%s", diff)
			}
		})
	}
}
