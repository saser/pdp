package day23

import (
	"testing"

	"github.com/Saser/pdp/adventofcode/go/testcase"
)

var inputFile = testcase.Runfile("adventofcode/data/year2020/day23/actual.in")

var (
	tcPart1 = testcase.NewFile("input", inputFile, "82635947")
	tcPart2 = testcase.NewFile("input", inputFile, "157047826689")
)

func TestPart1(t *testing.T) {
	for _, tc := range []testcase.TestCase{
		testcase.New("example", "389125467", "67384529"),
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
		testcase.New("example", "389125467", "149245887792"),
		tcPart2,
	} {
		tc.Test(t, Part2)
	}
}

func BenchmarkPart2(b *testing.B) {
	tcPart2.Benchmark(b, Part2)
}
