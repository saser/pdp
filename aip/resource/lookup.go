package resource

import (
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

type entry struct {
	md protoreflect.MessageDescriptor
	rd *annotations.ResourceDescriptor
}

var lookup = map[string]entry{}

func init() {
	protoregistry.GlobalTypes.RangeMessages(func(mt protoreflect.MessageType) bool {
		rd := DescriptorOf(mt.Zero().Interface())
		if rd == nil {
			return true
		}
		lookup[rd.GetType()] = entry{
			md: mt.Descriptor(),
			rd: rd,
		}
		return true
	})
}

func LookupMessage(typeString string) (protoreflect.MessageDescriptor, bool) {
	e, ok := lookup[typeString]
	return e.md, ok
}

func LookupResource(typeString string) (*annotations.ResourceDescriptor, bool) {
	e, ok := lookup[typeString]
	return e.rd, ok
}
