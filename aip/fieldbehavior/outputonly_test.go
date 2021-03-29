package fieldbehavior

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/proto"

	testingpb "github.com/Saser/pdp/aip/fieldbehavior/internal/testing/testing_go_proto"
)

func TestOutputOnlyPaths(t *testing.T) {
	for _, tt := range []struct {
		name string
		m    proto.Message
		want []string
	}{
		{
			name: "OK",
			m:    &testingpb.Test{},
			want: []string{
				"output_only",
				"nested.output_only",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := OutputOnlyPaths(tt.m)
			less := func(s1, s2 string) bool { return s1 < s2 }
			if diff := cmp.Diff(tt.want, got, cmpopts.EquateEmpty(), cmpopts.SortSlices(less)); diff != "" {
				t.Errorf("unexpected output (-want +got)\n%s", diff)
			}
		})
	}
}
