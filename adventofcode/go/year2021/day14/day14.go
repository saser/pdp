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
// https://github.com/lindskogen/advent-of-code-2021/blob/06a8a20d3805e4d52bad3c94084161ba56d0c83c/src/main/kotlin/day14/main.kt

func expand(template string, mapping map[string]string, steps int) map[rune]int {
	// The basic overview of this solution is:
	//     1. Keep track of how many times each pair occurs.
	//     2. Count the initial frequencies of the elements in the template.
	//     3. In each step, every pair is replaced by the same number of each of
	//        its expansion pairs. The expansion also creates the same number of
	//        the new element, so its frequency is increased.

	// Count the initial set of pairs.
	pairs := make(map[string]int) // pair -> how many times it occurs
	for i := 0; i < len(template)-1; i++ {
		pairs[template[i:i+2]]++
	}

	// Count the initial element frequencies.
	freq := make(map[rune]int) // element -> how many times it occurs
	for _, r := range template {
		freq[r]++
	}

	for i := 0; i < steps; i++ {
		next := make(map[string]int)
		for pair, n := range pairs {
			// Assume that a mapping exists for every pair. That was the case
			// for my input at least.
			elem := mapping[pair]
			next[pair[0:1]+elem] += n
			next[elem+pair[1:2]] += n
			// The expansion created new occurrences of elem, so count them
			// here.
			freq[rune(elem[0])] += n
		}
		pairs = next
	}

	return freq
}
