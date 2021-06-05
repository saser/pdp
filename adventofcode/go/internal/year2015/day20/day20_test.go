package day20

import (
	"testing"

	"github.com/Saser/adventofcode/internal/testcase"
)

const inputFile = "../testdata/20"

func TestPart1(t *testing.T) {
	for _, tc := range []testcase.TestCase{
		testcase.FromString("example1", "10", "1"),
		testcase.FromString("example2", "70", "4"),
		testcase.FromFile(t, inputFile, "665280"),
	} {
		testcase.Run(t, tc, Part1)
	}
}

func BenchmarkPart1(b *testing.B) {
	tc := testcase.FromFile(b, inputFile, "")
	testcase.Bench(b, tc, Part1)
}

func TestPart2(t *testing.T) {
	for _, tc := range []testcase.TestCase{
		testcase.FromString("example1", "10", "1"),
		testcase.FromString("example2", "70", "4"),
		testcase.FromFile(t, inputFile, ""),
	} {
		testcase.Run(t, tc, Part2)
	}
}

func BenchmarkPart2(b *testing.B) {
	tc := testcase.FromFile(b, inputFile, "")
	testcase.Bench(b, tc, Part2)
}
