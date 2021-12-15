package day15

import (
	"container/heap"
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
	g.Expanded = part == 2
	start := 0
	end := g.Rows()*g.Cols() - 1
	cost := dijkstra(g, start, end)
	return fmt.Sprint(cost), nil
}

type grid struct {
	risks      []int // row-major ordering of risks
	rows, cols int   // size of the grid

	Expanded bool
}

func parse(input string) *grid {
	var risks []int
	rows := 0
	cols := 0
	for i, r := range strings.TrimSpace(input) {
		if r == '\n' {
			rows++
			if cols == 0 {
				cols = i
			}
			continue
		}
		risks = append(risks, int(r-'0'))
	}
	rows++ // we trimmed off the last newline, if any
	return &grid{
		risks: risks,
		rows:  rows,
		cols:  cols,
	}
}

func (g *grid) Cols() int {
	c := g.cols
	if g.Expanded {
		c *= 5
	}
	return c
}

func (g *grid) Rows() int {
	r := g.rows
	if g.Expanded {
		r *= 5
	}
	return r
}

func (g *grid) Get(i int) int {
	// Calculate first the "virtual" x and y coordinates, i.e., the x and y
	// coordinates in a possibly expanded map.
	vx := i % g.Cols()
	vy := i / g.Cols()
	// Calculate the "physical" x and y coordinates. These are the ones that
	// correspond to a number in the actual grid. They may be the same as the
	// virtual coordinates.
	x := vx % g.cols
	y := vy % g.rows
	// Calculate which tile the virtual coordinates lie in. Which tile
	// determines how the physical value should be manipulated to account for
	// the increasing and wrapping risk levels.
	tileX := vx / g.cols
	tileY := vy / g.rows
	// Find the physical value.
	v := g.risks[y*g.cols+x]
	// Each tile-step to the right and down increase the value by 1.
	v += tileX + tileY
	// If the value has increased over 9, it should be wrapped around starting
	// from 1. This makes the range of possible values [1..9], which is 1 more
	// than the range of possible values mod 9, i.e., [0..8]. So we "translate"
	// v into the [0..8] range by decreasing by 1, then taking modulo 9, and
	// then "translate" back into the [1..9] range by increasing by 1.
	v = (v-1)%9 + 1
	return v
}

type item struct {
	Point int
	Cost  int
}

type priorityQueue []*item

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i int, j int) bool {
	return pq[i].Cost < pq[j].Cost
}

func (pq priorityQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*item))
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func dijkstra(g *grid, start, end int) int {
	rows := g.Rows()
	cols := g.Cols()
	max := rows*cols - 1

	shortest := make([]int, rows*cols) // point -> cost of cheapest path from start to that point

	var pq priorityQueue
	heap.Init(&pq)
	heap.Push(&pq, &item{
		Point: start,
		Cost:  0,
	})
	for len(pq) != 0 {
		i := heap.Pop(&pq).(*item)
		current := i.Point
		cost := i.Cost
		if current == end {
			return cost
		}
		shortest[current] = cost
		currentX := current % g.Cols()
		for _, d := range []int{
			-1,    // left
			+1,    // right
			-cols, // up
			+cols, // down
		} {
			neighbor := current + d
			if neighbor < 0 || neighbor > max {
				continue
			}
			neighborX := neighbor % cols
			if d < 0 && neighborX > currentX || d > 0 && neighborX < currentX {
				continue
			}
			if shortest[neighbor] == 0 {
				neighborCost := cost + g.Get(neighbor)
				shortest[neighbor] = neighborCost
				heap.Push(&pq, &item{
					Point: neighbor,
					Cost:  neighborCost,
				})
			}
		}
	}
	return -1
}
