package day12

// TODO(issues/63): use memoization for the search function.

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
	count := g.CountPaths(part == 2)
	return fmt.Sprint(count), nil
}

type graph struct {
	edges      [][]int // cave -> neighboring caves
	small      []bool  // cave -> whether it is small
	start, end int     // ids for start and end caves, respecitvely

	// State kept for the path counting.
	visited   []int // cave -> how many times it has been visited
	twiceCave int   // cave that has been visited twice, or -1 if none
}

func parse(input string) *graph {
	stringEdges := make(map[string][]string) // cave names -> neighboring cave names
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		parts := strings.Split(line, "-")
		c1, c2 := parts[0], parts[1]
		stringEdges[c1] = append(stringEdges[c1], c2)
		stringEdges[c2] = append(stringEdges[c2], c1)
	}
	caveIDs := make(map[string]int) // cave name -> cave ID
	id := 0
	for cave := range stringEdges {
		caveIDs[cave] = id
		id++
	}
	edges := make([][]int, len(stringEdges))
	small := make([]bool, len(stringEdges))
	for cave, neighbors := range stringEdges {
		id := caveIDs[cave]
		small[id] = cave[0] >= 'a' && cave[0] <= 'z'
		for _, neighbor := range neighbors {
			edges[id] = append(edges[id], caveIDs[neighbor])
		}
	}
	return &graph{
		edges:     edges,
		small:     small,
		start:     caveIDs["start"],
		end:       caveIDs["end"],
		visited:   make([]int, len(edges)),
		twiceCave: -1,
	}
}

func (g *graph) CountPaths(allowTwice bool) int {
	count := 0
	g.visited[g.start]++
	g.countPaths(&count, g.start, allowTwice)
	g.visited[g.start]--
	return count
}

func (g *graph) countPaths(count *int, current int, allowTwice bool) {
	if current == g.end {
		*count++
		return
	}
	for _, neighbor := range g.edges[current] {
		// Part 1: only visit each cave at most once.
		if !allowTwice {
			// Skip any small caves which we have already visited.
			if g.small[neighbor] && g.visited[neighbor] > 0 {
				continue
			}
			g.visited[neighbor]++
			g.countPaths(count, neighbor, allowTwice)
			g.visited[neighbor]--
			continue
		}
		// Part 2: allow visiting small caves twice.
		var canVisit bool
		switch {
		case !g.small[neighbor]:
			// Any big cave can be visited.
			canVisit = true
		case neighbor == g.start || neighbor == g.end:
			// Start and end can only be visited once each.
			canVisit = g.visited[neighbor] < 1
		case neighbor == g.twiceCave:
			// A twice-visited cave cannot be visited again.
			canVisit = false
		case g.twiceCave == -1:
			// No cave has been visited twice, so this cave can always be
			// visited.
			canVisit = true
		default:
			// Some cave has been visited twice, so this cave can be visited
			// if it hasn't been visited yet.
			canVisit = g.visited[neighbor] < 1
		}
		if !canVisit {
			continue
		}
		g.visited[neighbor]++
		// If there is no twice-cave yet, and the cave we are visiting
		// is a small cave visited for the second time, let it take the
		// role of the twice-cave.
		if g.twiceCave == -1 && g.small[neighbor] && g.visited[neighbor] == 2 {
			g.twiceCave = neighbor
		}
		// Recursively count paths from there.
		g.countPaths(count, neighbor, allowTwice)
		// Back out.
		g.visited[neighbor]--
		// If the cave we are visiting was the twice-cave, backing out
		// leaves the role of twice-cave open, so set it to -1 to mark
		// it as such.
		if g.twiceCave == neighbor {
			g.twiceCave = -1
		}
	}
}
