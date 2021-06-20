package day08

import (
	"testing"

	"github.com/Saser/pdp/adventofcode/go/testcase"
)

var (
	exampleFile = testcase.Runfile("adventofcode/go/year2020/day08/testdata/example")
	inputFile   = testcase.Runfile("adventofcode/data/year2020/day08/actual.in")

	tcPart1 = testcase.NewFile("input", inputFile, "1317")
	tcPart2 = testcase.NewFile("input", inputFile, "1033")
)

func TestPart1(t *testing.T) {
	for _, tc := range []testcase.TestCase{
		testcase.NewFile("example", exampleFile, "5"),
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
		testcase.NewFile("example", exampleFile, "8"),
		tcPart2,
	} {
		tc.Test(t, Part2)
	}
}

func BenchmarkPart2(b *testing.B) {
	tcPart2.Benchmark(b, Part2)
}
