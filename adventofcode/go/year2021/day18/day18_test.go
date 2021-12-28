package day18

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func parseNumber(s string, depth int) *number {
	n := newNumber()
	n.Parse(s, depth)
	return n
}

func TestNumber_Reduce(t *testing.T) {
	const depth = 0
	for _, tt := range []struct {
		n    *number
		want *number
	}{
		{
			n:    parseNumber("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", depth),
			want: parseNumber("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", depth),
		},
		{
			n:    parseNumber("[[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]", depth),
			want: parseNumber("[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]", depth),
		},
	} {
		t.Run(tt.n.String(), func(t *testing.T) {
			tt.n.Reduce(depth)
			got := tt.n // tt.n modified in-place
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Reduce(): unexpected result (-want +got)\n%s", diff)
			}
		})
	}
}

func TestNumber_Magnitude(t *testing.T) {
	for _, tt := range []struct {
		n    *number
		want int
	}{
		{
			n:    parseNumber("[9,1]", 0),
			want: 29,
		},
		{
			n:    parseNumber("[1,9]", 0),
			want: 21,
		},
		{
			n:    parseNumber("[[9,1],[1,9]]", 0),
			want: 129,
		},
		{
			n:    parseNumber("[[1,2],[[3,4],5]]", 0),
			want: 143,
		},
		{
			n:    parseNumber("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 0),
			want: 1384,
		},
		{
			n:    parseNumber("[[[[1,1],[2,2]],[3,3]],[4,4]]", 0),
			want: 445,
		},
		{
			n:    parseNumber("[[[[3,0],[5,3]],[4,4]],[5,5]]", 0),
			want: 791,
		},
		{
			n:    parseNumber("[[[[5,0],[7,4]],[5,5]],[6,6]]", 0),
			want: 1137,
		},
		{
			n:    parseNumber("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 0),
			want: 3488,
		},
	} {
		t.Run(tt.n.String(), func(t *testing.T) {
			if got := tt.n.Magnitude(); got != tt.want {
				t.Errorf("Magnitude() = %v; want %v", got, tt.want)
			}
		})
	}
}
