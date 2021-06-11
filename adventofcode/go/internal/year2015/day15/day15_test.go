package day15

import (
	"testing"

	"github.com/Saser/pdp/adventofcode/go/internal/testcase"
)

var (
	exampleFile = testcase.Runfile("adventofcode/go/internal/year2015/day15/testdata/example")
	inputFile   = testcase.Runfile("adventofcode/inputs/2015/15")

	tcPart1 = testcase.NewFile("input", inputFile, "13882464")
	tcPart2 = testcase.NewFile("input", inputFile, "11171160")
)

func TestPart1(t *testing.T) {
	for _, tc := range []testcase.TestCase{
		testcase.NewFile(exampleFile, exampleFile, "62842880"),
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
		testcase.NewFile(exampleFile, exampleFile, "57600000"),
		tcPart2,
	} {
		tc.Test(t, Part2)
	}
}

func BenchmarkPart2(b *testing.B) {
	tcPart2.Benchmark(b, Part2)
}
