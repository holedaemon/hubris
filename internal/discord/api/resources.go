package api

import (
	"github.com/holedaemon/hubris/internal/discord/api/resources"
	"github.com/holedaemon/hubris/internal/discord/api/resources/channel"
)

func (c *Client) Channel(id string) *channel.Resource {
	return channel.NewChannelResource(
		resources.NewRestClient(c.token, c.cli),
		id,
	)
}
