package problemid

import (
	"flag"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"

	adventofcodepb "github.com/Saser/pdp/adventofcode/adventofcode_go_proto"
)

func TestFlagValue(t *testing.T) {
	for _, tt := range []struct {
		arg  string
		want *adventofcodepb.ProblemID
	}{
		{
			arg: "-problem_id=year: 2015 day: 1 part: 1",
			want: &adventofcodepb.ProblemID{
				Year: 2015,
				Day:  1,
				Part: adventofcodepb.ProblemID_ONE,
			},
		},
		{
			arg: "-problem_id=year: 2015 day: 1",
			want: &adventofcodepb.ProblemID{
				Year: 2015,
				Day:  1,
				Part: adventofcodepb.ProblemID_PART_UNSPECIFIED,
			},
		},
	} {
		got := &adventofcodepb.ProblemID{}
		fs := flag.NewFlagSet("test", flag.ContinueOnError)
		fs.Var(&FlagValue{ProblemID: got}, "problem_id", "")
		arguments := []string{tt.arg}
		if err := fs.Parse(arguments); err != nil {
			t.Errorf("fs.Parse(%q) = %v; want nil", arguments, err)
			continue
		}
		if diff := cmp.Diff(tt.want, got, protocmp.Transform()); diff != "" {
			t.Errorf("unexpected result of argument parsing (-want +got)\n%s", diff)
		}
	}
}
