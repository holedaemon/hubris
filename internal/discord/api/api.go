// Package api implements an API client for Discord.
package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

const (
	base    = "https://discord.com/api/v"
	version = "9"

	root = base + version
)

type Client struct {
	token string

	cli *http.Client
}

// New creates a new Client.
func New(t string, opts ...Option) (*Client, error) {
	if t == "" {
		return nil, fmt.Errorf("discord/api: token required")
	}

	if !strings.HasPrefix(t, "Bot ") {
		t = "Bot " + t
	}

	c := &Client{
		token: t,
	}

	for _, o := range opts {
		o(c)
	}

	if c.cli == nil {
		c.cli = http.DefaultClient
	}

	return c, nil
}

func (c *Client) GetGateway(ctx context.Context) (string, error) {
	rc := NewRestClient("", c.cli)

	var resp struct {
		URL string `json:"url"`
	}

	if err := rc.Get(ctx, "/gateway", &resp); err != nil {
		return "", err
	}

	return resp.URL, nil
}
