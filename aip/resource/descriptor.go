package resource

import (
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
)

func Descriptor(m proto.Message) *annotations.ResourceDescriptor {
	return proto.GetExtension(m.ProtoReflect().Descriptor().Options(), annotations.E_Resource).(*annotations.ResourceDescriptor)
}
