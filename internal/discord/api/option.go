package api

import "net/http"

// Option configures a Client
type Option func(*Client)

// HTTPClient sets a Client's HTTP client.
func HTTPClient(cli *http.Client) Option {
	return func(c *Client) {
		c.cli = cli
	}
}
