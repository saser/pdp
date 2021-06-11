package day17

import (
	"testing"

	"github.com/Saser/pdp/adventofcode/go/internal/testcase"
)

var (
	exampleFile = testcase.Runfile("adventofcode/go/internal/year2020/day17/testdata/example")
	inputFile   = testcase.Runfile("adventofcode/inputs/2020/17")

	tcPart1 = testcase.NewFile("input", inputFile, "295")
	tcPart2 = testcase.NewFile("input", inputFile, "1972")
)

func TestPart1(t *testing.T) {
	for _, tc := range []testcase.TestCase{
		testcase.NewFile("example", exampleFile, "112"),
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
		testcase.NewFile("example", exampleFile, "848"),
		tcPart2,
	} {
		tc.Test(t, Part2)
	}
}

func BenchmarkPart2(b *testing.B) {
	tcPart2.Benchmark(b, Part2)
}
