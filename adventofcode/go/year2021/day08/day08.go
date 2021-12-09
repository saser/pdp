package day08

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
	entries, err := parse(input)
	if err != nil {
		return "", fmt.Errorf("solve part %d: %w", part, err)
	}
	if part == 1 {
		count := 0
		for _, e := range entries {
			for _, d := range e.Output {
				switch d.Count() {
				case
					2, // 1
					4, // 4
					3, // 7
					7: // 8
					count++
				}
			}
		}
		return fmt.Sprint(count), nil
	}

	sum := 0
	for _, e := range entries {
		// numbers and patterns must be kept in sync -- if one is updated, then the
		// other must be updated as well.
		numbers := make(map[digit]int)  // bit pattern -> which number is displayed
		patterns := make(map[int]digit) // number -> which bit pattern represents it
		// First fill out numbers and patterns with the patterns of unique length,
		// representing numbers 1, 4, 7, and 8.
		for _, d := range e.Patterns {
			if len(numbers) == 4 {
				break
			}
			switch d.Count() {
			case 2: // number 1
				numbers[d] = 1
				patterns[1] = d
			case 4: // number 4
				numbers[d] = 4
				patterns[4] = d
			case 3: // number 7
				numbers[d] = 7
				patterns[7] = d
			case 7: // number 8
				numbers[d] = 8
				patterns[8] = d
			}
		}
		output := 0
		for _, d := range e.Output {
			var number int
			switch n := numbers[d]; n {
			case 1, 4, 7, 8:
				number = n
			default:
				and1 := d.And(patterns[1]).Count()
				and4 := d.And(patterns[4]).Count()
				switch n := d.Count(); {
				case n == 5 && and4 == 2:
					number = 2
				case n == 5 && and1 == 2:
					number = 3
				case n == 5 && and1 == 1:
					number = 5

				case n == 6 && and4 == 4:
					number = 9
				case n == 6 && and1 == 2:
					number = 0
				case n == 6 && and1 == 1:
					number = 6
				}
			}
			output = output*10 + number
		}
		sum += output
	}
	return fmt.Sprint(sum), nil
}

// digit represents a 7-segment display pattern as a bit mask. For example, the
// pattern "fab" is represented as follows:
//     0 0 1 0 0 0 1 1
//     _ g f e d c b a
// which in decimal notation is the number 35. The most significant bit is
// unused and will always be 0.
type digit int8

func parseDigit(s string) (digit, error) {
	if len(s) > 7 {
		return 0, fmt.Errorf("parse digit: too long: %q", s)
	}
	var d digit
	for _, r := range s {
		bit := r - 'a'
		d |= (1 << bit)
	}
	return d, nil
}

func (d digit) Count() int {
	n := 0
	for i := 0; i <= 7; i++ {
		if (d>>i)&1 == 1 {
			n++
		}
	}
	return n
}

func (d digit) And(other digit) digit {
	return d & other
}

type entry struct {
	Patterns [10]digit
	Output   [4]digit
}

func parseEntry(line string) (entry, error) {
	parts := strings.FieldsFunc(line, func(r rune) bool {
		return r == ' ' || r == '|'
	})
	var e entry
	if len(parts) != len(e.Patterns)+len(e.Output) {
		return entry{}, fmt.Errorf("parse entry: line does not contain %d + %d patterns: %q", len(e.Patterns), len(e.Output), line)
	}
	for i, p := range parts[:len(e.Patterns)] {
		d, err := parseDigit(p)
		if err != nil {
			return entry{}, err
		}
		e.Patterns[i] = d
	}
	for i, p := range parts[len(e.Patterns):] {
		d, err := parseDigit(p)
		if err != nil {
			return entry{}, err
		}
		e.Output[i] = d
	}
	return e, nil
}

func parse(input string) ([]entry, error) {
	var entries []entry
	for i, line := range strings.Split(strings.TrimSpace(input), "\n") {
		e, err := parseEntry(line)
		if err != nil {
			return nil, fmt.Errorf("parse line %d: %w", i+1, err)
		}
		entries = append(entries, e)
	}
	return entries, nil
}
