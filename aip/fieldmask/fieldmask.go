// Package fieldmask contains helpful functions for working with field masks.
package fieldmask

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Merge updates fields in dst with the values of the corresponding fields in src, based on the
// mask. A nil mask, or a mask with an empty list of paths, makes dst equal to src. Currently only
// scalar values are supported. Merge returns an error if dst and src do not have the same message
// type, or if the field mask is not valid for the given message type.
func Merge(dst, src proto.Message, mask *fieldmaskpb.FieldMask) error {
	if dstT, srcT := proto.MessageName(dst), proto.MessageName(src); dstT != srcT {
		return fmt.Errorf("merge with field mask: dst has type %q but src has type %q", dstT, srcT)
	}
	if mask == nil {
		proto.Reset(dst)
		proto.Merge(dst, src)
		return nil
	}
	if !mask.IsValid(dst) {
		return fmt.Errorf("merge with field mask: mask %v is invalid", mask)
	}
	for _, path := range mask.GetPaths() {
		if strings.Contains(path, ".") {
			return fmt.Errorf("merge with field mask: nested path %q is invalid", path)
		}
	}
	dstpref := dst.ProtoReflect()
	srcpref := src.ProtoReflect()
	for _, path := range mask.GetPaths() {
		name := protoreflect.Name(path)
		fd := dstpref.Descriptor().Fields().ByName(name)
		dstpref.Set(fd, srcpref.Get(fd))
	}
	return nil
}
