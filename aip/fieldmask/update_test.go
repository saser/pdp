package fieldmask

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	testingpb "github.com/Saser/pdp/aip/fieldmask/internal/testing/testing_go_proto"
)

func TestUpdate(t *testing.T) {
	type testCase struct {
		name     string
		dst, src proto.Message
		mask     *fieldmaskpb.FieldMask
		want     proto.Message
	}
	test := func(t *testing.T, tt testCase) {
		t.Run(tt.name, func(t *testing.T) {
			dst := proto.Clone(tt.dst)
			src := proto.Clone(tt.src)
			if err := Update(dst, src, tt.mask); err != nil {
				t.Errorf("Update(%v, %v, %v) = %v; want nil", tt.dst, tt.src, tt.mask, err)
				if diff := cmp.Diff(tt.dst, dst, protocmp.Transform()); diff != "" {
					t.Errorf("dst updated when error was returned (-before +after)\n%s", diff)
				}
				t.FailNow()
			}
			if diff := cmp.Diff(tt.want, dst, protocmp.Transform()); diff != "" {
				t.Errorf("unexpected result of Update (-want +got)\n%s", diff)
			}
		})
	}

	nilMaskCases := []testCase{
		{
			name: "UpdateString",
			dst: &testingpb.Test{
				S: "string before",
			},
			src: &testingpb.Test{
				S: "string after",
			},
			mask: nil,
			want: &testingpb.Test{
				S: "string after",
			},
		},
		{
			name: "UpdateRepeatedString",
			dst: &testingpb.Test{
				RepS: []string{"first string before", "second string before"},
			},
			src: &testingpb.Test{
				RepS: []string{"first string after", "second string after"},
			},
			mask: nil,
			want: &testingpb.Test{
				RepS: []string{"first string after", "second string after"},
			},
		},
		{
			name: "UpdateNested_FullDst_FullSrc",
			dst: &testingpb.Test{
				Nested: &testingpb.Nested{
					Foo: "foo before",
					Bar: "bar before",
				},
			},
			src: &testingpb.Test{
				Nested: &testingpb.Nested{
					Foo: "foo after",
					Bar: "bar after",
				},
			},
			mask: nil,
			want: &testingpb.Test{
				Nested: &testingpb.Nested{
					Foo: "foo after",
					Bar: "bar after",
				},
			},
		},
		{
			name: "UpdateNested_PartialDst_FullSrc",
			dst: &testingpb.Test{
				Nested: &testingpb.Nested{
					Bar: "bar before",
				},
			},
			src: &testingpb.Test{
				Nested: &testingpb.Nested{
					Foo: "foo after",
					Bar: "bar after",
				},
			},
			mask: nil,
			want: &testingpb.Test{
				Nested: &testingpb.Nested{
					Foo: "foo after",
					Bar: "bar after",
				},
			},
		},
		{
			name: "UpdateNested_FullDst_PartialSrc",
			dst: &testingpb.Test{
				Nested: &testingpb.Nested{
					Foo: "foo before",
					Bar: "bar before",
				},
			},
			src: &testingpb.Test{
				Nested: &testingpb.Nested{
					Bar: "bar after",
				},
			},
			mask: nil,
			want: &testingpb.Test{
				Nested: &testingpb.Nested{
					Foo: "foo before",
					Bar: "bar after",
				},
			},
		},
		{
			name: "UpdateNested_PartialDst_PartialSrc_Replace",
			dst: &testingpb.Test{
				Nested: &testingpb.Nested{
					Foo: "foo before",
				},
			},
			src: &testingpb.Test{
				Nested: &testingpb.Nested{
					Foo: "foo after",
				},
			},
			mask: nil,
			want: &testingpb.Test{
				Nested: &testingpb.Nested{
					Foo: "foo after",
				},
			},
		},
		{
			name: "UpdateNested_PartialDst_PartialSrc_Merge",
			dst: &testingpb.Test{
				Nested: &testingpb.Nested{
					Foo: "foo before",
				},
			},
			src: &testingpb.Test{
				Nested: &testingpb.Nested{
					Bar: "bar after",
				},
			},
			mask: nil,
			want: &testingpb.Test{
				Nested: &testingpb.Nested{
					Foo: "foo before",
					Bar: "bar after",
				},
			},
		},
		{
			name: "UpdateRepeatedNested",
			dst: &testingpb.Test{
				RepNested: []*testingpb.Nested{
					{Foo: "first foo before", Bar: "first bar before"},
					{Foo: "second foo before", Bar: "second bar before"},
				},
			},
			src: &testingpb.Test{
				RepNested: []*testingpb.Nested{
					{Foo: "first foo after", Bar: "first bar after"},
					{Foo: "second foo after", Bar: "second bar after"},
				},
			},
			mask: nil,
			want: &testingpb.Test{
				RepNested: []*testingpb.Nested{
					{Foo: "first foo after", Bar: "first bar after"},
					{Foo: "second foo after", Bar: "second bar after"},
				},
			},
		},
		{
			name: "UpdateRepeatedNested_FullMessages",
			dst: &testingpb.Test{
				RepNested: []*testingpb.Nested{
					{Foo: "first foo before", Bar: "first bar before"},
					{Foo: "second foo before", Bar: "second bar before"},
				},
			},
			src: &testingpb.Test{
				RepNested: []*testingpb.Nested{
					{Foo: "first foo after", Bar: "first bar after"},
					{Foo: "second foo after", Bar: "second bar after"},
				},
			},
			mask: nil,
			want: &testingpb.Test{
				RepNested: []*testingpb.Nested{
					{Foo: "first foo after", Bar: "first bar after"},
					{Foo: "second foo after", Bar: "second bar after"},
				},
			},
		},
		{
			name: "UpdateRepeatedNested_PartialMessages",
			dst: &testingpb.Test{
				RepNested: []*testingpb.Nested{
					{Foo: "first foo before"},
					{Bar: "second bar before"},
				},
			},
			src: &testingpb.Test{
				RepNested: []*testingpb.Nested{
					{Bar: "first bar after"},
					{Foo: "second foo after"},
				},
			},
			mask: nil,
			want: &testingpb.Test{
				RepNested: []*testingpb.Nested{
					{Bar: "first bar after"},
					{Foo: "second foo after"},
				},
			},
		},
		{
			name: "UpdateOneof_SameCase",
			dst: &testingpb.Test{
				Oo: &testingpb.Test_OoS{OoS: "oneof string before"},
			},
			src: &testingpb.Test{
				Oo: &testingpb.Test_OoS{OoS: "oneof string after"},
			},
			mask: nil,
			want: &testingpb.Test{
				Oo: &testingpb.Test_OoS{OoS: "oneof string after"},
			},
		},
		{
			name: "UpdateOneof_DifferentCase",
			dst: &testingpb.Test{
				Oo: &testingpb.Test_OoS{OoS: "oneof string before"},
			},
			src: &testingpb.Test{
				Oo: &testingpb.Test_OoNested{OoNested: &testingpb.Nested{
					Foo: "oneof foo after",
					Bar: "oneof bar after",
				}},
			},
			mask: nil,
			want: &testingpb.Test{
				Oo: &testingpb.Test_OoNested{OoNested: &testingpb.Nested{
					Foo: "oneof foo after",
					Bar: "oneof bar after",
				}},
			},
		},
		{
			name: "UpdateWithinNestedOneof_FullDst_PartialSrc",
			dst: &testingpb.Test{
				Oo: &testingpb.Test_OoNested{OoNested: &testingpb.Nested{
					Foo: "oneof foo before",
					Bar: "oneof bar before",
				}},
			},
			src: &testingpb.Test{
				Oo: &testingpb.Test_OoNested{OoNested: &testingpb.Nested{
					Foo: "oneof foo after",
				}},
			},
			mask: nil,
			want: &testingpb.Test{
				Oo: &testingpb.Test_OoNested{OoNested: &testingpb.Nested{
					Foo: "oneof foo after",
					Bar: "oneof bar before",
				}},
			},
		},
		{
			name: "UpdateWithinNestedOneof_PartialDst_FullSrc",
			dst: &testingpb.Test{
				Oo: &testingpb.Test_OoNested{OoNested: &testingpb.Nested{
					Foo: "oneof foo before",
				}},
			},
			src: &testingpb.Test{
				Oo: &testingpb.Test_OoNested{OoNested: &testingpb.Nested{
					Foo: "oneof foo after",
					Bar: "oneof bar after",
				}},
			},
			mask: nil,
			want: &testingpb.Test{
				Oo: &testingpb.Test_OoNested{OoNested: &testingpb.Nested{
					Foo: "oneof foo after",
					Bar: "oneof bar after",
				}},
			},
		},
		{
			name: "UpdateWithinNestedOneof_PartialDst_PartialSrc_Replace",
			dst: &testingpb.Test{
				Oo: &testingpb.Test_OoNested{OoNested: &testingpb.Nested{
					Foo: "oneof foo before",
				}},
			},
			src: &testingpb.Test{
				Oo: &testingpb.Test_OoNested{OoNested: &testingpb.Nested{
					Foo: "oneof foo after",
				}},
			},
			mask: nil,
			want: &testingpb.Test{
				Oo: &testingpb.Test_OoNested{OoNested: &testingpb.Nested{
					Foo: "oneof foo after",
				}},
			},
		},
		{
			name: "UpdateWithinNestedOneof_PartialDst_PartialSrc_Merge",
			dst: &testingpb.Test{
				Oo: &testingpb.Test_OoNested{OoNested: &testingpb.Nested{
					Foo: "oneof foo before",
				}},
			},
			src: &testingpb.Test{
				Oo: &testingpb.Test_OoNested{OoNested: &testingpb.Nested{
					Bar: "oneof bar after",
				}},
			},
			mask: nil,
			want: &testingpb.Test{
				Oo: &testingpb.Test_OoNested{OoNested: &testingpb.Nested{
					Foo: "oneof foo before",
					Bar: "oneof bar after",
				}},
			},
		},
	}
	t.Run("NilMask", func(t *testing.T) {
		for _, tt := range nilMaskCases {
			test(t, tt)
		}
	})
	t.Run("EmptyMask", func(t *testing.T) {
		for _, tt := range nilMaskCases {
			tt := tt
			tt.mask = &fieldmaskpb.FieldMask{}
			test(t, tt)
		}
	})

	t.Run("StarMask", func(t *testing.T) {
		for _, tt := range []testCase{
			{
				name: "ClearEverything",
				dst: &testingpb.Test{
					S:    "string before",
					RepS: []string{"first string before", "second string before"},
					Nested: &testingpb.Nested{
						Foo: "foo before",
						Bar: "bar before",
					},
					RepNested: []*testingpb.Nested{
						{Foo: "first foo before", Bar: "first bar before"},
						{Foo: "second foo before", Bar: "second bar before"},
					},
					Oo: &testingpb.Test_OoS{OoS: "oneof string before"},
				},
				src:  &testingpb.Test{},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"*"}},
				want: &testingpb.Test{},
			},
			{
				name: "ClearSomeThings",
				dst: &testingpb.Test{
					S:    "string before",
					RepS: []string{"first string before", "second string before"},
					Nested: &testingpb.Nested{
						Foo: "foo before",
						Bar: "bar before",
					},
					RepNested: []*testingpb.Nested{
						{Foo: "first foo before", Bar: "first bar before"},
						{Foo: "second foo before", Bar: "second bar before"},
					},
					Oo: &testingpb.Test_OoS{OoS: "oneof string before"},
				},
				src: &testingpb.Test{
					S: "string after",
					Nested: &testingpb.Nested{
						Foo: "foo after",
						Bar: "bar after",
					},
				},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"*"}},
				want: &testingpb.Test{
					S: "string after",
					Nested: &testingpb.Nested{
						Foo: "foo after",
						Bar: "bar after",
					},
				},
			},
		} {
			test(t, tt)
		}
	})

	t.Run("NonEmptyMask", func(t *testing.T) {
		for _, tt := range []testCase{
			{
				name: "UpdateString",
				dst: &testingpb.Test{
					S: "string before",
				},
				src: &testingpb.Test{
					S:    "string after",
					RepS: []string{"this should not", "end up in dst"},
				},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"s"}},
				want: &testingpb.Test{
					S: "string after",
				},
			},
			{
				name: "UpdateRepeatedString",
				dst: &testingpb.Test{
					RepS: []string{"first string before", "second string before"},
				},
				src: &testingpb.Test{
					S:    "this should not end up in dst",
					RepS: []string{"first string after", "second string after"},
				},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"rep_s"}},
				want: &testingpb.Test{
					RepS: []string{"first string after", "second string after"},
				},
			},
			{
				name: "UpdateNested_EntireMessage_PartialSrc",
				dst: &testingpb.Test{
					Nested: &testingpb.Nested{
						Foo: "foo before",
						Bar: "bar before",
					},
				},
				src: &testingpb.Test{
					S: "this should not end up in dst",
					Nested: &testingpb.Nested{
						Foo: "foo after",
					},
				},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"nested"}},
				want: &testingpb.Test{
					Nested: &testingpb.Nested{
						Foo: "foo after",
					},
				},
			},
			{
				name: "UpdateNested_EntireMessage_FullSrc",
				dst: &testingpb.Test{
					Nested: &testingpb.Nested{
						Foo: "foo before",
						Bar: "bar before",
					},
				},
				src: &testingpb.Test{
					S: "this should not end up in dst",
					Nested: &testingpb.Nested{
						Foo: "foo after",
						Bar: "bar after",
					},
				},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"nested"}},
				want: &testingpb.Test{
					Nested: &testingpb.Nested{
						Foo: "foo after",
						Bar: "bar after",
					},
				},
			},
			{
				name: "UpdateWithinNested_SinglePath_PartialSrc",
				dst: &testingpb.Test{
					Nested: &testingpb.Nested{
						Foo: "foo before",
						Bar: "bar before",
					},
				},
				src: &testingpb.Test{
					S: "this should not end up in dst",
					Nested: &testingpb.Nested{
						Foo: "foo after",
					},
				},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"nested.foo"}},
				want: &testingpb.Test{
					Nested: &testingpb.Nested{
						Foo: "foo after",
						Bar: "bar before",
					},
				},
			},
			{
				name: "UpdateWithinNested_AllPaths_PartialSrc",
				dst: &testingpb.Test{
					Nested: &testingpb.Nested{
						Foo: "foo before",
						Bar: "bar before",
					},
				},
				src: &testingpb.Test{
					S: "this should not end up in dst",
					Nested: &testingpb.Nested{
						Foo: "foo after",
					},
				},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"nested.foo", "nested.bar"}},
				want: &testingpb.Test{
					Nested: &testingpb.Nested{
						Foo: "foo after",
					},
				},
			},
			{
				name: "UpdateWithinNested_SinglePath_FullSrc",
				dst: &testingpb.Test{
					Nested: &testingpb.Nested{
						Foo: "foo before",
						Bar: "bar before",
					},
				},
				src: &testingpb.Test{
					S: "this should not end up in dst",
					Nested: &testingpb.Nested{
						Foo: "foo after",
						Bar: "bar after",
					},
				},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"nested.foo"}},
				want: &testingpb.Test{
					Nested: &testingpb.Nested{
						Foo: "foo after",
						Bar: "bar before",
					},
				},
			},
			{
				name: "UpdateWithinNested_AllPaths_FullSrc",
				dst: &testingpb.Test{
					Nested: &testingpb.Nested{
						Foo: "foo before",
						Bar: "bar before",
					},
				},
				src: &testingpb.Test{
					S: "this should not end up in dst",
					Nested: &testingpb.Nested{
						Foo: "foo after",
						Bar: "bar after",
					},
				},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"nested.foo", "nested.bar"}},
				want: &testingpb.Test{
					Nested: &testingpb.Nested{
						Foo: "foo after",
						Bar: "bar after",
					},
				},
			},
			{
				name: "UpdateOneofString_SinglePath",
				dst: &testingpb.Test{
					Oo: &testingpb.Test_OoS{OoS: "oneof string before"},
				},
				src: &testingpb.Test{
					S:  "this should not end up in dst",
					Oo: &testingpb.Test_OoS{OoS: "oneof string after"},
				},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"oo_s"}},
				want: &testingpb.Test{
					Oo: &testingpb.Test_OoS{OoS: "oneof string after"},
				},
			},
			{
				name: "UpdateOneofString_AllPaths",
				dst: &testingpb.Test{
					Oo: &testingpb.Test_OoS{OoS: "oneof string before"},
				},
				src: &testingpb.Test{
					S:  "this should not end up in dst",
					Oo: &testingpb.Test_OoS{OoS: "oneof string after"},
				},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"oo_s", "oo_nested"}},
				want: &testingpb.Test{
					Oo: &testingpb.Test_OoS{OoS: "oneof string after"},
				},
			},
			{
				name: "UpdateOneofString_SetEmpty",
				dst: &testingpb.Test{
					Oo: &testingpb.Test_OoS{OoS: "oneof string before"},
				},
				src: &testingpb.Test{
					S:  "this should not end up in dst",
					Oo: &testingpb.Test_OoS{OoS: ""},
				},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"oo_s"}},
				want: &testingpb.Test{
					Oo: &testingpb.Test_OoS{OoS: ""},
				},
			},
			{
				name: "ClearOneofString",
				dst: &testingpb.Test{
					Oo: &testingpb.Test_OoS{OoS: "oneof string before"},
				},
				src: &testingpb.Test{
					S:  "this should not end up in dst",
					Oo: nil, // set explicitly for clarity
				},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"oo_s"}},
				want: &testingpb.Test{},
			},
			{
				name: "ReplaceOneofStringWithNested",
				dst: &testingpb.Test{
					Oo: &testingpb.Test_OoS{OoS: "oneof string before"},
				},
				src: &testingpb.Test{
					S: "this should not end up in dst",
					Oo: &testingpb.Test_OoNested{OoNested: &testingpb.Nested{
						Foo: "oneof foo after",
						Bar: "oneof bar after",
					}},
				},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"oo_s", "oo_nested"}}, // note that both must be specified
				want: &testingpb.Test{
					Oo: &testingpb.Test_OoNested{OoNested: &testingpb.Nested{
						Foo: "oneof foo after",
						Bar: "oneof bar after",
					}},
				},
			},
			{
				name: "UpdateWithinOneofNested",
				dst: &testingpb.Test{
					Oo: &testingpb.Test_OoNested{OoNested: &testingpb.Nested{
						Foo: "oneof foo before",
						Bar: "oneof bar before",
					}},
				},
				src: &testingpb.Test{
					S: "this should not end up in dst",
					Oo: &testingpb.Test_OoNested{OoNested: &testingpb.Nested{
						Foo: "oneof foo after",
						Bar: "this should not end up in dst",
					}},
				},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"oo_nested.foo"}},
				want: &testingpb.Test{
					Oo: &testingpb.Test_OoNested{OoNested: &testingpb.Nested{
						Foo: "oneof foo after",
						Bar: "oneof bar before",
					}},
				},
			},
			{
				name: "CreateWithinNested",
				dst: &testingpb.Test{
					Nested: nil, // explicitly set for clarity
				},
				src: &testingpb.Test{
					S: "this should not end up in dst",
					Nested: &testingpb.Nested{
						Foo: "foo after",
						Bar: "this should not end up in dst",
					},
				},
				mask: &fieldmaskpb.FieldMask{Paths: []string{"nested.foo"}}, // note that both must be specified
				want: &testingpb.Test{
					Nested: &testingpb.Nested{
						Foo: "foo after",
					},
				},
			},
			{
				name: "ClearEverything",
				dst: &testingpb.Test{
					S:    "string before",
					RepS: []string{"first string before", "second string before"},
					Nested: &testingpb.Nested{
						Foo: "foo before",
						Bar: "bar before",
					},
					RepNested: []*testingpb.Nested{
						{Foo: "first foo before", Bar: "first bar before"},
						{Foo: "second foo before", Bar: "second bar before"},
					},
					Oo: &testingpb.Test_OoS{OoS: "oneof string before"},
				},
				src: &testingpb.Test{},
				mask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"s",
						"rep_s",
						"nested",
						"rep_nested",
						"oo_s",
						"oo_nested",
					},
				},
				want: &testingpb.Test{},
			},
		} {
			test(t, tt)
		}
	})
}
