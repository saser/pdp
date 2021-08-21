package testfile

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"

	adventofcodepb "github.com/Saser/pdp/adventofcode/adventofcode_go_proto"
)

var (
	cmd = &cobra.Command{
		Use:   "testfile",
		Short: "Generate test code from adventofcode.Instance messages.",
		RunE:  runE,
	}
)

func Cmd() *cobra.Command {
	return cmd
}

func runE(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("at least one textproto file must be given as argument")
	}
	var part1Instances, part2Instances []*adventofcodepb.Instance
	for _, arg := range args {
		b, err := os.ReadFile(arg)
		if err != nil {
			return err
		}
		instance := &adventofcodepb.Instance{}
		if err := prototext.Unmarshal(b, instance); err != nil {
			return fmt.Errorf("read %q: %w", arg, err)
		}
		switch instance.GetProblem().GetPart() {
		case 1:
			part1Instances = append(part1Instances, instance)
		case 2:
			part2Instances = append(part2Instances, instance)
		}
	}

	if goOut != "" {
		if goData.Package == "" {
			return errors.New("--go_package is required if --go_out is non-empty")
		}
		if goData.Part1 == "" {
			return errors.New("--go_part1 is required if --go_out is non-empty")
		}
		if goData.Part2 == "" {
			return errors.New("--go_part2 is required if --go_out is non-empty")
		}
		for _, instance := range part1Instances {
			b, err := proto.Marshal(instance)
			if err != nil {
				return err
			}
			goData.Part1Blobs = append(goData.Part1Blobs, b)
		}
		for _, instance := range part2Instances {
			b, err := proto.Marshal(instance)
			if err != nil {
				return err
			}
			goData.Part2Blobs = append(goData.Part2Blobs, b)
		}
		src, err := goData.render()
		if err != nil {
			return fmt.Errorf("render Go code: %w", err)
		}
		if err := os.WriteFile(goOut, src, 0644); err != nil {
			return fmt.Errorf("render Go code: %w", err)
		}
	}
	return nil
}
