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
		return fmt.Sprint(g.ActiveCount()), nil
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
	dots       []bool // laid out row-wise
	rows, cols int
}

func parseGrid(s string) (*grid, error) {
	var ps []geo.Point
	xmax := -1
	ymax := -1
	for _, line := range strings.Split(s, "\n") {
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("parse grid: %w", err)
		}
		if x > xmax {
			xmax = x
		}
		y, err := strconv.Atoi(parts[1])
		if y > ymax {
			ymax = y
		}
		if err != nil {
			return nil, fmt.Errorf("parse grid: %w", err)
		}
		ps = append(ps, geo.Point{
			X: x,
			Y: y,
		})
	}
	// xmax and ymax are one less than the number of columns and rows,
	// respectively.
	// Example: if xmax = 4, then the largest x index is 4, which means 5
	// columns.
	cols := xmax + 1
	rows := ymax + 1
	dots := make([]bool, rows*cols)
	g := &grid{
		dots: dots,
		rows: rows,
		cols: cols,
	}
	for _, p := range ps {
		i, _ := g.index(p.X, p.Y)
		g.dots[i] = true
	}
	return g, nil
}

func (g *grid) index(x, y int) (int, bool) {
	if x < 0 || x >= g.cols || y < 0 || y >= g.rows {
		return 0, false
	}
	return y*g.cols + x, true
}

func (g *grid) Get(x, y int) (value bool, ok bool) {
	i, ok := g.index(x, y)
	if !ok {
		return false, false
	}
	return g.dots[i], ok
}

func (g *grid) Set(x, y int, value bool) bool {
	i, ok := g.index(x, y)
	if !ok {
		return false
	}
	g.dots[i] = value
	return true
}

func (g *grid) ActiveCount() int {
	n := 0
	for _, b := range g.dots {
		if b {
			n++
		}
	}
	return n
}

func (g *grid) String() string {
	var ss []string
	for y := 0; y < g.rows; y++ {
		var sb strings.Builder
		for x := 0; x < g.cols; x++ {
			r := '.'
			if v, _ := g.Get(x, y); v {
				r = '#'
			}
			sb.WriteRune(r)
		}
		ss = append(ss, sb.String())
	}
	return strings.Join(ss, "\n")
}

func (g *grid) Fold(i instruction) {
	if i.Along == "x" {
		g.foldX(i.Line)
	} else {
		g.foldY(i.Line)
	}
}

// The code for folding is a little unintuitive, so this is an summary of how it
// works, with examples to help illustrate. The examples will be for folding to
// the left ("along x") but the same principles apply for folding up ("along
// y").
//
// Example 1: 7 columns, fold along x=4. There is only a two rows in this
// example, to make it a bit smaller.
//
//     #...|.#
//     #..#|#.
//
// Folding along a line splits the grid in two parts: one unchanged part and one
// flipped part. The unchanged part is the one that ends up "below" the flipped
// part in the fold.
//
//     #... | .#
//     #..# | #.
//     ^^^^   ^^
//     ||||   || flipped part
//     |||| unchanged part
//
// The flipped part is reversed in place.
//
//     #... | .#  becomes  #... | #.
//     #..# | #.           #..# | .#
//
// The overlapping columns in the unchanged part and the flipped part are ORed
// together and stored in a new, empty grid with the right size. That is done by
// iterating over the number of overlapping columns, which is calculated as the
// minimum of "number of columns in unchanged part" and "number of columns in
// flipped part". In this case, that number is 2. Starting from the folding
// line, the iteration starts 2 columns to the left, and 1 column to the right.
// Then the iteration happens in lockstep.
//
// First loop of iteration:
//
//     #.[.]. | [#].  becomes  ..[#].
//     #.[.]# | [.]#           ..[.].
//     old grid                new grid
//
// Second loop of iteration:
//
//     #..[.] | #[.]  becomes  ..#[.]
//     #..[#] | .[#]           ...[#]
//     old grid                new grid
//
// What remains is filling in the initial part of the new grid. In this case,
// the unchanged part is larger than the flipped part, so the values are taken
// from the unchanged part. If the flipped part was larger, the values would be
// taken from the flipped part.
//
// First loop of iteration:
//
//     [#]...|#.  becomes  [#].#.
//     [#]..#|.#           [#]..#
//     old grid            new grid
//
// Second loop of iteration:
//
//     #[.]..|#.  becomes  #[.]#.
//     #[.].#|.#           #[.].#
//     old grid            new grid
//
// The result of the fold is then:
//
//     #.#.
//     #..#
//
// ----
//
// Example 2: 7 columns, fold along n=2.
//
//     #.|...#
//     #.|#.#.
//
// Reverse the flipped part.
//
//     #.|...#  becomes  #.|#...
//     #.|#.#.           #.|.#.#
//
// OR together the overlapping part, store in a new grid. Two things to note:
//     * the iteration starts 2 steps from the end of the flipped part, not 1
//       step to the right of the line (in the previous example, those were the
//       same thing).
//     * the result is again stored in the end of the new grid.
//
// First loop of iteration:
//
//     [#]. | #.[.].  becomes  ..[#].
//     [#]. | .#[.]#           ..[#].
//     old grid                new grid
//
// Second loop of iteration:
//
//     #[.] | #..[.]  becomes  ..#[.]
//     #[.] | .#.[#]           ..#[#]
//     old grid                new grid
//
// Finally, store the initial values, this time of the flipped part as it is
// larger than the unchanged part.
//
// First loop of iteration:
//
//     #.|[#]...  becomes  [#].#.
//     #.|[.]#.#           [.].##
//     old grid            new grid
//
// Second loop of iteration:
//
//     #.|#[.]..  becomes  #[.]#.
//     #.|.[#].#           .[#]##
//     old grid            new grid
//
// The result of the fold:
//
//     #.#.
//     .###

func (g *grid) foldX(n int) {
	// Reverse the flipped part first.
	for y := 0; y < g.rows; y++ {
		// Use the fact that g.dots is row-major: create a slice representing
		// the partial row that we want to reverse, and then reverse only that
		// slice.
		start := y*g.cols + (n + 1) // start of row + first column of reversed part
		end := (y + 1) * g.cols     // start of next row
		partial := g.dots[start:end]
		for i, j := 0, len(partial)-1; i < j; i, j = i+1, j-1 {
			partial[i], partial[j] = partial[j], partial[i]
		}
	}

	// Create a new grid of appropriate size.
	unchangedCols := n              // Everything up to but not including the fold line.
	flippedCols := g.cols - (n + 1) // Skip the first n columns, plus the fold line.
	rows2 := g.rows
	cols2 := intmath.Max(unchangedCols, flippedCols)
	dots2 := make([]bool, rows2*cols2)
	g2 := &grid{
		dots: dots2,
		rows: rows2,
		cols: cols2,
	}

	// OR together unchanged and flipped part.
	overlappingCols := intmath.Min(unchangedCols, flippedCols)
	for y := 0; y < g.rows; y++ {
		for i := 0; i < overlappingCols; i++ {
			unchangedX := n - overlappingCols + i    // Relative to fold line.
			flippedX := g.cols - overlappingCols + i // Relative to end of flipped part.
			x2 := g2.cols - overlappingCols + i      // Relative to end of new grid.
			unchangedV, _ := g.Get(unchangedX, y)
			flippedV, _ := g.Get(flippedX, y)
			g2.Set(x2, y, unchangedV || flippedV)
		}
	}

	// Copy over initial part of either unchanged or flipped part.
	copyCols := g2.cols - overlappingCols
	// Assume the unchanged part is bigger -- then we start copying from the
	// first column.
	copyStart := 0
	if flippedCols > unchangedCols {
		// The flipped part is bigger, so we start copying from 1 step right of
		// the fold line.
		copyStart = n + 1
	}
	for y := 0; y < g.rows; y++ {
		for i := 0; i < copyCols; i++ {
			srcX := copyStart + i
			dstX := i
			v, _ := g.Get(srcX, y)
			g2.Set(dstX, y, v)
		}
	}

	// We are done, so overwrite this grid with the result of the fold.
	*g = *g2
}

func (g *grid) foldY(n int) {
	// Reverse the flipped part first.
	for x := 0; x < g.cols; x++ {
		for yi, yj := n+1, g.rows-1; yi < yj; yi, yj = yi+1, yj-1 {
			ii, _ := g.index(x, yi)
			ij, _ := g.index(x, yj)
			g.dots[ii], g.dots[ij] = g.dots[ij], g.dots[ii]
		}
	}

	// Create a new grid of appropriate size.
	unchangedRows := n              // Everything up to but not including the fold line.
	flippedRows := g.rows - (n + 1) // Skip the first n columns, plus the fold line.
	rows2 := intmath.Max(unchangedRows, flippedRows)
	cols2 := g.cols
	dots2 := make([]bool, rows2*cols2)
	g2 := &grid{
		dots: dots2,
		rows: rows2,
		cols: cols2,
	}

	// OR together unchanged and flipped part.
	overlappingRows := intmath.Min(unchangedRows, flippedRows)
	for x := 0; x < g.cols; x++ {
		for i := 0; i < overlappingRows; i++ {
			unchangedY := n - overlappingRows + i    // Relative to fold line.
			flippedY := g.rows - overlappingRows + i // Relative to end of flipped part.
			y2 := g2.rows - overlappingRows + i      // Relative to end of new grid.
			unchangedV, _ := g.Get(x, unchangedY)
			flippedV, _ := g.Get(x, flippedY)
			g2.Set(x, y2, unchangedV || flippedV)
		}
	}

	// Copy over initial part of either unchanged or flipped part.
	copyRows := g2.rows - overlappingRows
	// Assume the unchanged part is bigger -- then we start copying from the
	// first row.
	copyStart := 0
	if flippedRows > unchangedRows {
		// The flipped part is bigger, so we start copying from 1 step below the
		// fold line.
		copyStart = n + 1
	}
	for x := 0; x < g.cols; x++ {
		for i := 0; i < copyRows; i++ {
			srcY := copyStart + i
			dstY := i
			v, _ := g.Get(x, srcY)
			g2.Set(x, dstY, v)
		}
	}

	// We are done, so overwrite this grid with the result of the fold.
	*g = *g2
}
