package input

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	cmd = &cobra.Command{
		Use:   "input",
		Short: "Download the input to a problem from the Advent of Code website.",
		RunE:  runE,
	}

	year int
	day  int
)

func init() {
	cmd.Flags().IntVar(&year, "year", 0, "The year in the range [2015, 2020].")
	cmd.MarkFlagRequired("year")
	cmd.Flags().IntVar(&day, "day", 0, "The day in the range [1, 25].")
	cmd.MarkFlagRequired("day")
}

func Cmd() *cobra.Command {
	return cmd
}

func runE(cmd *cobra.Command, args []string) error {
	session, err := cmd.Flags().GetString("session")
	if err != nil {
		return err
	}
	fmt.Printf("fetch input: year=%d, day=%d (session=%q)\n", year, day, session)
	return nil
}
