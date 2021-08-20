package instance

import (
	"errors"
	"fmt"
	"io"
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
	name       string
	input      string
	inputFile  string
	answer     string
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

	cmd.Flags().StringVar(&name, "name", "", "The name.")
	cmd.MarkFlagRequired("name")

	cmd.Flags().StringVar(&input, "input", "", "The puzzle input. This flag is suitable for short inputs, such as examples. Exactly one of --input or --input_file must be specified.")
	cmd.Flags().StringVar(&inputFile, "input_file", "", "Path to a file to read the puzzle input from. Exactly one of --input or --input_file must be specified.")

	cmd.Flags().StringVar(&answer, "answer", "", "The puzzle answer. This flag is suitable for short answers, such as examples. Exactly one of --answer or --answer_file must be specified.")
	cmd.Flags().StringVar(&answerFile, "answer_file", "", "Path to a file to read the puzzle answer from. Exactly one of --answer or --answer_file must be specified.")

	cmd.Flags().StringVar(&format, "format", "", `The format to write the protobuf message in. Must be either "prototext" or "binary".`)
	cmd.MarkFlagRequired("format")

	cmd.Flags().StringVar(&outFile, "out_file", "", `Path to a file where the Instance message should be written. If set to "-", the output will be written to stdout.`)
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
	if (input == "" && inputFile == "") || (input != "" && inputFile != "") {
		return fmt.Errorf("--input=%q and --input_file=%q is invalid; exactly one of them must be non-empty", input, inputFile)
	}
	var input2 string
	switch {
	case input != "":
		input2 = input
	case inputFile != "":
		inputBytes, err := os.ReadFile(inputFile)
		if err != nil {
			return err
		}
		input2 = string(inputBytes)
	}
	if (answer == "" && answerFile == "") || (answer != "" && answerFile != "") {
		return fmt.Errorf("--answer=%q and --answer_file=%q is invalid; exactly one of them must be non-empty", answer, answerFile)
	}
	var answer2 string
	switch {
	case answer != "":
		answer2 = answer
	case answerFile != "":
		answerBytes, err := os.ReadFile(answerFile)
		if err != nil {
			return err
		}
		answer2 = string(answerBytes)
	}
	if format != "prototext" && format != "binary" {
		return fmt.Errorf(`--format=%q is not a valid format; must be either "prototext" or "binary"`, format)
	}

	instance := &adventofcodepb.Instance{
		Problem: &adventofcodepb.Problem{
			Year: int32(year),
			Day:  int32(day),
			Part: int32(part),
		},
		Name:   name,
		Input:  input2,
		Answer: answer2,
	}
	var (
		out []byte
		err error
	)
	switch format {
	case "prototext":
		out, err = prototext.MarshalOptions{Multiline: true}.Marshal(instance)
	case "binary":
		out, err = proto.Marshal(instance)
	}
	if err != nil {
		return err
	}
	w := os.Stdout
	if outFile != "-" {
		f, err := os.Create(outFile)
		if err != nil {
			return nil
		}
		defer f.Close()
		w = f
	}
	if err := writeAll(w, out); err != nil {
		return err
	}
	return nil
}

func Cmd() *cobra.Command {
	return cmd
}

func writeAll(w io.Writer, data []byte) error {
	written := 0
	for written < len(data) {
		n, err := w.Write(data[written:])
		if err != nil {
			return err
		}
		written += n
	}
	return nil
}
