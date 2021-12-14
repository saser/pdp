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
	steps := 10
	if part == 2 {
		steps = 40
	}
	freq := expand(template, mapping, steps)
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

// This solution was not my own. I learned it from glancing on this solution:
// https://github.com/mlonn/adventofcode-2021/blob/16284827a11ac16147d85e8bc405542cfa3848b6/14/extended_polymerization.go

func expand(template string, mapping map[string]string, steps int) map[rune]int {
	// The basic overview of this solution is:
	//     1. Keep track of how many times a certain pair occurs in the current expansion.
	//     2. In each step, every instance of a pair is replaced by the same
	//        number of each of its two expansion pairs (if the pair has an
	//        expansion).
	//     3. After all steps, count the number of occurrences of each element
	//        in each pair (accounting for double-counting as necessary).

	// We will proactively keep track of some double-counting, so we create freq
	// here.
	freq := make(map[rune]int) // element -> how many times it occurs

	// Count the initial set of pairs.
	pairs := make(map[string]int) // pair -> how many times it occurs
	for i := 0; i < len(template)-1; i++ {
		pairs[template[i:i+2]]++
	}
	// All elements except the first and the last in the template are now
	// double-counted, so proactively decrease their frequency here.
	for _, r := range template[1 : len(template)-1] {
		freq[r]--
	}

	// Iterate through the steps. The unexpanded map is used as a slight
	// optimization(?).
	unexpanded := make(map[string]int) // pair without expansion -> how many times it occurs
	for i := 0; i < steps; i++ {
		next := make(map[string]int)
		for pair, n := range pairs {
			elem, ok := mapping[pair]
			if !ok {
				unexpanded[pair] += n
				continue
			}
			next[pair[0:1]+elem] += n
			next[elem+pair[1:2]] += n
			// elem gets double counted above, so proactively decrease its
			// eventual frequency.
			freq[rune(elem[0])] -= n
		}
		pairs = next
	}

	// Accumulate the counts among all pairs, both unexpanded and the results of
	// the last iteration.
	for pair, n := range unexpanded {
		for _, r := range pair {
			freq[r] += n
		}
	}
	for pair, n := range pairs {
		for _, r := range pair {
			freq[r] += n
		}
	}
	return freq
}
