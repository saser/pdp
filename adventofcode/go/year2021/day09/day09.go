package day09

// TODO(issues/47): replace maps with slices to hopefully improve runtime.

import (
	"fmt"
	"sort"
	"strings"
)

func Part1(input string) (string, error) {
	return solve(input, 1)
}

func Part2(input string) (string, error) {
	return solve(input, 2)
}

func solve(input string, part int) (string, error) {
	g, err := parse(input)
	if err != nil {
		return "", fmt.Errorf("solve part %d: %w", part, err)
	}
	sum := 0
	if part == 1 {
		for y := 0; y < g.Rows(); y++ {
		pointloop:
			for x := 0; x < g.Cols(); x++ {
				h, _ := g.Get(pos{X: x, Y: y})
			deltaloop:
				for _, d := range []struct {
					dx, dy int
				}{
					{dx: -1, dy: -1}, // top left
					{dx: 0, dy: -1},  // above
					{dx: +1, dy: -1}, // top right
					{dx: -1, dy: 0},  // left
					{dx: +1, dy: 0},  // right
					{dx: -1, dy: +1}, // bottom left
					{dx: 0, dy: +1},  // bottom
					{dx: +1, dy: +1}, // bottom right
				} {
					h2, ok := g.Get(pos{X: x + d.dx, Y: y + d.dy})
					if !ok {
						continue deltaloop
					}
					if h2 < h {
						continue pointloop
					}
				}
				// No adjacent point was lower, so this is a low point.
				sum += h + 1
			}
		}
		return fmt.Sprint(sum), nil
	}
	basins := make(map[pos]*int) // position -> pointer to number of positions in the same basin
	ps := make([]pos, 0, g.Rows()*g.Cols())
	for x := 0; x < g.Cols(); x++ {
		for y := 0; y < g.Rows(); y++ {
			ps = append(ps, pos{X: x, Y: y})
		}
	}
	for _, p := range ps {
		if _, ok := basins[p]; ok { // this position is already in a basin
			continue
		}
		if h, _ := g.Get(p); h == 9 { // this position will never be in a basin
			continue
		}
		n := 0
		queue := []pos{p}
		for {
			if len(queue) == 0 {
				break
			}
			// Pop the queue.
			pp := queue[0]
			queue = queue[1:]
			// Skip if we have already seen this position.
			if _, ok := basins[pp]; ok {
				continue
			}
			// We have a new position in the same basin.
			n++
			basins[pp] = &n
			// Find all neighbors of the position. If their value is 9, or if
			// they're already in a basin, skip it.
			for _, d := range []struct {
				dx, dy int
			}{
				{dx: 0, dy: -1}, // above
				{dx: -1, dy: 0}, // left
				{dx: +1, dy: 0}, // right
				{dx: 0, dy: +1}, // bottom
			} {
				neighbor := pos{X: pp.X + d.dx, Y: pp.Y + d.dy}
				if _, ok := basins[neighbor]; ok {
					continue
				}
				if n, ok := g.Get(neighbor); !ok || n == 9 {
					continue
				}
				queue = append(queue, neighbor)
			}
		}
	}
	seen := make(map[*int]bool) // a set of seen "basins", i.e., pointers to ints with the size of the basin
	var sizes []int
	for _, b := range basins {
		if seen[b] {
			continue
		}
		seen[b] = true
		sizes = append(sizes, *b)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	return fmt.Sprint(sizes[0] * sizes[1] * sizes[2]), nil
}

type grid struct {
	heights            []int
	rowCount, colCount int
}

func parse(input string) (*grid, error) {
	var heights []int
	rowCount := 0
	colCount := 0
	for i, r := range strings.TrimSpace(input) {
		if r == '\n' {
			rowCount++
			if colCount == 0 {
				colCount = i
			}
			continue
		}
		if r < '0' || r > '9' {
			return nil, fmt.Errorf("parse: invalid rune: %q", r)
		}
		heights = append(heights, int(r-'0'))
	}
	rowCount++ // we trimmed the last \n off of the input
	return &grid{
		heights:  heights,
		rowCount: rowCount,
		colCount: colCount,
	}, nil
}

type pos struct {
	X, Y int
}

func (g *grid) Get(p pos) (int, bool) {
	if p.X < 0 || p.X >= g.colCount || p.Y < 0 || p.Y >= g.rowCount {
		return 0, false
	}
	return g.heights[p.Y*g.colCount+p.X], true
}

func (g *grid) Rows() int {
	return g.rowCount
}

func (g *grid) Cols() int {
	return g.colCount
}
