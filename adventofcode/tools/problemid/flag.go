package problemid

import (
	"errors"
	"flag"
	"fmt"

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"

	adventofcodepb "github.com/Saser/pdp/adventofcode/adventofcode_go_proto"
)

type problemIDValue struct {
	id *adventofcodepb.ProblemID
}

func (v *problemIDValue) String() string {
	data, err := prototext.MarshalOptions{Multiline: false}.Marshal(v.id)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (v *problemIDValue) Set(s string) error {
	id := &adventofcodepb.ProblemID{}
	if err := prototext.Unmarshal([]byte(s), id); err != nil {
		return err
	}
	if year := id.GetYear(); year < 2015 || year > 2021 {
		return fmt.Errorf("problem ID has year outside range [2015, 2021]: %d", year)
	}
	if day := id.GetDay(); day < 1 || day > 25 {
		return fmt.Errorf("problem ID has day outside range [1, 25]: %d", day)
	}
	switch part := id.GetPart(); part {
	case adventofcodepb.ProblemID_PART_UNSPECIFIED:
		return errors.New("a part must be specified")
	case adventofcodepb.ProblemID_ONE:
		// this is fine
	case adventofcodepb.ProblemID_TWO:
		if id.GetDay() == 25 {
			return errors.New("problem ID specifies part 2 for day 25")
		}
	default:
		return fmt.Errorf("invalid part: %v", part)
	}
	proto.Reset(v.id)
	proto.Merge(v.id, id)
	return nil
}

func (v *problemIDValue) Get() interface{} {
	return v.id
}

func Flag(fs *flag.FlagSet, name string, value *adventofcodepb.ProblemID, usage string) *adventofcodepb.ProblemID {
	v := &problemIDValue{id: proto.Clone(value).(*adventofcodepb.ProblemID)}
	fs.Var(v, name, usage)
	return v.id
}
