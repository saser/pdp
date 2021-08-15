package testfile

import (
	"errors"

	"github.com/spf13/cobra"
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
	return errors.New("unimplemented")
}
