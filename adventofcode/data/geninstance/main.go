package main

import (
	"errors"
	"flag"
	"io"
	"log"
	"os"

	adventofcodepb "github.com/Saser/pdp/adventofcode/adventofcode_go_proto"
)

var (
	year       = flag.Int("year", 0, "The year the problem is in.")
	day        = flag.Int("day", 0, "The day of the year.")
	part       = flag.Int("part", 0, "The part of the day.")
	name       = flag.String("name", "", "The name of the problem instance.")
	inputFile  = flag.String("input_file", "", "The path to a file containing the input of the problem instance.")
	answerFile = flag.String("answer_file", "", "The path to a file containing the answer of the problem instance.")
	outputFile = flag.String("output_file", "", "The path to the file where the instance should be written as a textproto.")
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
	f, err := os.Create(*outputFile)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("error closing output file: %v", err)
		}
	}()
	if err := writeProblemInstance(instance, f); err != nil {
		return err
	}
	return nil
}

func readFile(path string) (string, error) {
	return "", errors.New("function readFile not implemented")
}

func validateProblemID(id *adventofcodepb.ProblemID) error {
	return errors.New("function validateProblemID not implemented")
}

func validateProblemInstance(instance *adventofcodepb.ProblemInstance) error {
	return errors.New("function validateProblemInstance not implemented")
}

func writeProblemInstance(instance *adventofcodepb.ProblemInstance, w io.Writer) error {
	return errors.New("function writeProblemInstance not implemented")
}

func main() {
	if err := emain(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
