package day01

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
	depths, err := parse(input)
	if err != nil {
		return "", fmt.Errorf("solve part %d: %w", part, err)
	}
	windowSize := 1
	if part == 2 {
		windowSize = 3
	}
	increasingCount := 0
	for i := 0; i < len(depths)-windowSize; i++ {
		sum1 := 0
		for _, d := range depths[i : i+windowSize] {
			sum1 += d
		}
		sum2 := 0
		for _, d := range depths[i+1 : i+1+windowSize] {
			sum2 += d
		}
		if sum1 < sum2 {
			increasingCount++
		}
	}
	return fmt.Sprint(increasingCount), nil
}

func parse(input string) ([]int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var depths []int
	for i, line := range lines {
		d, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("parse line %d: %w", i+1, err)
		}
		depths = append(depths, d)
	}
	return depths, nil
}
