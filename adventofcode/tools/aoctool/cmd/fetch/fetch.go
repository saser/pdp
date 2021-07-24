package fetch

import (
	"github.com/Saser/pdp/adventofcode/tools/aoctool/cmd/fetch/answer"
	"github.com/Saser/pdp/adventofcode/tools/aoctool/cmd/fetch/input"
	"github.com/spf13/cobra"
)

var (
	cmd = &cobra.Command{
		Use:   "fetch",
		Short: "Fetch stuff from the Advent of Code website (https://adventofcode.com).",
	}

	Session = cmd.PersistentFlags().String("session", "", "The value of the session cookie.")
)

func init() {
	cmd.MarkPersistentFlagRequired("session")

	cmd.AddCommand(answer.Cmd())
	cmd.AddCommand(input.Cmd())
}

func Cmd() *cobra.Command {
	return cmd
}
