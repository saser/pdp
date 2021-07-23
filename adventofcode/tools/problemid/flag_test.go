package problemid

import (
	"flag"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"

	adventofcodepb "github.com/Saser/pdp/adventofcode/adventofcode_go_proto"
)

func TestFlag(t *testing.T) {
	orig := &adventofcodepb.ProblemID{
		Year: 2015,
		Day:  1,
		Part: adventofcodepb.ProblemID_ONE,
	}
	wantOrig := proto.Clone(orig).(*adventofcodepb.ProblemID)

	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	got := Flag(fs, "problem_id", orig, "Problem ID")
	if got == orig {
		t.Error("Flag() did not return a copy of the passed in problem ID")
	}

	arguments := []string{"-problem_id=year: 2016  day: 2  part: 2"}
	if err := fs.Parse(arguments); err != nil {
		t.Fatalf("fs.Parse(%q) = %v; want nil", arguments, err)
	}
	want := &adventofcodepb.ProblemID{
		Year: 2016,
		Day:  2,
		Part: adventofcodepb.ProblemID_TWO,
	}
	if diff := cmp.Diff(want, got, protocmp.Transform()); diff != "" {
		t.Errorf("unexpected result of flag parsing (-want +got)\n%s", diff)
	}

	if diff := cmp.Diff(wantOrig, orig, protocmp.Transform()); diff != "" {
		t.Errorf("unexpectedly changed passed in ProblemID value (-want +got)\n%s", diff)
	}
}

func TestFlagVar(t *testing.T) {
	got := &adventofcodepb.ProblemID{
		Year: 2015,
		Day:  1,
		Part: adventofcodepb.ProblemID_ONE,
	}

	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	FlagVar(fs, got, "problem_id", "")
	arguments := []string{"-problem_id=year: 2016  day: 2  part: 2"}
	if err := fs.Parse(arguments); err != nil {
		t.Fatalf("fs.Parse(%q) = %v; want nil", arguments, err)
	}
	want := &adventofcodepb.ProblemID{
		Year: 2016,
		Day:  2,
		Part: adventofcodepb.ProblemID_TWO,
	}
	if diff := cmp.Diff(want, got, protocmp.Transform()); diff != "" {
		t.Errorf("unexpected result of flag parsing (-want +got)\n%s", diff)
	}
}
