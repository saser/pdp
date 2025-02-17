package day14

import (
	"testing"

	"github.com/Saser/pdp/adventofcode/go/testcase"
)

var (
	part1ExampleFile = testcase.Runfile("adventofcode/go/year2020/day14/testdata/p1example")
	part2ExampleFile = testcase.Runfile("adventofcode/go/year2020/day14/testdata/p2example")
	inputFile        = testcase.Runfile("adventofcode/data/year2020/day14/actual.in")

	tcPart1 = testcase.NewFile("input", inputFile, "11884151942312")
	tcPart2 = testcase.NewFile("input", inputFile, "2625449018811")
)

func Test_mask_ApplyTo(t *testing.T) {
	maskStr := "mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"
	for _, tt := range []struct {
		n    int64
		want int64
	}{
		{n: 11, want: 73},
		{n: 101, want: 101},
		{n: 0, want: 64},
	} {
		if got := parseMask(maskStr).ApplyTo(tt.n); got != tt.want {
			t.Errorf("parseMask(%q).ApplyTo(%v) = %v; want %v", maskStr, tt.n, got, tt.want)
		}
	}
}

func TestPart1(t *testing.T) {
	for _, tc := range []testcase.TestCase{
		testcase.NewFile("example", part1ExampleFile, "165"),
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
		testcase.NewFile("example", part2ExampleFile, "208"),
		tcPart2,
	} {
		tc.Test(t, Part2)
	}
}

func BenchmarkPart2(b *testing.B) {
	tcPart2.Benchmark(b, Part2)
}
