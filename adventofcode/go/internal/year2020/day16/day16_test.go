package day16

import (
	"testing"

	"github.com/Saser/pdp/adventofcode/go/internal/testcase"
)

var (
	exampleFile = testcase.Runfile("adventofcode/go/internal/year2020/day16/testdata/example")
	inputFile   = testcase.Runfile("adventofcode/inputs/2020/16")

	tcPart1 = testcase.NewFile("input", inputFile, "27898")
	tcPart2 = testcase.NewFile("input", inputFile, "2766491048287")
)

func TestPart1(t *testing.T) {
	for _, tc := range []testcase.TestCase{
		testcase.NewFile("example", exampleFile, "71"),
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
