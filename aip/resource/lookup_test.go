package resource

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/testing/protocmp"

	testingpb "github.com/Saser/pdp/aip/resource/internal/testing/testing_go_proto"
)

func TestLookupMessage(t *testing.T) {
	for _, tt := range []struct {
		name       string
		typeString string
		wantMD     protoreflect.MessageDescriptor
		wantOK     bool
	}{
		{
			name:       "Publisher",
			typeString: "type.api.saser.se/aip.resource.internal.testing.Publisher",
			wantMD:     (&testingpb.Publisher{}).ProtoReflect().Descriptor(),
			wantOK:     true,
		},
		{
			name:       "Book",
			typeString: "type.api.saser.se/aip.resource.internal.testing.Book",
			wantMD:     (&testingpb.Book{}).ProtoReflect().Descriptor(),
			wantOK:     true,
		},
		{
			name:       "Invalid",
			typeString: "type.api.saser.se/aip.resource.internal.testing.Invalid",
			wantMD:     nil,
			wantOK:     false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			gotMD, gotOK := LookupMessage(tt.typeString)
			if gotOK != tt.wantOK {
				t.Errorf("LookupMessage(%q) ok = %v; want %v", tt.typeString, gotOK, tt.wantOK)
			}
			if gotMD != tt.wantMD {
				t.Errorf("LookupMessage(%q) md = %v; want %v", tt.typeString, gotMD, tt.wantMD)
			}
		})
	}
}

func TestLookupResource(t *testing.T) {
	for _, tt := range []struct {
		name       string
		typeString string
		wantRD     *annotations.ResourceDescriptor
		wantOK     bool
	}{
		{
			name:       "Publisher",
			typeString: "type.api.saser.se/aip.resource.internal.testing.Publisher",
			wantRD: &annotations.ResourceDescriptor{
				Type:    "type.api.saser.se/aip.resource.internal.testing.Publisher",
				Pattern: []string{"publishers/{publisher}"},
			},
			wantOK: true,
		},
		{
			name:       "Book",
			typeString: "type.api.saser.se/aip.resource.internal.testing.Book",
			wantRD: &annotations.ResourceDescriptor{
				Type:    "type.api.saser.se/aip.resource.internal.testing.Book",
				Pattern: []string{"publishers/{publisher}/books/{book}"},
			},
			wantOK: true,
		},
		{
			name:       "Invalid",
			typeString: "type.api.saser.se/aip.resource.internal.testing.Invalid",
			wantRD:     nil,
			wantOK:     false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			gotRD, gotOK := LookupResource(tt.typeString)
			if gotOK != tt.wantOK {
				t.Errorf("LookupMessage(%q) ok = %v; want %v", tt.typeString, gotOK, tt.wantOK)
			}
			if diff := cmp.Diff(tt.wantRD, gotRD, protocmp.Transform()); diff != "" {
				t.Errorf("diff between resource descriptors (-want +got)\n%s", diff)
			}
		})
	}
}
