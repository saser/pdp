package fieldmask

import (
	"testing"

	"github.com/Saser/pdp/testing/errtest"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	testingpb "github.com/Saser/pdp/aip/fieldmask/internal/testing/testing_go_proto"
)

func TestValidate(t *testing.T) {
	for _, tt := range []struct {
		name string
		m    proto.Message
		mask *fieldmaskpb.FieldMask
		tf   errtest.TestFunc
	}{
		{
			name: "SingleWildcard",
			m:    &testingpb.Test{},
			mask: &fieldmaskpb.FieldMask{Paths: []string{"*"}},
			tf:   errtest.IsNil(),
		},
		{
			name: "AllPaths",
			m:    &testingpb.Test{},
			mask: &fieldmaskpb.FieldMask{Paths: []string{
				"s",
				"rep_s",
				"nested",
				"rep_nested",
				"oo_s",
				"oo_nested",
			}},
			tf: errtest.IsNil(),
		},
		{
			name: "InvalidFields",
			m:    &testingpb.Test{},
			mask: &fieldmaskpb.FieldMask{Paths: []string{
				"first_invalid_field",
				"second_invalid_field",
				"field with spaces",
				"invalid.nested.field",
			}},
			tf: errtest.All(
				errtest.IsNonNil(),
				errtest.ErrorContains(`"first_invalid_field"`),
				errtest.ErrorContains(`"second_invalid_field"`),
				errtest.ErrorContains(`"field with spaces"`),
				errtest.ErrorContains(`"invalid.nested.field"`),
			),
		},
		{
			name: "InvalidFieldWithinNested",
			m:    &testingpb.Test{},
			mask: &fieldmaskpb.FieldMask{Paths: []string{"nested.invalid"}},
			tf: errtest.All(
				errtest.IsNonNil(),
				errtest.ErrorContains(`"nested.invalid"`),
			),
		},
		{
			name: "WithinRepeatedNested",
			m:    &testingpb.Test{},
			mask: &fieldmaskpb.FieldMask{Paths: []string{"rep_nested.foo"}},
			tf: errtest.All(
				errtest.IsNonNil(),
				errtest.ErrorContains("repeated"),
				errtest.ErrorContains("last position"),
			),
		},
		{
			name: "MultiplePathsWithWildcard",
			m:    &testingpb.Test{},
			mask: &fieldmaskpb.FieldMask{Paths: []string{"s", "*"}},
			tf: errtest.All(
				errtest.IsNonNil(),
				errtest.ErrorContains("wildcard"),
			),
		},
		{
			name: "OneofType",
			m:    &testingpb.Test{},
			mask: &fieldmaskpb.FieldMask{Paths: []string{"oo"}},
			tf: errtest.All(
				errtest.IsNonNil(),
				errtest.ErrorContains("oneof type"),
			),
		},
		{
			name: "EmptyPath",
			m:    &testingpb.Test{},
			mask: &fieldmaskpb.FieldMask{Paths: []string{""}},
			tf: errtest.All(
				errtest.IsNonNil(),
				errtest.ErrorContains("empty path"),
			),
		},
		{
			name: "EmptyPathPrefix",
			m:    &testingpb.Test{},
			mask: &fieldmaskpb.FieldMask{Paths: []string{".foo"}},
			tf: errtest.All(
				errtest.IsNonNil(),
				errtest.ErrorContains("empty path"),
			),
		},
		{
			name: "NestedPathOfScalar",
			m:    &testingpb.Test{},
			mask: &fieldmaskpb.FieldMask{Paths: []string{"s.foo"}},
			tf: errtest.All(
				errtest.IsNonNil(),
				errtest.ErrorContains("scalar"),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			tt.tf(t, Validate(tt.m, tt.mask))
		})
	}
}
