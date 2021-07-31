package answer

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/Saser/pdp/adventofcode/tools/aoctool/cmd/fetch/client"
	"github.com/spf13/cobra"
)

var (
	cmd = &cobra.Command{
		Use:   "answer",
		Short: "Download the answer to a problem from the Advent of Code website.",
		Long:  `Fetches the problem description page from the Advent of Code website, and uses an ugly hack based on a regex to try and parse out the answer from the HTML. The answer, if found, is printed to stdout.`,
		RunE:  runE,
	}

	year int
	day  int
	part int

	answerRE = regexp.MustCompile(`Your puzzle answer was \<code\>(.+?)\</code\>`)
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
	if year < 2015 || year > 2020 {
		return fmt.Errorf("--year=%d is outside range [2015, 2020]", year)
	}
	if day < 1 || day > 25 {
		return fmt.Errorf("--day=%d is outside range [1, 25]", day)
	}
	if part < 1 || part > 2 {
		return fmt.Errorf("--part=%d is not one of 1 or 2", part)
	}
	if day == 25 && part == 2 {
		return errors.New("there is no part 2 for day 25")
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
	page, err := c.GetPage(ctx, year, day)
	if err != nil {
		return err
	}
	answers := parseAnswers(page)
	if len(answers) < part {
		return fmt.Errorf("no answer found for year %d, day %d, part %d", year, day, part)
	}
	fmt.Print(answers[part-1])
	return nil
}

func parseAnswers(body string) []string {
	var answers []string
	for _, matches := range answerRE.FindAllStringSubmatch(body, -1) {
		answers = append(answers, matches[1])
	}
	return answers
}
