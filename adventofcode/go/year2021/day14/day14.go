package day14

import (
	"fmt"
	"strings"

	"github.com/Saser/pdp/adventofcode/go/intmath"
)

func Part1(input string) (string, error) {
	return solve(input, 1)
}

func Part2(input string) (string, error) {
	return solve(input, 2)
}

func solve(input string, part int) (string, error) {
	template, mapping := parse(input)
	e := newExpander(mapping)
	steps := 10
	if part == 2 {
		steps = 40
	}
	freq := e.Expand(template, steps)
	var counts []int
	for _, n := range freq {
		counts = append(counts, n)
	}
	min := intmath.Min(counts[0], counts[1:]...)
	max := intmath.Max(counts[0], counts[1:]...)
	return fmt.Sprint(max - min), nil
}

func parse(input string) (string, map[string]string) {
	paragraphs := strings.Split(strings.TrimSpace(input), "\n\n")
	template := paragraphs[0]
	mapping := make(map[string]string) // pair -> element
	for _, line := range strings.Split(paragraphs[1], "\n") {
		parts := strings.Split(line, " -> ")
		mapping[parts[0]] = parts[1]
	}
	return template, mapping
}

type frequency map[rune]int // element -> how many times it occurs

func (f frequency) Merge(ff frequency) {
	for r, n := range ff {
		f[r] += n
	}
}

type expander struct {
	mapping map[string]string // pair -> element

	memo map[args]frequency // expandPair arguments -> return value
}

func newExpander(mapping map[string]string) *expander {
	return &expander{
		mapping: mapping,
		memo:    make(map[args]frequency),
	}
}

type args struct {
	Pair  string // e.g., "CH"
	Steps int    // how many steps the pair should be expanded
}

func (e *expander) expandPair(a args) (f frequency) {
	f = make(frequency)
	// Early return: if there is no mapping for this pair, then no matter the
	// number of steps, the result will be the same. Similarly, if we are to
	// take 0 steps, the result is just the input pair.
	elem, ok := e.mapping[a.Pair]
	if !ok || a.Steps == 0 {
		for _, r := range a.Pair {
			f[r]++
		}
		return f
	}
	// Retrieve the result from the memo, if it is stored, otherwise make sure
	// it is stored when we return.
	if m, ok := e.memo[a]; ok {
		return m
	}
	defer func() {
		e.memo[a] = f
	}()
	// Recurse into the left pair and the right pair, and merge the results.
	a1 := args{
		Pair:  a.Pair[0:1] + elem,
		Steps: a.Steps - 1,
	}
	f.Merge(e.expandPair(a1))
	a2 := args{
		Pair:  elem + a.Pair[1:2],
		Steps: a.Steps - 1,
	}
	f.Merge(e.expandPair(a2))
	// elem has been double counted, so decrease it by one.
	f[rune(elem[0])]--
	return f
}

func (e *expander) Expand(template string, steps int) frequency {
	f := make(frequency)
	for i := 0; i < len(template)-1; i++ {
		a := args{
			Pair:  template[i : i+2],
			Steps: steps,
		}
		f.Merge(e.expandPair(a))
	}
	// All elements except the first and the last have been double counted, so
	// decrease them by one.
	for _, r := range template[1 : len(template)-1] {
		f[r]--
	}
	return f
}
