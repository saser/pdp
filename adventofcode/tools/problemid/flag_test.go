package problemid

import (
	"flag"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"

	adventofcodepb "github.com/Saser/pdp/adventofcode/adventofcode_go_proto"
)

func TestFlag(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	got := Flag(fs, "problem_id", &adventofcodepb.ProblemID{}, "Problem ID")
	arguments := []string{"-problem_id=year: 2015  day: 10  part: 1"}
	if err := fs.Parse(arguments); err != nil {
		t.Fatalf("fs.Parse(%q) = %v; want nil", arguments, err)
	}
	want := &adventofcodepb.ProblemID{
		Year: 2015,
		Day:  10,
		Part: adventofcodepb.ProblemID_ONE,
	}
	if diff := cmp.Diff(want, got, protocmp.Transform()); diff != "" {
		t.Errorf("unexpected result of flag parsing (-want +got)\n%s", diff)
	}
}
