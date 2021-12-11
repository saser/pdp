package day11

import (
	"fmt"
	"strings"
)

func Part1(input string) (string, error) {
	return solve(input, 1)
}

func Part2(input string) (string, error) {
	return solve(input, 2)
}

func solve(input string, part int) (string, error) {
	g := parse(input)
	if part == 1 {
		flashCount := 0
		for i := 0; i < 100; i++ {
			flashCount += g.Step()
		}
		return fmt.Sprint(flashCount), nil
	}
	step := 0
	for {
		n := g.Step()
		step++
		if n == 100 {
			break
		}
	}
	return fmt.Sprint(step), nil
}

type grid struct {
	g [10][10]int
}

type pos struct {
	Row, Col int
}

func parse(input string) *grid {
	var g [10][10]int
	for row, line := range strings.Split(strings.TrimSpace(input), "\n") {
		for col, r := range line {
			g[row][col] = int(r - '0')
		}
	}
	return &grid{
		g: g,
	}
}

func (g *grid) Step() int {
	var flashed [10][10]bool
	flashCount := 0
	queue := make([]pos, 0, 100)
	for row := range g.g {
		for col := range g.g[row] {
			g.g[row][col]++
			if g.g[row][col] > 9 {
				queue = append(queue, pos{
					Row: row,
					Col: col,
				})
			}
		}
	}
	reset := make([]pos, 0, 100)
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if flashed[p.Row][p.Col] {
			continue
		}
		flashed[p.Row][p.Col] = true
		flashCount++
		reset = append(reset, p)
		for _, d := range []pos{
			{Row: -1, Col: -1},
			{Row: -1, Col: 0},
			{Row: -1, Col: +1},
			{Row: 0, Col: -1},
			{Row: 0, Col: +1},
			{Row: +1, Col: -1},
			{Row: +1, Col: 0},
			{Row: +1, Col: +1},
		} {
			row := p.Row + d.Row
			col := p.Col + d.Col
			if row < 0 || row > 9 || col < 0 || col > 9 {
				continue
			}
			g.g[row][col]++
			if g.g[row][col] > 9 && !flashed[row][col] {
				queue = append(queue, pos{
					Row: row,
					Col: col,
				})
			}
		}
	}
	for _, p := range reset {
		g.g[p.Row][p.Col] = 0
	}
	return flashCount
}
