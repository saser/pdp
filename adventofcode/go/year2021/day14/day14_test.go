package day14

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func makeExpander(input string) *expander {
	_, mapping := parse(input)
	return newExpander(mapping)
}

func makeFrequency(s string) frequency {
	f := make(frequency)
	for _, r := range s {
		f[r]++
	}
	return f
}

func TestExpander_Expand(t *testing.T) {
	// Input from the example.
	const input = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`
	e := makeExpander(input)
	for _, tt := range []struct {
		template string
		steps    int
		want     frequency
	}{
		// The following three test cases test that expanding 1 + 1 steps works,
		// as well as expanding 2 steps.
		{
			template: "NN",
			steps:    1,
			want:     makeFrequency("NCN"),
		},
		{
			template: "NCN",
			steps:    1,
			want:     makeFrequency("NBCCN"),
		},
		{
			template: "NN",
			steps:    2,
			want:     makeFrequency("NBCCN"),
		},
		// The following test cases use the example from the instructions with
		// the result of the first 4 steps.
		{
			template: "NNCB",
			steps:    1,
			want:     makeFrequency("NCNBCHB"),
		},
		{
			template: "NNCB",
			steps:    2,
			want:     makeFrequency("NBCCNBBBCBHCB"),
		},
		{
			template: "NNCB",
			steps:    3,
			want:     makeFrequency("NBBBCNCCNBBNBNBBCHBHHBCHB"),
		},
		{
			template: "NNCB",
			steps:    4,
			want:     makeFrequency("NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB"),
		},
	} {
		t.Run(fmt.Sprintf("template=%s,steps=%d", tt.template, tt.steps), func(t *testing.T) {
			got := e.Expand(tt.template, tt.steps)
			strgot := make(map[string]int)
			for r, n := range got {
				strgot[string(r)] = n
			}
			strwant := make(map[string]int)
			for r, n := range tt.want {
				strwant[string(r)] = n
			}
			if diff := cmp.Diff(strwant, strgot); diff != "" {
				t.Errorf("unexpected result of Expand (-want +got)\n%s", diff)
			}
		})
	}
}
