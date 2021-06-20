package day24

import (
	"testing"

	"github.com/Saser/pdp/adventofcode/go/testcase"
)

var (
	exampleFile = testcase.Runfile("adventofcode/go/year2020/day24/testdata/example")
	inputFile   = testcase.Runfile("adventofcode/data/year2020/day24/actual.in")

	tcPart1 = testcase.NewFile("input", inputFile, "500")
	tcPart2 = testcase.NewFile("input", inputFile, "4280")
)

func TestPart1(t *testing.T) {
	for _, tc := range []testcase.TestCase{
		testcase.NewFile("example", exampleFile, "10"),
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
		testcase.NewFile("example", exampleFile, "2208"),
		tcPart2,
	} {
		tc.Test(t, Part2)
	}
}

func BenchmarkPart2(b *testing.B) {
	tcPart2.Benchmark(b, Part2)
}
