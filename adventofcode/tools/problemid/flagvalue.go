package problemid

import (
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"

	adventofcodepb "github.com/Saser/pdp/adventofcode/adventofcode_go_proto"
)

// FlagValue implements the flag.Value interface for ProblemID. The prototext
// format is used to encode and decode a ProblemID as a string flag.
type FlagValue struct {
	ProblemID *adventofcodepb.ProblemID
}

func (v FlagValue) String() string {
	data, err := prototext.MarshalOptions{Multiline: false}.Marshal(v.ProblemID)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (v FlagValue) Set(s string) error {
	id := &adventofcodepb.ProblemID{}
	if err := prototext.Unmarshal([]byte(s), id); err != nil {
		return err
	}
	proto.Reset(v.ProblemID)
	proto.Merge(v.ProblemID, id)
	return nil
}

func (v FlagValue) Get() interface{} {
	return v.ProblemID
}
