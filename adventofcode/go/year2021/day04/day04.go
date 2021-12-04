package day04

// TODO(issues/28): try using maps much more here to see if it becomes faster.

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
	paragraphs := strings.Split(input, "\n\n")
	var numbers []int
	for _, s := range strings.Split(paragraphs[0], ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			return "", fmt.Errorf("parse bingo numbers: %w", err)
		}
		numbers = append(numbers, i)
	}
	var boards []*board
	for _, p := range paragraphs[1:] {
		b, err := parseBoard(p)
		if err != nil {
			return "", fmt.Errorf("solve part %v: %w", part, err)
		}
		boards = append(boards, b)
	}
	completed := make(map[*board]bool) // set of completed boards
	for _, n := range numbers {
		for _, b := range boards {
			if completed[b] {
				continue
			}
			if b.Mark(n) {
				if sum, bingo := b.Bingo(); bingo {
					if part == 1 {
						return fmt.Sprint(sum * n), nil
					}
					completed[b] = true
					if len(completed) == len(boards) {
						return fmt.Sprint(sum * n), nil
					}
				}
			}
		}
	}
	return "", fmt.Errorf("no solution found for part %d, this is a bug", part)
}

type board struct {
	grid        [5][5]int  // row -> column -> number
	marked      [5][5]bool // row -> column -> marked or not
	unmarkedSum int        // sum of all unmarked numbers, updated by Mark
}

func parseBoard(s string) (*board, error) {
	var grid [5][5]int // row -> column -> number
	sum := 0
	for i, f := range strings.Fields(s) {
		n, err := strconv.Atoi(f)
		if err != nil {
			return nil, fmt.Errorf("parse board: %w", err)
		}
		grid[i/5][i%5] = n
		sum += n
	}
	return &board{
		grid:        grid,
		unmarkedSum: sum,
	}, nil
}

func (b *board) Mark(n int) bool {
	for row := 0; row < len(b.grid); row++ {
		for col := 0; col < len(b.grid[row]); col++ {
			if b.grid[row][col] == n {
				b.marked[row][col] = true
				b.unmarkedSum -= n
				return true
			}
		}
	}
	return false
}

func (b *board) Bingo() (int, bool) {
	// First check if there's bingo on a row.
rowLoop:
	for row := 0; row < len(b.grid); row++ {
		for col := 0; col < len(b.grid[row]); col++ {
			if !b.marked[row][col] {
				continue rowLoop
			}
		}
		// If we reach this point then all numbers in the row have been marked,
		// and we have bingo.
		return b.unmarkedSum, true
	}

	// No bingo in any row, now let's check columns.
colLoop:
	for col := 0; col < len(b.grid[0]); col++ {
		for row := 0; row < len(b.grid); row++ {
			if !b.marked[row][col] {
				continue colLoop
			}
		}
		// If we reach this point then all numbers in the column have been
		// marked, and we have bingo.
		return b.unmarkedSum, true
	}

	return 0, false
}
