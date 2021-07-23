package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "aoctool",
		Short: "A tool for working with Advent of Code solutions",
		Long:  "Aoctool is a tool for interacting with the Advent of Code website (https://adventofcode.com) and with the solutions stored in this repository.",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
