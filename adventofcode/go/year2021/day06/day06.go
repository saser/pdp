package day06

import (
	"fmt"
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
	fishes, err := parse(input)
	if err != nil {
		return "", fmt.Errorf("solve part %d: %w", part, err)
	}
	var timers [9]int64 // number of days until a new fish is created -> how many fishes have that number of days
	for _, f := range fishes {
		timers[f]++
	}
	days := 80
	if part == 2 {
		days = 256
	}
	for n := 0; n < days; n++ {
		tmp := timers[8]
		timers[8] = 0 // simulate "taking out" these fishes
		for i := 7; i >= 0; i-- {
			tmp, timers[i] = timers[i], tmp
		}
		// At this point, tmp holds the old value of timers[0]. Now tmp fishes are created at timers[8] and their parents loop back to timers[6].
		timers[8] += tmp
		timers[6] += tmp
	}
	sum := int64(0)
	for _, n := range timers {
		sum += n
	}
	return fmt.Sprint(sum), nil
}

func parse(input string) ([]int, error) {
	var fishes []int
	for _, s := range strings.Split(strings.TrimSpace(input), ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("couldn't parse %q: %w", s, err)
		}
		fishes = append(fishes, i)
	}
	return fishes, nil
}
