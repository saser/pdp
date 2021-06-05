package day09

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"

	"github.com/Saser/adventofcode/internal/permutations"
)

func Part1(r io.Reader) (string, error) {
	return solve(r, 1)
}

func Part2(r io.Reader) (string, error) {
	return solve(r, 2)
}

func solve(r io.Reader, part int) (string, error) {
	distances, err := parse(r)
	if err != nil {
		return "", fmt.Errorf("year 2015, day 09, part 1: %w", err)
	}
	places := make([]string, 0, len(distances))
	for k, _ := range distances {
		places = append(places, k)
	}
	routes := permutations.Strings(places)
	var compare func(int, int) int
	switch part {
	case 1:
		compare = func(a, b int) int {
			if a < b {
				return a
			} else {
				return b
			}
		}
	case 2:
		compare = func(a, b int) int {
			if a > b {
				return a
			} else {
				return b
			}
		}
	}
	answer := optimalDistance(routes, distances, compare)
	return fmt.Sprint(answer), nil
}

func parse(r io.Reader) (map[string]map[string]int, error) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)

	re, err := regexp.Compile(`^(\w+) to (\w+) = (\d+)$`)
	if err != nil {
		return nil, fmt.Errorf("parse: %w", err)
	}
	distances := make(map[string]map[string]int)
	for sc.Scan() {
		line := sc.Text()
		matches := re.FindStringSubmatch(line)
		if matches == nil {
			return nil, fmt.Errorf("parse: invalid line: %s", line)
		}
		from := matches[1]
		to := matches[2]
		distance, err := strconv.Atoi(matches[3])
		if err != nil {
			return nil, fmt.Errorf("parse: invalid distance: %s", matches[3])
		}
		if _, ok := distances[from]; !ok {
			distances[from] = make(map[string]int)
		}
		distances[from][to] = int(distance)
		if _, ok := distances[to]; !ok {
			distances[to] = make(map[string]int)
		}
		distances[to][from] = int(distance)
	}
	return distances, nil
}

func optimalDistance(routes [][]string, distances map[string]map[string]int, compare func(int, int) int) int {
	optimal := -1
	for _, route := range routes {
		distance := 0
		for i := 0; i < len(route)-1; i++ {
			from := route[i]
			to := route[i+1]
			distance += distances[from][to]
		}
		if optimal == -1 {
			optimal = distance
		} else {
			optimal = compare(optimal, distance)
		}
	}
	return optimal
}
