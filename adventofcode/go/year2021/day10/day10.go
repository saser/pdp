package day10

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
	sum := 0
	var scores []int
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		stack, err := parse(line)
		if err == nil {
			continue
		}
		if _, ok := err.(*incompleteError); ok {
			if part == 1 {
				continue
			}
			score := 0
			for i := len(stack) - 1; i >= 0; i-- {
				var points int
				switch stack[i] {
				case ')':
					points = 1
				case ']':
					points = 2
				case '}':
					points = 3
				case '>':
					points = 4
				}
				score = score*5 + points
			}
			scores = append(scores, score)
		}
		if e, ok := err.(*corruptError); ok {
			if part == 2 {
				continue
			}
			var points int
			switch line[e.Index] {
			case ')':
				points = 3
			case ']':
				points = 57
			case '}':
				points = 1197
			case '>':
				points = 25137
			}
			sum += points
		}
	}
	if part == 1 {
		return fmt.Sprint(sum), nil
	}
	sort.Ints(scores)
	return fmt.Sprint(scores[len(scores)/2]), nil
}

type incompleteError struct {
	Line string
}

func (e *incompleteError) Error() string {
	return fmt.Sprintf("incomplete line %q", e.Line)
}

type corruptError struct {
	Line  string
	Index int
}

func (e *corruptError) Error() string {
	return fmt.Sprintf("corrupt line %q: illegal character %q at index %d", e.Line, e.Line[e.Index], e.Index)
}

func parse(line string) ([]rune, error) {
	// The stack contains the expected closing delimiters. It is preallocated to
	// make it a bit faster. It can never contain more runes than there are in
	// the input.
	stack := make([]rune, 0, len(line))
	for i, r := range line {
		isOpener := false
		switch r {
		case '(', '[', '{', '<':
			isOpener = true
		}
		if isOpener {
			var closer rune
			switch r {
			case '(':
				closer = ')'
			case '[':
				closer = ']'
			case '{':
				closer = '}'
			case '<':
				closer = '>'
			}
			stack = append(stack, closer)
			continue
		}
		if r != stack[len(stack)-1] {
			return stack, &corruptError{
				Line:  line,
				Index: i,
			}
		}
		stack = stack[:len(stack)-1]
	}
	if len(stack) != 0 {
		return stack, &incompleteError{
			Line: line,
		}
	}
	return nil, nil
}
