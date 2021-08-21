package instancetest

import (
	"fmt"
	"os"
	"testing"

	"google.golang.org/protobuf/encoding/prototext"

	adventofcodepb "github.com/Saser/pdp/adventofcode/adventofcode_go_proto"
)

// SolveFunc is the type of a function solving one particular problem.
type SolveFunc func(input string) (string, error)

// FromFiles takes the paths to a set of files containing
// adventofcodepb.Instance messages in prototext format, and runs
// tests using f. If f returns an error or an answer different than
// the one in the Instance message, the test fails.
func FromFiles(t *testing.T, paths []string, f SolveFunc) {
	var instances []*adventofcodepb.Instance
	for _, path := range paths {
		b, err := os.ReadFile(path)
		if err != nil {
			t.Error(err)
			continue
		}
		instance := &adventofcodepb.Instance{}
		if err := prototext.Unmarshal(b, instance); err != nil {
			t.Errorf("couldn't parse %q as adventofcodepb.Instance: %v", path, err)
			continue
		}
		instances = append(instances, instance)
	}
	for _, instance := range instances {
		p := instance.GetProblem()
		name := fmt.Sprintf(
			"Year%d_Day%d_Part%d_%s",
			p.GetYear(),
			p.GetDay(),
			p.GetPart(),
			instance.GetName(),
		)
		t.Run(name, func(t *testing.T) {
			got, err := f(instance.GetInput())
			if err != nil {
				t.Fatal(err)
			}
			want := instance.GetAnswer()
			if got != want {
				t.Errorf("Solve() answer = %q; want %q", got, want)
				if testing.Verbose() {
					t.Logf("input = %q", instance.GetInput())
				}
				t.FailNow()
			}
		})
	}
}
