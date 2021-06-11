package day19

import (
	"testing"

	"github.com/Saser/pdp/adventofcode/go/internal/testcase"
)

var (
	part1ExampleFile = testcase.Runfile("adventofcode/go/internal/year2020/day19/testdata/p1example")
	part2ExampleFile = testcase.Runfile("adventofcode/go/internal/year2020/day19/testdata/p2example")
	inputFile        = testcase.Runfile("adventofcode/inputs/2020/19")

	tcPart1 = testcase.NewFile("input", inputFile, "111")
	tcPart2 = testcase.NewFile("input", inputFile, "343")
)

func TestPart1(t *testing.T) {
	for _, tc := range []testcase.TestCase{
		testcase.NewFile("example", part1ExampleFile, "2"),
		tcPart1,
	} {
		tc.Test(t, Part1)
	}
}

func BenchmarkPart1(b *testing.B) {
	tcPart1.Benchmark(b, Part1)
}

func TestPart2(t *testing.T) {
	for _, tc := range []testcase.TestCase{
		testcase.NewFile("example", part2ExampleFile, "12"),
		tcPart2,
	} {
		tc.Test(t, Part2)
	}
}

func BenchmarkPart2(b *testing.B) {
	tcPart2.Benchmark(b, Part2)
}
