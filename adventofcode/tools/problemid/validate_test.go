package problemid

import (
	"testing"

	adventofcodepb "github.com/Saser/pdp/adventofcode/adventofcode_go_proto"
)

func TestValidate(t *testing.T) {
	for _, id := range []*adventofcodepb.ProblemID{} {
		if err := Validate(id); err != nil {
			t.Errorf("Validate(%v) = %v; want nil", id, err)
		}
	}
}

func TestValidate_Error(t *testing.T) {
	for _, id := range []*adventofcodepb.ProblemID{
		// Invalid years.
		{Year: -2015, Day: 1, Part: adventofcodepb.ProblemID_ONE},
		{Year: 2014, Day: 1, Part: adventofcodepb.ProblemID_ONE},
		{Year: 2030, Day: 1, Part: adventofcodepb.ProblemID_ONE},

		// Invalid days.
		{Year: 2015, Day: 0, Part: adventofcodepb.ProblemID_ONE},
		{Year: 2015, Day: 26, Part: adventofcodepb.ProblemID_ONE},
		{Year: 2015, Day: -25, Part: adventofcodepb.ProblemID_ONE},

		// Invalid parts.
		{Year: 2015, Day: 25, Part: adventofcodepb.ProblemID_PART_UNSPECIFIED},
		{Year: 2015, Day: 25, Part: adventofcodepb.ProblemID_TWO}, // there is no day 25, part 2
	} {
		if err := Validate(id); err == nil {
			t.Errorf("Validate(%v) = nil error; want non-nil", id)
		}
	}
}
