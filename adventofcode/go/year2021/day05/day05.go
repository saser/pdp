package day05

// TODO(issues/30): consider using a slice-backed grid instead of a map-based
// one.

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Saser/pdp/adventofcode/go/geo"
)

func Part1(input string) (string, error) {
	return solve(input, 1)
}

func Part2(input string) (string, error) {
	return solve(input, 2)
}

func solve(input string, part int) (string, error) {
	segments, err := parse(input)
	if err != nil {
		return "", fmt.Errorf("solve part %d: %w", part, err)
	}
	marked := make(map[geo.Point]int) // coordinate -> how many lines cover it
	for _, s := range segments {
		switch {
		case !s.Horizontal() && !s.Vertical():
			if part == 1 {
				continue
			}
			dy := 1                // positive Y direction
			if s.From.Y > s.To.Y { // negative Y direction
				dy = -1
			}
			steps := s.To.X - s.From.X
			for i := 0; i <= steps; i++ {
				p := geo.Point{
					X: s.From.X + i,
					Y: s.From.Y + dy*i,
				}
				marked[p]++
			}
		case s.Horizontal():
			for x := s.From.X; x <= s.To.X; x++ {
				p := geo.Point{
					X: x,
					Y: s.From.Y,
				}
				marked[p]++
			}
		case s.Vertical():
			for y := s.From.Y; y <= s.To.Y; y++ {
				p := geo.Point{
					X: s.From.X,
					Y: y,
				}
				marked[p]++
			}
		}
	}
	coveredMultiple := 0
	for _, n := range marked {
		if n >= 2 {
			coveredMultiple++
		}
	}
	return fmt.Sprint(coveredMultiple), nil
}

type segment struct {
	From, To geo.Point
}

func parse(input string) ([]segment, error) {
	var segments []segment
	for i, line := range strings.Split(strings.TrimSpace(input), "\n") {
		s, err := parseLine(line)
		if err != nil {
			return nil, fmt.Errorf("parse line %d: %w", i+1, err)
		}
		segments = append(segments, s)
	}
	return segments, nil
}

// parseLine reads the line and parses a segment out of it. If the segment is
// parsed successfully it holds that:
//   * if the segment is horizontal, then From.X < To.X
//   * if the segment is vertical, then From.Y < To.Y
//   * if the segment is diagonal, then From.X < To.X
func parseLine(line string) (segment, error) {
	parts := strings.Split(line, " -> ")
	if len(parts) != 2 {
		return segment{}, fmt.Errorf(`line does not split into two parts around " -> ": %q`, line)
	}
	fromStr, toStr := parts[0], parts[1]

	fromParts := strings.Split(fromStr, ",")
	if len(fromParts) != 2 {
		return segment{}, fmt.Errorf(`first coordinate pair does not split into two parts around ",": %q`, fromStr)
	}
	x, err := strconv.Atoi(fromParts[0])
	if err != nil {
		return segment{}, fmt.Errorf(`parse X in first coordinate pair: %w`, err)
	}
	y, err := strconv.Atoi(fromParts[1])
	if err != nil {
		return segment{}, fmt.Errorf(`parse X in first coordinate pair: %w`, err)
	}
	from := geo.Point{
		X: x,
		Y: y,
	}

	toParts := strings.Split(toStr, ",")
	if len(toParts) != 2 {
		return segment{}, fmt.Errorf(`second coordinate pair does not split into two parts around ",": %q`, toStr)
	}
	x, err = strconv.Atoi(toParts[0])
	if err != nil {
		return segment{}, fmt.Errorf(`parse X in second coordinate pair: %w`, err)
	}
	y, err = strconv.Atoi(toParts[1])
	if err != nil {
		return segment{}, fmt.Errorf(`parse X in second coordinate pair: %w`, err)
	}
	to := geo.Point{
		X: x,
		Y: y,
	}

	swap := (from.X == to.X && from.Y > to.Y) || // vertical line going in negative Y direction
		(from.Y == to.Y && from.X > to.X) || // horizontal line going in negative X direction
		(from.X > to.X) // diagonal line going in negative X direction
	if swap {
		from, to = to, from
	}
	return segment{
		From: from,
		To:   to,
	}, nil
}

func (s segment) Horizontal() bool {
	return s.From.Y == s.To.Y
}

func (s segment) Vertical() bool {
	return s.From.X == s.To.X
}

func (s segment) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d", s.From.X, s.From.Y, s.To.X, s.To.Y)
}
