package day22

import (
	"testing"

	"github.com/Saser/pdp/adventofcode/go/testcase"
)

var inputFile = testcase.Runfile("adventofcode/data/year2015/day22/actual.in")

var (
	tcPart1 = testcase.NewFile("input", inputFile, "1269")
	tcPart2 = testcase.NewFile("input", inputFile, "1309")
)

func TestPart1(t *testing.T) {
	tcPart1.Test(t, Part1)
}

func BenchmarkPart1(b *testing.B) {
	tcPart1.Benchmark(b, Part1)
}

func TestPart2(t *testing.T) {
	tcPart2.Test(t, Part2)
}

func BenchmarkPart2(b *testing.B) {
	tcPart2.Benchmark(b, Part2)
}
