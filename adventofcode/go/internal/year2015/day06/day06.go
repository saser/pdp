package day06

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"regexp"
	"strconv"

	"github.com/Saser/adventofcode/internal/geo"
)

func Part1(r io.Reader) (string, error) {
	instructions, err := parse(r, 1)
	if err != nil {
		return "", fmt.Errorf("year 2015, day 06, part 1: %w", err)
	}
	return solve(instructions)
}

func Part2(r io.Reader) (string, error) {
	instructions, err := parse(r, 2)
	if err != nil {
		return "", fmt.Errorf("year 2015, day 06, part 1: %w", err)
	}
	return solve(instructions)
}

func solve(instructions []instruction) (string, error) {
	grid := make([][]int, 0, 1000)
	for x := 0; x < 1000; x++ {
		grid = append(grid, make([]int, 1000))
	}
	for _, instruction := range instructions {
		instruction.apply(grid)
	}
	count := 0
	for _, row := range grid {
		for _, light := range row {
			count += light
		}
	}
	return fmt.Sprint(count), nil
}

type operation func(int) int

type instruction struct {
	start geo.Point
	end   geo.Point
	op    operation
}

func parse(r io.Reader, part int) ([]instruction, error) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)
	re, err := regexp.Compile(`(turn on|turn off|toggle) (\d{1,3}),(\d{1,3}) through (\d{1,3}),(\d{1,3})`)
	if err != nil {
		return nil, fmt.Errorf("parse: %w", err)
	}
	instructions := make([]instruction, 0)
	for sc.Scan() {
		line := sc.Text()
		matches := re.FindStringSubmatch(line)
		var op operation
		switch matches[1] {
		case "turn on":
			if part == 1 {
				op = turnOn
			} else {
				op = increaseBy(1)
			}
		case "turn off":
			if part == 1 {
				op = turnOff
			} else {
				op = decrease
			}
		case "toggle":
			if part == 1 {
				op = toggle
			} else {
				op = increaseBy(2)
			}
		}
		startX, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, fmt.Errorf("parse: %w", err)
		}
		startY, err := strconv.Atoi(matches[3])
		if err != nil {
			return nil, fmt.Errorf("parse: %w", err)
		}
		start := geo.Point{X: startX, Y: startY}
		endX, err := strconv.Atoi(matches[4])
		if err != nil {
			return nil, fmt.Errorf("parse: %w", err)
		}
		endY, err := strconv.Atoi(matches[5])
		if err != nil {
			return nil, fmt.Errorf("parse: %w", err)
		}
		end := geo.Point{X: endX, Y: endY}
		instructions = append(instructions, instruction{start: start, end: end, op: op})
	}
	return instructions, nil
}

func (d *instruction) apply(grid [][]int) {
	xMin := int(math.Min(float64(d.start.X), float64(d.end.X)))
	xMax := int(math.Max(float64(d.start.X), float64(d.end.X)))
	yMin := int(math.Min(float64(d.start.Y), float64(d.end.Y)))
	yMax := int(math.Max(float64(d.start.Y), float64(d.end.Y)))
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			grid[x][y] = d.op(grid[x][y])
		}
	}
}

func turnOn(_ int) int {
	return 1
}

func turnOff(_ int) int {
	return 0
}

func toggle(i int) int {
	if i == 0 {
		return 1
	} else {
		return 0
	}
}

func increaseBy(d int) operation {
	return func(i int) int {
		return i + d
	}
}

func decrease(i int) int {
	return int(math.Max(0, float64(i-1)))
}

func decrease2(i int) int {
	if i == 0 {
		return 0
	} else {
		return i - 1
	}
}
