package day02

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
	var instructions []instruction
	for i, line := range strings.Split(strings.TrimSpace(input), "\n") {
		ins, err := parseInstruction(line)
		if err != nil {
			return "", fmt.Errorf("parse input, line %d: %w", i+1, err)
		}
		instructions = append(instructions, ins)
	}
	if part == 1 {
		depth := 0
		horizontal := 0
		for _, i := range instructions {
			switch i.Direction {
			case directionForward:
				horizontal += i.Amount
			case directionDown:
				depth += i.Amount
			case directionUp:
				depth -= i.Amount
			}
		}
		return fmt.Sprint(depth * horizontal), nil
	}
	depth := 0
	horizontal := 0
	aim := 0
	for _, i := range instructions {
		switch i.Direction {
		case directionForward:
			horizontal += i.Amount
			depth += aim * i.Amount
		case directionDown:
			aim += i.Amount
		case directionUp:
			aim -= i.Amount
		}
	}
	return fmt.Sprint(depth * horizontal), nil
}

type direction int

const (
	directionForward direction = iota
	directionDown
	directionUp
)

func parseDirection(s string) (direction, error) {
	var (
		d   direction
		err error
	)
	switch s {
	case "forward":
		d = directionForward
	case "down":
		d = directionDown
	case "up":
		d = directionUp
	default:
		err = fmt.Errorf("parse direction: invalid direction %q", s)
	}
	return d, err
}

type instruction struct {
	Direction direction
	Amount    int
}

func parseInstruction(s string) (instruction, error) {
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return instruction{}, fmt.Errorf("parse instruction: invalid instruction %q", s)
	}
	d, err := parseDirection(parts[0])
	if err != nil {
		return instruction{}, fmt.Errorf("parse instruction: %w", err)
	}
	n, err := strconv.Atoi(parts[1])
	if err != nil {
		return instruction{}, fmt.Errorf("parse instruction: %w", err)
	}
	return instruction{
		Direction: d,
		Amount:    n,
	}, nil
}
