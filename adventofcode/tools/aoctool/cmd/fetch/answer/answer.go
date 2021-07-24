package answer

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	cmd = &cobra.Command{
		Use:   "answer",
		Short: "Download the answer to a problem from the Advent of Code website.",
		RunE:  runE,
	}

	year int
	day  int
	part int
)

func init() {
	cmd.Flags().IntVar(&year, "year", 0, "The year in the range [2015, 2020].")
	cmd.MarkFlagRequired("year")
	cmd.Flags().IntVar(&day, "day", 0, "The day in the range [1, 25].")
	cmd.MarkFlagRequired("day")
	cmd.Flags().IntVar(&part, "part", 0, "The part, which must be either 1 or 2.")
	cmd.MarkFlagRequired("part")
}

func Cmd() *cobra.Command {
	return cmd
}

func runE(cmd *cobra.Command, args []string) error {
	session, err := cmd.Flags().GetString("session")
	if err != nil {
		return err
	}
	fmt.Printf("fetch answer: year=%d, day=%d, part=%d (session=%q)\n", year, day, part, session)
	return nil
}
