package day18

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func equalNumber(n1, n2 *number) bool {
	if n1.root.String() != n2.root.String() {
		return false
	}
	var rs1 []string
	for r := n1.front; r != nil; r = r.next {
		rs1 = append(rs1, fmt.Sprint(r.Value))
	}
	var rs2 []string
	for r := n2.front; r != nil; r = r.next {
		rs2 = append(rs2, fmt.Sprint(r.Value))
	}
	return strings.Join(rs1, ",") == strings.Join(rs2, ",")
}

func TestParseNumber(t *testing.T) {
	// p is a shorthand for constructing a pair. Left and right must either be
	// an int or a *pair.
	p := func(left, right any) *node {
		p := &node{}
		switch v := left.(type) {
		case int:
			p.Left = &node{
				Value: v,
			}
		case *node:
			p.Left = v
		}
		switch v := right.(type) {
		case int:
			p.Right = &node{
				Value: v,
			}
		case *node:
			p.Right = v
		}
		return p
	}
	for _, tt := range []struct {
		line string
		want *node
	}{
		{
			line: "[1,2]",
			want: p(1, 2),
		},
		{
			line: "[[1,2],3]",
			want: p(p(1, 2), 3),
		},
		{
			line: "[9,[8,7]]",
			want: p(9, p(8, 7)),
		},
		{
			line: "[[1,9],[8,5]]",
			want: p(p(1, 9), p(8, 5)),
		},
		{
			line: "[[[[1,2],[3,4]],[[5,6],[7,8]]],9]",
			want: p(p(p(p(1, 2), p(3, 4)), p(p(5, 6), p(7, 8))), 9),
		},
		{
			line: "[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]",
			want: p(p(p(9, p(3, 8)), p(p(0, 9), 6)), p(p(p(3, 7), p(4, 9)), 3)),
		},
		{
			line: "[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]",
			want: p(p(p(p(1, 3), p(5, 3)), p(p(1, 3), p(8, 7))), p(p(p(4, 9), p(6, 9)), p(p(8, 2), p(7, 3)))),
		},
	} {
		got := parseNode(tt.line)
		if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreUnexported(node{})); diff != "" {
			t.Errorf("unexpected result from parsing (-want +got)\n%s", diff)
		}
	}
}

func TestNumber_Magnitude(t *testing.T) {
	for _, tt := range []struct {
		n    *number
		want int
	}{
		{
			n:    parseNumber("10"),
			want: 10,
		},
		{
			n:    parseNumber("[9,1]"),
			want: 29,
		},
		{
			n:    parseNumber("[[9,1],[1,9]]"),
			want: 129,
		},
		{
			n:    parseNumber("[[1,2],[[3,4],5]]"),
			want: 143,
		},
		{
			n:    parseNumber("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"),
			want: 1384,
		},
		{
			n:    parseNumber("[[[[1,1],[2,2]],[3,3]],[4,4]]"),
			want: 445,
		},
		{
			n:    parseNumber("[[[[3,0],[5,3]],[4,4]],[5,5]]"),
			want: 791,
		},
		{
			n:    parseNumber("[[[[5,0],[7,4]],[5,5]],[6,6]]"),
			want: 1137,
		},
		{
			n:    parseNumber("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"),
			want: 3488,
		},
	} {
		if got := tt.n.Magnitude(); got != tt.want {
			t.Errorf("Magnitude() = %v; want %v", got, tt.want)
		}
	}
}

func TestNumber_Add(t *testing.T) {
	for _, tt := range []struct {
		n     *number
		other *number
		want  *number
	}{
		{
			n:     parseNumber("[[[[4,3],4],4],[7,[[8,4],9]]]"),
			other: parseNumber("[1,1]"),
			want:  parseNumber("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"),
		},
	} {
		tt.n.Add(tt.other)
		got := tt.n // Add() modified the receiver in-place
		if diff := cmp.Diff(tt.want, got, cmp.Comparer(equalNumber)); diff != "" {
			t.Errorf("unexpected result from Add() (-want +got)\n%s", diff)
		}
	}
}

func TestNumber_Reduce(t *testing.T) {
	for _, tt := range []struct {
		n    *number
		want *number
	}{
		{
			n:    parseNumber("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"),
			want: parseNumber("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"),
		},
	} {
		tt.n.Reduce()
		got := tt.n // Reduce() modified tt.n in-place
		if diff := cmp.Diff(tt.want, got, cmp.Comparer(equalNumber)); diff != "" {
			t.Errorf("unexpected result of Reduce() (-want +got)\n%s", diff)
		}
	}
}

func TestNumber_Explode(t *testing.T) {
	for _, tt := range []struct {
		n    *number
		want *number
	}{
		{
			n:    parseNumber("[[[[[9,8],1],2],3],4]"),
			want: parseNumber("[[[[0,9],2],3],4]"),
		},
		{
			n:    parseNumber("[7,[6,[5,[4,[3,2]]]]]"),
			want: parseNumber("[7,[6,[5,[7,0]]]]"),
		},
		{
			n:    parseNumber("[[6,[5,[4,[3,2]]]],1]"),
			want: parseNumber("[[6,[5,[7,0]]],3]"),
		},
		{
			n:    parseNumber("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"),
			want: parseNumber("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"),
		},
		{
			n:    parseNumber("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"),
			want: parseNumber("[[3,[2,[8,0]]],[9,[5,[7,0]]]]"),
		},
	} {
		t.Run(tt.n.root.String(), func(t *testing.T) {
			changed := tt.n.Explode()
			if got, want := changed, true; got != want {
				t.Fatalf("Explode() = %v; want %v", got, want)
			}
			got := tt.n // Explode() modified tt.n in-place
			if diff := cmp.Diff(tt.want, got, cmp.Comparer(equalNumber)); diff != "" {
				t.Errorf("unexpected result of Explode() (-want +got)\n%s", diff)
			}
		})
	}
}

func TestNumber_Split(t *testing.T) {
	for _, tt := range []struct {
		n    *number
		want *number
	}{
		{
			n:    parseNumber("[[[[0,7],4],[15,[0,13]]],[1,1]]"),
			want: parseNumber("[[[[0,7],4],[[7,8],[0,13]]],[1,1]]"),
		},
		{
			n:    parseNumber("[[[[0,7],4],[[7,8],[0,13]]],[1,1]]"),
			want: parseNumber("[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]"),
		},
	} {
		changed := tt.n.Split()
		if got, want := changed, true; got != want {
			t.Fatalf("Split() = %v; want %v", got, want)
		}
		got := tt.n // Split() modified tt.n in-place
		if diff := cmp.Diff(tt.want, got, cmp.Comparer(equalNumber)); diff != "" {
			t.Errorf("unexpected result of Split() (-want +got)\n%s", diff)
		}
	}
}
