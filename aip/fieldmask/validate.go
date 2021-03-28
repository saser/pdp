package fieldmask

import (
	"errors"
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func Validate(m proto.Message, mask *fieldmaskpb.FieldMask) error {
	if m == nil {
		panic("validate: proto message is nil")
	}
	paths := mask.GetPaths()
	// Special case: a single "*" wildcard is valid.
	if len(paths) == 1 && paths[0] == "*" {
		return nil
	}
	var errs []error
	for _, path := range paths {
		if err := validateRecursive(m.ProtoReflect(), path, path); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		ss := make([]string, len(errs))
		for i, err := range errs {
			ss[i] = err.Error()
		}
		return fmt.Errorf("validate: %s", strings.Join(ss, "; "))
	}
	return nil
}

func validateRecursive(m protoreflect.Message, fullPath, path string) error {
	desc := m.Descriptor()
	// Split the path into its dot-separated segments.
	segments := strings.Split(path, ".")
	// Doing segments[0] is safe since strings.Split always returns a slice with at least one
	// element.
	field := segments[0]
	// Look for special cases first.
	switch field {
	case "":
		return errors.New("empty path")
	case "*":
		return errors.New("wildcards are not supported together with other paths")
	}
	// field may point to a valid field within the message.
	fd := desc.Fields().ByName(protoreflect.Name(field))
	if fd == nil {
		// A path pointing to the a type (i.e., the name that is given to the oneof as
		// opposed to the fields within the oneof) is invalid. A oneof type is not a valid
		// field, so getting the FieldDescriptor for it will return nil, but getting the
		// OneofDescriptor will succeed.
		if desc.Oneofs().ByName(protoreflect.Name(field)) != nil {
			return fmt.Errorf("path %q points to a oneof type which is not supported", fullPath)
		}
		return fmt.Errorf("path %q points to a field that does not exist", fullPath)
	}
	// The path contains only one segment, which points to a valid field, so the path is valid.
	if len(segments) == 1 {
		return nil
	}
	// There are more segments. For the path to be valid, fd must point to a message-typed
	// field. Repeated fields, map fields, and groups are not supported.
	if fd.IsList() || fd.IsMap() {
		return fmt.Errorf("path %q contains a repeated/map field that is not in the last position", fullPath)
	}
	switch fd.Kind() {
	case protoreflect.MessageKind:
		return validateRecursive(m.Get(fd).Message(), fullPath, strings.TrimPrefix(path, field+"."))
	case protoreflect.GroupKind:
		return fmt.Errorf("path %q points to a group which is not supported", fullPath)
	default:
		// fd points to a single, scalar value, but the path contains more elements, which
		// is invalid.
		return fmt.Errorf("path %q points to a field within a scalar", fullPath)
	}
}
