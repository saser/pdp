package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"

	"google.golang.org/protobuf/encoding/prototext"

	adventofcodepb "github.com/Saser/pdp/adventofcode/adventofcode_go_proto"
)

var (
	year       = flag.Int("year", 0, "The year the problem is in.")
	day        = flag.Int("day", 0, "The day of the year.")
	part       = flag.Int("part", 0, "The part of the day.")
	name       = flag.String("name", "", "The name of the problem instance.")
	inputFile  = flag.String("input_file", "", "The path to a file containing the input of the problem instance.")
	answerFile = flag.String("answer_file", "", "The path to a file containing the answer of the problem instance.")
	outputFile = flag.String("output_file", "", "The path to the file where the instance should be written as a text proto.")
)

func emain() error {
	flag.Parse()
	input, err := readFile(*inputFile)
	if err != nil {
		return err
	}
	answer, err := readFile(*answerFile)
	if err != nil {
		return err
	}
	instance := &adventofcodepb.ProblemInstance{
		ProblemId: &adventofcodepb.ProblemID{
			Year: int32(*year),
			Day:  int32(*day),
			Part: adventofcodepb.ProblemID_Part(*part),
		},
		Name:   *name,
		Input:  input,
		Answer: answer,
	}
	if err := validateProblemInstance(instance); err != nil {
		return err
	}
	if err := writeProblemInstance(instance, *outputFile); err != nil {
		return err
	}
	return nil
}

func readFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("error closing file: %v", err)
		}
	}()
	bytes, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func validateProblemID(id *adventofcodepb.ProblemID) error {
	if year := id.GetYear(); year < 2015 || year > 2021 {
		return fmt.Errorf("problem ID has year outside range [2015, 2021]: %d", year)
	}
	if day := id.GetDay(); day < 1 || day > 25 {
		return fmt.Errorf("problem ID has day outside range [1, 25]: %d", day)
	}
	switch part := id.GetPart(); part {
	case adventofcodepb.ProblemID_PART_UNSPECIFIED:
		return errors.New("a part must be specified")
	case adventofcodepb.ProblemID_ONE:
		// this is fine
	case adventofcodepb.ProblemID_TWO:
		if id.GetDay() == 25 {
			return errors.New("problem ID specifies part 2 for day 25")
		}
	default:
		return fmt.Errorf("invalid part: %v", part)
	}
	return nil
}

func validateProblemInstance(instance *adventofcodepb.ProblemInstance) error {
	if err := validateProblemID(instance.GetProblemId()); err != nil {
		return err
	}
	if instance.GetName() == "" {
		return errors.New("problem instance has empty name")
	}
	if instance.GetInput() == "" {
		return errors.New("problem instance has empty input")
	}
	if instance.GetAnswer() == "" {
		return errors.New("problem instance has empty answer")
	}
	return nil
}

func writeProblemInstance(instance *adventofcodepb.ProblemInstance, path string) error {
	bytes, err := prototext.MarshalOptions{
		Multiline: true,
	}.Marshal(instance)
	if err != nil {
		return err
	}
	if err := os.WriteFile(path, bytes, fs.FileMode(0644)); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := emain(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
