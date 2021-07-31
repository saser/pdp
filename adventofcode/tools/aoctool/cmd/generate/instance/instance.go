package instance

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	cmd = &cobra.Command{
		Use:   "instance",
		Short: "Generate a adventofcode.Instance message and write it to stdout.",
		RunE:  runE,
	}

	year       int
	day        int
	part       int
	inputFile  string
	answerFile string
	format     string
)

func init() {
	cmd.Flags().IntVar(&year, "year", 0, "The year in the range [2015, 2020].")
	cmd.MarkFlagRequired("year")
	cmd.Flags().IntVar(&day, "day", 0, "The day in the range [1, 25].")
	cmd.MarkFlagRequired("day")
	cmd.Flags().IntVar(&part, "part", 0, "The part, which must be either 1 or 2.")
	cmd.MarkFlagRequired("part")
	cmd.Flags().StringVar(&inputFile, "input_file", "", "Path to a file to take the puzzle input from.")
	cmd.MarkFlagRequired("input_file")
	cmd.Flags().StringVar(&answerFile, "answer_file", "", "Path to a file to take the puzzle answer from.")
	cmd.MarkFlagRequired("answer_file")
	cmd.Flags().StringVar(&format, "format", "", `The format to write the protobuf message in. Must be either "prototext" or "binary".`)
	cmd.MarkFlagRequired("format")
}

func runE(cmd *cobra.Command, args []string) error {
	if year < 2015 || year > 2020 {
		return fmt.Errorf("--year=%d is outside range [2015, 2020]", year)
	}
	if day < 1 || day > 25 {
		return fmt.Errorf("--day=%d is outside range [1, 25]", day)
	}
	if part < 1 || part > 2 {
		return fmt.Errorf("--part=%d is not one of 1 or 2", part)
	}
	if day == 25 && part == 2 {
		return errors.New("there is no part 2 for day 25")
	}
	if format != "prototext" && format != "binary" {
		return fmt.Errorf(`--format=%q is not a valid format; must be either "prototext" or "binary"`, format)
	}
	fmt.Printf("year=%d day=%d part=%d input_file=%q answer_file=%q format=%q\n", year, day, part, inputFile, answerFile, format)
	return nil
}

func Cmd() *cobra.Command {
	return cmd
}
