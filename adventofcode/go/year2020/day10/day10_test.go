package day10

import (
	"testing"

	"github.com/Saser/pdp/adventofcode/go/testcase"
)

var (
	example1File = testcase.Runfile("adventofcode/go/year2020/day10/testdata/example1")
	example2File = testcase.Runfile("adventofcode/go/year2020/day10/testdata/example2")
	inputFile    = testcase.Runfile("adventofcode/data/year2020/day10/actual.in")

	tcPart1 = testcase.NewFile("input", inputFile, "1920")
	tcPart2 = testcase.NewFile("input", inputFile, "1511207993344")
)

func TestPart1(t *testing.T) {
	for _, tc := range []testcase.TestCase{
		testcase.NewFile("example1", example1File, "35"),
		testcase.NewFile("example2", example2File, "220"),
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
		testcase.NewFile("example1", example1File, "8"),
		testcase.NewFile("example2", example2File, "19208"),
		tcPart2,
	} {
		tc.Test(t, Part2)
	}
}

func BenchmarkPart2(b *testing.B) {
	tcPart2.Benchmark(b, Part2)
}
