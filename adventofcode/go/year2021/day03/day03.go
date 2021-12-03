package day03

import (
	"fmt"
	"strings"
)

func Part1(input string) (string, error) {
	counts := make(map[int]int) // position -> count. First position == 0 == left-most bit
	pos := 0
	lines := 0
	for _, r := range input {
		switch r {
		case '\n':
			pos = 0
			lines++
			continue
		case '1':
			counts[pos]++
		}
		pos++
	}
	if input[len(input)-1] != '\n' {
		lines++
	}
	gamma := 0
	epsilon := 0
	for pos := 0; pos < len(counts); pos++ {
		mostCommon := 0
		if counts[pos] > lines-counts[pos] {
			mostCommon = 1
		}
		gamma = (gamma << 1) + mostCommon
		epsilon = (epsilon << 1) + (1 - mostCommon)
	}
	return fmt.Sprint(gamma * epsilon), nil
}

// TODO(issues/24): replace this with a binary tree backed by a slice.
func Part2(input string) (string, error) {
	t := new(tree)
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		t.Add(line)
	}
	oxygen := 0
	n := t.Root
	for {
		if n.Zero == nil && n.One == nil {
			break
		}
		oxygen <<= 1
		if n.Zero != nil && n.One != nil {
			if n.Zero.Count > n.One.Count {
				n = n.Zero
			} else {
				oxygen += 1
				n = n.One
			}
		} else if n.Zero != nil {
			n = n.Zero
		} else {
			oxygen += 1
			n = n.One
		}
	}
	co2 := 0
	n = t.Root
	for {
		if n.Zero == nil && n.One == nil {
			break
		}
		co2 <<= 1
		if n.Zero != nil && n.One != nil {
			if n.Zero.Count <= n.One.Count {
				n = n.Zero
			} else {
				co2 += 1
				n = n.One
			}
		} else if n.Zero != nil {
			n = n.Zero
		} else {
			co2 += 1
			n = n.One
		}
	}
	return fmt.Sprint(oxygen * co2), nil
}

// node and tree below are used in a stupid prefix tree implementation. It's
// probably way more complicated than it needs to be.

type node struct {
	Count     int
	Zero, One *node
}

type tree struct {
	Root *node
}

func (t *tree) Add(s string) {
	if t.Root == nil {
		t.Root = new(node)
	}
	add(t.Root, s)
}

func add(n *node, s string) {
	n.Count++
	if s == "" {
		return
	}
	var next *node
	switch s[0] {
	case '0':
		if n.Zero == nil {
			n.Zero = new(node)
		}
		next = n.Zero
	case '1':
		if n.One == nil {
			n.One = new(node)
		}
		next = n.One
	}
	add(next, s[1:])
}
