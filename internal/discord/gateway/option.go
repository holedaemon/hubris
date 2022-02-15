package gateway

type Option func(*Client)

func Debug() Option {
	return func(c *Client) {
		c.debug = true
	}
}
