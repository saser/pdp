package day25

import (
	"testing"

	"github.com/Saser/pdp/adventofcode/go/testcase"
)

var (
	exampleFile = testcase.Runfile("adventofcode/go/year2020/day25/testdata/example")
	inputFile   = testcase.Runfile("adventofcode/data/year2020/day25/actual.in")

	tcPart1 = testcase.NewFile("input", inputFile, "17980581")
)

func TestPart1(t *testing.T) {
	for _, tc := range []testcase.TestCase{
		testcase.NewFile("example", exampleFile, "14897079"),
		tcPart1,
	} {
		tc.Test(t, Part1)
	}
}

func BenchmarkPart1(b *testing.B) {
	tcPart1.Benchmark(b, Part1)
}
