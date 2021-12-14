package day13

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Saser/pdp/adventofcode/go/geo"
	"github.com/Saser/pdp/adventofcode/go/intmath"
)

func Part1(input string) (string, error) {
	return solve(input, 1)
}

func Part2(input string) (string, error) {
	return solve(input, 2)
}

func solve(input string, part int) (string, error) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	g, err := parseGrid(parts[0])
	if err != nil {
		return "", fmt.Errorf("solve part %d: %w", part, err)
	}
	instructions, err := parseInstructions(parts[1])
	if err != nil {
		return "", fmt.Errorf("solve part %d: %w", part, err)
	}
	if part == 1 {
		instructions = instructions[0:1]
	}
	for _, i := range instructions {
		g.Fold(i)
	}
	if part == 1 {
		return fmt.Sprint(g.Len()), nil
	}
	return g.String(), nil
}

type instruction struct {
	Along string // either "x" or "y"
	Line  int    // the 5 in "fold along x=5", for example
}

func parseInstructions(s string) ([]instruction, error) {
	var is []instruction
	for _, line := range strings.Split(s, "\n") {
		parts := strings.Split(strings.TrimPrefix(line, "fold along "), "=")
		along := parts[0]
		n, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("parse instructions: %w", err)
		}
		is = append(is, instruction{
			Along: along,
			Line:  n,
		})
	}
	return is, nil
}

type grid struct {
	points map[geo.Point]bool
}

func parseGrid(s string) (*grid, error) {
	points := make(map[geo.Point]bool)
	for _, line := range strings.Split(s, "\n") {
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("parse grid: %w", err)
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("parse grid: %w", err)
		}
		points[geo.Point{X: x, Y: y}] = true
	}
	return &grid{
		points: points,
	}, nil
}

func (g *grid) Len() int {
	return len(g.points)
}

func (g *grid) String() string {
	var xs, ys []int
	for p := range g.points {
		xs = append(xs, p.X)
		ys = append(ys, p.Y)
	}
	maxX := intmath.Max(xs[0], xs[1:]...)
	maxY := intmath.Max(ys[0], ys[1:]...)
	var ss []string
	for y := 0; y <= maxY; y++ {
		var sb strings.Builder
		for x := 0; x <= maxX; x++ {
			if g.points[geo.Point{X: x, Y: y}] {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		ss = append(ss, sb.String())
	}
	return strings.Join(ss, "\n")
}

// It is assumed in Fold, foldX, and foldY, that we never do folds that result
// in negative x or y values. Why? The instructions are not entirely clear on
// what would happen if that were the case. They do not specifically state that
// folds only ever result in positive values, but seeing as all the instructions
// only are for non-negative values, it seems like a reasonable assumption.

func (g *grid) Fold(i instruction) {
	if i.Along == "x" {
		g.foldX(i.Line)
	} else {
		g.foldY(i.Line)
	}
}

func (g *grid) foldX(line int) {
	// "Fold" to the left by doing the following for each point:
	//     1. If the point is to the left of the fold line, ignore it.
	//     2. Calculate the new x-value of the point as
	//            new_x = old_x - 2*(old_x - line).
	//        The -1 is because we need to skip over the fold line.
	for p := range g.points {
		if p.X < line {
			continue
		}
		p2 := p
		p2.X = p.X - 2*(p.X-line)
		delete(g.points, p)
		g.points[p2] = true
	}
}

func (g *grid) foldY(line int) {
	// See the comment in foldX for an explanation of how this works. It is
	// basically the same but manipulating y-values instead.
	for p := range g.points {
		if p.Y < line {
			continue
		}
		p2 := p
		p2.Y = p.Y - 2*(p.Y-line)
		delete(g.points, p)
		g.points[p2] = true
	}
}
