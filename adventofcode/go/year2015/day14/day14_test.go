package day14

import (
	"testing"

	"github.com/Saser/pdp/adventofcode/go/testcase"
)

var inputFile = testcase.Runfile("adventofcode/data/year2015/day14/actual.in")

var (
	tcPart1 = testcase.NewFile("input", inputFile, "2696")
	tcPart2 = testcase.NewFile("input", inputFile, "1084")
)

func TestPart1(t *testing.T) {
	for _, tc := range []testcase.TestCase{
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
		tcPart2,
	} {
		tc.Test(t, Part2)
	}
}

func BenchmarkPart2(b *testing.B) {
	tcPart2.Benchmark(b, Part2)
}
