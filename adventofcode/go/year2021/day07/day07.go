package day07

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	return solve(input, 1)
}

func Part2(input string) (string, error) {
	return solve(input, 2)
}

func solve(input string, part int) (string, error) {
	crabs, err := parse(input)
	if err != nil {
		return "", fmt.Errorf("solve part %d: %w", part, err)
	}
	if part == 1 {
		sort.Ints(crabs)
		median := crabs[len(crabs)/2]
		sum := 0
		for _, c := range crabs {
			if c < median {
				sum += median - c
			} else {
				sum += c - median
			}
		}
		return fmt.Sprint(sum), nil
	}
	crabSum := 0
	for _, c := range crabs {
		crabSum += c
	}
	meanFloor := crabSum / len(crabs)
	fuelSumFloor := 0
	for _, c := range crabs {
		fuelSumFloor += part2Fuel(c, meanFloor)
	}
	meanCeil := meanFloor + 1
	fuelSumCeil := 0
	for _, c := range crabs {
		fuelSumCeil += part2Fuel(c, meanCeil)
	}
	minFuel := fuelSumFloor
	if fuelSumCeil < fuelSumFloor {
		minFuel = fuelSumCeil
	}
	return fmt.Sprint(minFuel), nil
}

func parse(input string) ([]int, error) {
	var crabs []int
	for _, s := range strings.Split(strings.TrimSpace(input), ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("couldn't parse %q: %w", s, err)
		}
		crabs = append(crabs, i)
	}
	return crabs, nil
}

// part2Fuel returns the amount of fuel required to move to a position from a
// given starting position.
func part2Fuel(from, to int) int {
	min, max := from, to
	if min > max {
		min, max = max, min
	}
	d := max - min
	return (d * (d + 1)) / 2 // sum of the first d natural numbers
}
