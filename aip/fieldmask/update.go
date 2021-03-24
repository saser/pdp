package fieldmask

import (
	"errors"
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Update sets fields in dst equal to their counterparts in src, subject to the field mask. Repeated
// fields are replaced rather than appended (in contrast to the
// "google.golang.org/protobuf/proto".Merge function). A nil mask, or a mask with no paths, is
// treated as a mask that specifies all values that are set on the wire. As a special case, a mask
// with a single path that is set to "*" is treated as if all fields were specified.
func Update(dst, src proto.Message, mask *fieldmaskpb.FieldMask) error {
	if dst == nil {
		return errors.New("dst is nil")
	}
	if src == nil {
		return errors.New("src is nil")
	}
	if dstType, srcType := proto.MessageName(dst), proto.MessageName(src); dstType != srcType {
		return fmt.Errorf("src has type %q but dst has type %q", dstType, srcType)
	}

	dstrefl := dst.ProtoReflect()
	srcrefl := src.ProtoReflect()

	// If an empty field mask is given, it should be treated as a mask that includes all fields
	// set on the wire.
	if len(mask.GetPaths()) == 0 {
		// We should replace, instead of merge, the values of a repeated field. Therefore,
		// go over all wire-set repeated fields in src, and clear the corresponding field in
		// dst.
		srcrefl.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
			if fd.IsList() {
				dstrefl.Clear(fd)
			}
			return true
		})
		// proto.Merge implements the logic for setting fields that are set on the wire, so
		// we can just use that.
		proto.Merge(dst, src)
		return nil
	}

	// Special case: a single path that is the `*` wildcard. In this case, src should fully
	// replace dst.
	if paths := mask.GetPaths(); len(paths) == 1 && paths[0] == "*" {
		// Resetting dst and then merging the now empty message will make dst equal to src.
		proto.Reset(dst)
		proto.Merge(dst, src)
		return nil
	}

	if !mask.IsValid(dst) {
		return errors.New("mask is invalid") // TODO: make this error message better.
	}

	mask.Normalize()
	for _, path := range mask.GetPaths() {
		updateRecursive(dstrefl, srcrefl, strings.Split(path, "."))
	}

	return nil
}

func updateRecursive(dst, src protoreflect.Message, path []string) {
	fd := src.Descriptor().Fields().ByName(protoreflect.Name(path[0]))
	// If there is only one segment in the path, we should do a full replacement, regardless of
	// whether it is a scalar or a message.
	if len(path) == 1 {
		// If the field is not populated in src, it means that the field should be cleared.
		if !src.Has(fd) {
			dst.Clear(fd)
			return
		}
		// The field is populated in src, so simply copy its value over to dst.
		dst.Set(fd, src.Get(fd))
		return
	}
	// There are more than one segment, so we should recurse. We use dst.Mutable since, well, we
	// are going to mutate the message. However, there is also the fact that using dst.Get means
	// that if fd is not set in dst, dst.Get returns an empty, read-only view of the value,
	// which causes panics if we try to mutate it down the line.
	updateRecursive(dst.Mutable(fd).Message(), src.Get(fd).Message(), path[1:])
}
