package problemid

import (
	"errors"
	"fmt"

	adventofcodepb "github.com/Saser/pdp/adventofcode/adventofcode_go_proto"
)

func Validate(id *adventofcodepb.ProblemID) error {
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
	return nil
}
