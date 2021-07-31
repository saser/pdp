package generate

import (
	"github.com/Saser/pdp/adventofcode/tools/aoctool/cmd/generate/instance"
	"github.com/spf13/cobra"
)

var (
	cmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate various kinds of helpers and boilerplate relating to Advent of Code.",
	}
)

func init() {
	cmd.AddCommand(instance.Cmd())
}

func Cmd() *cobra.Command {
	return cmd
}
