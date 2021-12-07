package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"golang.org/x/net/publicsuffix"
)

type Client struct {
	client *http.Client
}

func New(session string) (*Client, error) {
	jar, err := createJar(session)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Jar: jar,
	}
	return &Client{
		client: client,
	}, nil
}

func (c *Client) GetPage(ctx context.Context, year int, day int) (string, error) {
	u := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
	body, err := get(ctx, c.client, u)
	if err != nil {
		return "", err
	}
	return body, nil
}

func (c *Client) GetInput(ctx context.Context, year int, day int) (string, error) {
	u := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	body, err := get(ctx, c.client, u)
	if err != nil {
		return "", err
	}
	return body, nil
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

func get(ctx context.Context, c *http.Client, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}
	res, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
