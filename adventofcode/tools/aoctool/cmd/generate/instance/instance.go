package instance

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"

	adventofcodepb "github.com/Saser/pdp/adventofcode/adventofcode_go_proto"
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
	outFile    string
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
	cmd.Flags().StringVar(&outFile, "out_file", "", "Path to a file where the Instance message should be written.")
	cmd.MarkFlagRequired("out_file")
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
	inputBytes, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}
	input := string(inputBytes)
	answerBytes, err := os.ReadFile(answerFile)
	if err != nil {
		return err
	}
	answer := string(answerBytes)
	if format != "prototext" && format != "binary" {
		return fmt.Errorf(`--format=%q is not a valid format; must be either "prototext" or "binary"`, format)
	}

	instance := &adventofcodepb.Instance{
		Problem: &adventofcodepb.Problem{
			Year: int32(year),
			Day:  int32(day),
			Part: int32(part),
		},
		Input:  input,
		Answer: answer,
	}
	var out []byte
	switch format {
	case "prototext":
		out, err = prototext.MarshalOptions{Multiline: true}.Marshal(instance)
	case "binary":
		out, err = proto.Marshal(instance)
	}
	if err != nil {
		return err
	}
	if err := os.WriteFile(outFile, out, fs.FileMode(0644)); err != nil {
		return err
	}
	return nil
}

func Cmd() *cobra.Command {
	return cmd
}
