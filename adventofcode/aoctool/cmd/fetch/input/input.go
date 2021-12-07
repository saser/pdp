package input

import (
	"fmt"

	"github.com/Saser/pdp/adventofcode/aoctool/cmd/fetch/client"
	"github.com/spf13/cobra"
)

const (
	minYear = 2015
	maxYear = 2021
)

var (
	cmd = &cobra.Command{
		Use:   "input",
		Short: "Download the input to a problem from the Advent of Code website.",
		Long:  "Authenticates using the given session to download the input from the Advent of Code website. The input is printed to stdout.",
		RunE:  runE,
	}

	year int
	day  int
)

func init() {
	cmd.Flags().IntVar(&year, "year", 0, fmt.Sprintf("The year in the range [%d, %d].", minYear, maxYear))
	cmd.MarkFlagRequired("year")
	cmd.Flags().IntVar(&day, "day", 0, "The day in the range [1, 25].")
	cmd.MarkFlagRequired("day")
}

func Cmd() *cobra.Command {
	return cmd
}

func runE(cmd *cobra.Command, args []string) error {
	if year < minYear || year > maxYear {
		return fmt.Errorf("--year=%d is outside range [%d, %d]", year, minYear, maxYear)
	}
	if day < 1 || day > 25 {
		return fmt.Errorf("--day=%d is outside range [1, 25]", day)
	}
	session, err := cmd.Flags().GetString("session")
	if err != nil {
		return err
	}

	c, err := client.New(session)
	if err != nil {
		return err
	}

	ctx := cmd.Context()
	input, err := c.GetInput(ctx, year, day)
	if err != nil {
		return err
	}
	fmt.Print(input)
	return nil
}
