package day17

import (
	"fmt"
	"regexp"
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
	xmin, xmax, ymin, ymax := parse(strings.TrimSpace(input))
	if part == 1 {
		// Assume that the target is beneath the starting position.
		vy := -ymin - 1
		return fmt.Sprint((vy * (vy + 1)) / 2), nil
	}
	_ = xmin
	_ = xmax
	_ = ymin
	_ = ymax
	return "", fmt.Errorf("solution not implemented for part %v", part)
}

var (
	inputRE   = regexp.MustCompile(`target area: x=(?P<xmin>-?\d+)..(?P<xmax>-?\d+), y=(?P<ymin>-?\d+)..(?P<ymax>-?\d+)`)
	xminIndex = inputRE.SubexpIndex("xmin")
	xmaxIndex = inputRE.SubexpIndex("xmax")
	yminIndex = inputRE.SubexpIndex("ymin")
	ymaxIndex = inputRE.SubexpIndex("ymax")
)

func parse(input string) (xmin, xmax, ymin, ymax int) {
	matches := inputRE.FindStringSubmatch(input)
	var err error
	xmin, err = strconv.Atoi(matches[xminIndex])
	if err != nil {
		panic(err)
	}
	xmax, err = strconv.Atoi(matches[xmaxIndex])
	if err != nil {
		panic(err)
	}
	ymin, err = strconv.Atoi(matches[yminIndex])
	if err != nil {
		panic(err)
	}
	ymax, err = strconv.Atoi(matches[ymaxIndex])
	if err != nil {
		panic(err)
	}
	return xmin, xmax, ymin, ymax
}
