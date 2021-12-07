package testfile

import (
	"bytes"
	"go/format"
	"text/template"

	_ "embed" // for embedding Go template
)

var (
	//go:embed go.tmpl
	goTmplText string
	goTmpl     = template.Must(template.New("go.tmpl").Parse(goTmplText))

	goOut  string
	goData goTmplData
)

func init() {
	cmd.Flags().StringVar(&goOut, "go_out", "", "Path to file where Go code should be written.")
	cmd.Flags().StringVar(&goData.Package, "go_package", "", "Name of Go package to generate code in. Required if --go_out is non-empty.")
	cmd.Flags().StringVar(&goData.Part1, "go_part1", "", "Name of solution function for part 1. If empty then no code will be generated for testing part 1. At least one of --go_part1 and --go_part2 is required if --go_out is non-empty.")
	cmd.Flags().StringVar(&goData.Part2, "go_part2", "", "Name of solution function for part 2. If empty then no code will be generated for testing part 2. At least one of --go_part1 and --go_part2 is required if --go_out is non-empty.")
}

// goTmplData holds the data for the Go template.
type goTmplData struct {
	// Package is the name of the package the test should be
	// embedded in. Example: "day01".
	Package string
	// Part1Blobs and Part2Blobs are the Instance messages
	// marshalled in binary format for part 1 and part 2,
	// respectively.
	Part1Blobs, Part2Blobs [][]byte
	// Part1 and Part2 are the names of the solution functions for
	// part 1 and part 2, respectively. They should have the
	// signature func(input string) (string, error).
	Part1, Part2 string
}

func (d goTmplData) render() ([]byte, error) {
	var buf bytes.Buffer
	if err := goTmpl.Execute(&buf, d); err != nil {
		return nil, err
	}
	src, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, err
	}
	return src, nil
}
