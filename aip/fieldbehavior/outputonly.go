package fieldbehavior

import (
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func OutputOnlyPaths(m proto.Message) []string {
	return outputOnlyPathsRecursive(m.ProtoReflect())
}

func outputOnlyPathsRecursive(m protoreflect.Message) []string {
	var out []string
	fields := m.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		name := string(fd.Name())
		if fd.Kind() == protoreflect.MessageKind {
			for _, oo := range outputOnlyPathsRecursive(m.Get(fd).Message()) {
				out = append(out, name+"."+oo)
			}
			continue
		}
		opts := fd.Options().(*descriptorpb.FieldOptions)
		fbs := proto.GetExtension(opts, annotations.E_FieldBehavior).([]annotations.FieldBehavior)
		for _, fb := range fbs {
			if fb == annotations.FieldBehavior_OUTPUT_ONLY {
				out = append(out, name)
			}
		}
	}
	return out
}
