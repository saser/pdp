package fieldmask

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	examplepb "github.com/Saser/pdp/aip/fieldmask/example/example_go_proto"
)

func TestMerge_OK_Scalars(t *testing.T) {
	dst := &examplepb.Scalars{
		S:   "before",
		I32: 32,
		I64: 64,
		F:   3.2,
		D:   6.4,
		B:   false,
	}
	src := &examplepb.Scalars{
		S:   "after",
		I32: 320,
		I64: 640,
		F:   3.25,
		D:   6.45,
		B:   true,
	}
	makeMask := func(paths ...string) *fieldmaskpb.FieldMask {
		return &fieldmaskpb.FieldMask{Paths: paths}
	}
	copyField := func(field protoreflect.Name) *examplepb.Scalars {
		dst2 := proto.Clone(dst).(*examplepb.Scalars)
		v := src.ProtoReflect().Get(src.ProtoReflect().Descriptor().Fields().ByName(field))
		dst2.ProtoReflect().Set(dst2.ProtoReflect().Descriptor().Fields().ByName(field), v)
		return dst2
	}
	for _, tt := range []struct {
		name string
		mask *fieldmaskpb.FieldMask
		want *examplepb.Scalars
	}{
		{
			name: "NilMask",
			mask: nil,
			want: src,
		},
		{
			name: "EmptyMask",
			mask: &fieldmaskpb.FieldMask{},
			want: dst,
		},
		{
			name: "MultipleFields",
			mask: makeMask("s", "i32", "i64"),
			want: func() *examplepb.Scalars {
				dst2 := proto.Clone(dst).(*examplepb.Scalars)
				dst2.S = src.S
				dst2.I32 = src.I32
				dst2.I64 = src.I64
				return dst2
			}(),
		},
		{
			name: "AllFields",
			mask: makeMask("s", "i32", "i64", "f", "d", "b"),
			want: src,
		},
		{
			name: "String",
			mask: makeMask("s"),
			want: copyField("s"),
		},
		{
			name: "Int32",
			mask: makeMask("i32"),
			want: copyField("i32"),
		},
		{
			name: "Int64",
			mask: makeMask("i64"),
			want: copyField("i64"),
		},
		{
			name: "Float",
			mask: makeMask("f"),
			want: copyField("f"),
		},
		{
			name: "Double",
			mask: makeMask("d"),
			want: copyField("d"),
		},
		{
			name: "Bool",
			mask: makeMask("b"),
			want: copyField("b"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			dst := proto.Clone(dst)
			src := proto.Clone(src)
			if err := Merge(dst, src, tt.mask); err != nil {
				t.Errorf("Merge(%v, %v, %v) = %v; want nil", dst, src, tt.mask, err)
			}
			if diff := cmp.Diff(tt.want, dst, protocmp.Transform()); diff != "" {
				t.Errorf("wrong result of merge (-want +got)\n%s", diff)
			}
		})
	}
}

func TestMerge_Errors(t *testing.T) {
	for _, tt := range []struct {
		name     string
		dst, src proto.Message
		mask     *fieldmaskpb.FieldMask
	}{
		{
			name: "DifferentTypes",
			dst:  &examplepb.Scalars{},
			src:  &examplepb.Nested{},
			mask: nil,
		},
		{
			name: "InvalidMask",
			dst:  &examplepb.Scalars{},
			src:  &examplepb.Scalars{},
			mask: &fieldmaskpb.FieldMask{
				Paths: []string{"invalid"},
			},
		},
		{
			name: "NestedPath",
			dst:  &examplepb.Nested{},
			src:  &examplepb.Nested{},
			mask: &fieldmaskpb.FieldMask{
				Paths: []string{"scalars.s"},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			dst := proto.Clone(tt.dst)
			src := proto.Clone(tt.src)
			if err := Merge(dst, src, tt.mask); err == nil {
				t.Errorf("Merge(%v, %v, %v) = nil; want non-nil", dst, src, tt.mask)
			}
		})
	}
}
