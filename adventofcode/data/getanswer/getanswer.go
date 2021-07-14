package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"os/signal"
	"regexp"

	"golang.org/x/net/publicsuffix"
	"golang.org/x/sync/errgroup"
)

const (
	baseURL = "https://adventofcode.com"
)

var (
	session = flag.String("session", "", `The value to use for the "session" cookie.`)

	answerRE = regexp.MustCompile(`Your puzzle answer was \<code\>(.+?)\</code\>`)
)

func emain() error {
	flag.Parse()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	g, ctx := errgroup.WithContext(ctx)
	jar, err := createJar(*session)
	if err != nil {
		return err
	}
	client := &http.Client{
		Jar: jar,
	}
	for year := 2015; year <= 2020; year++ {
		for day := 1; day <= 25; day++ {
			year, day := year, day
			g.Go(func() error {
				body, err := getProblem(ctx, client, year, day)
				if err != nil {
					return err
				}
				for i, answer := range parseAnswers(body) {
					log.Printf("year%d/day%02d/part%d: %s", year, day, i+1, answer)
				}
				return nil
			})
		}
	}
	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func createJar(session string) (*cookiejar.Jar, error) {
	if session == "" {
		return nil, errors.New("empty session")
	}
	jar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if err != nil {
		return nil, err
	}
	u := &url.URL{
		Scheme: "https",
		Host:   "adventofcode.com",
	}
	cookie := &http.Cookie{
		Name:  "session",
		Value: session,
	}
	jar.SetCookies(u, []*http.Cookie{cookie})
	return jar, nil
}

func getProblem(ctx context.Context, c *http.Client, year, day int) (body string, err error) {
	u := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
	req, err := http.NewRequestWithContext(ctx, "GET", u, nil)
	if err != nil {
		return "", err
	}
	res, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		if closeErr := res.Body.Close(); err == nil && closeErr != nil {
			err = closeErr
		}
	}()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func parseAnswers(body string) []string {
	var answers []string
	for _, matches := range answerRE.FindAllStringSubmatch(body, -1) {
		answers = append(answers, matches[1])
	}
	return answers
}

func main() {
	if err := emain(); err != nil {
		log.Print(err)
	}
}

func problemRequest(ctx context.Context, year, day int) (*http.Request, error) {
	url := fmt.Sprintf("%s/%d/day/%d", baseURL, year, day)
	r, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	return r, nil
}
