package cmd

import (
	"github.com/Saser/pdp/adventofcode/aoctool/cmd/fetch"
	"github.com/Saser/pdp/adventofcode/aoctool/cmd/generate"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoctool",
	Short: "A tool for working with Advent of Code in this repository.",
	Long:  `Aoctool is a tool for working with Advent of Code (https://adventofcode) within the context of this repository. It can be used to, for example, download stuff from the Advent of Code website, or generate boilerplate code for the solutions.`,
}

func init() {
	rootCmd.AddCommand(fetch.Cmd())
	rootCmd.AddCommand(generate.Cmd())
}

func RootCmd() *cobra.Command {
	return rootCmd
}

func Execute() error {
	return rootCmd.Execute()
}
