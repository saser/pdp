package generate

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	cmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate boilerplate code for solutions in various languages.",
		RunE:  runE,
	}

	year     int
	day      int
	language string
	basedir  string
)

func init() {
	cmd.Flags().IntVar(&year, "year", 0, "The year in the range [2015, 2020].")
	cmd.MarkFlagRequired("year")
	cmd.Flags().IntVar(&day, "day", 0, "The day in the range [1, 25].")
	cmd.MarkFlagRequired("day")
	cmd.Flags().StringVar(&language, "language", "", `Which language to generate code for. Must be one of "cpp", "go", "java", or "rust".`)
	cmd.MarkFlagRequired("language")
	cmd.Flags().StringVar(&basedir, "basedir", "", `The path to the top-level "adventofcode" directory.`)
	cmd.MarkFlagRequired("basedir")
}

func Cmd() *cobra.Command {
	return cmd
}

func runE(cmd *cobra.Command, args []string) error {
	fmt.Printf("generate: year=%d, day=%d, language=%q, basedir=%q\n", year, day, language, basedir)
	return nil
}
