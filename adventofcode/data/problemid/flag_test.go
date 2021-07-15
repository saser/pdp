package problemid

import (
	"flag"
	"testing"

	adventofcodepb "github.com/Saser/pdp/adventofcode/adventofcode_go_proto"
)

func TestFlag(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	id := Flag(fs, "problem_id", &adventofcodepb.ProblemID{}, "Problem ID")
	arguments := []string{"-problem_id=year: 2015  day: 10  part: 1"}
	if err := fs.Parse(arguments); err != nil {
		t.Fatalf("fs.Parse(%q) = %v; want nil", arguments, err)
	}
	if got, want := id.GetYear(), int32(2015); got != want {
		t.Errorf("id.GetYear() = %v; want %v", got, want)
	}
	if got, want := id.GetDay(), int32(10); got != want {
		t.Errorf("id.GetDay() = %v; want %v", got, want)
	}
	if got, want := id.GetPart(), adventofcodepb.ProblemID_ONE; got != want {
		t.Errorf("id.GetDay() = %v; want %v", got, want)
	}
}
