package day13

import (
	"io/ioutil"
	"testing"

	"github.com/Saser/adventofcode/internal/testcase"
	"github.com/stretchr/testify/require"
)

const (
	exampleFile = "testdata/example"
	inputFile   = "../testdata/13"
)

var (
	tcPart1 = testcase.NewFile("input", inputFile, "618")
	tcPart2 = testcase.NewFile("input", inputFile, "601")
)

func Test_parsePreference(t *testing.T) {
	for _, tt := range []struct {
		name string
		s    string
		p    preference
	}{
		{
			name: "AliceBobGain54",
			s:    "Alice would gain 54 happiness units by sitting next to Bob.",
			p: preference{
				from:   "Alice",
				to:     "Bob",
				change: 54,
			},
		},
		{
			name: "AliceCarolLose79",
			s:    "Alice would gain 79 happiness units by sitting next to Carol.",
			p: preference{
				from:   "Alice",
				to:     "Carol",
				change: 79,
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			p, err := parsePreference(tt.s)
			require.NoError(t, err)
			require.Equal(t, tt.p, p)
		})
	}
}

func Test_parse(t *testing.T) {
	data, err := ioutil.ReadFile(exampleFile)
	require.NoError(t, err)
	m, err := parse(string(data))
	require.NoError(t, err)
	expected := map[string]map[string]int{
		"Alice": map[string]int{
			"Bob":   54,
			"Carol": -79,
			"David": -2,
		},
		"Bob": map[string]int{
			"Alice": 83,
			"Carol": -7,
			"David": -63,
		},
		"Carol": map[string]int{
			"Alice": -62,
			"Bob":   60,
			"David": 55,
		},
		"David": map[string]int{
			"Alice": 46,
			"Bob":   -7,
			"Carol": 41,
		},
	}
	require.Equal(t, expected, m)
}

func Test_score(t *testing.T) {
	data, err := ioutil.ReadFile(exampleFile)
	require.NoError(t, err)
	m, err := parse(string(data))
	require.NoError(t, err)
	names := []string{"Alice", "Bob", "Carol", "David"}
	require.Equal(t, 330, score(names, m, 1))
}

func TestPart1(t *testing.T) {
	for _, tc := range []testcase.TestCase{
		testcase.NewFile(exampleFile, exampleFile, "330"),
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
