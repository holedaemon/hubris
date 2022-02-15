package api

import (
	"github.com/holedaemon/hubris/internal/discord/api/resources"
	"github.com/holedaemon/hubris/internal/discord/api/resources/application"
	"github.com/holedaemon/hubris/internal/discord/api/resources/channel"
)

func (c *Client) Channel(id string) *channel.Resource {
	return channel.NewChannelResource(
		resources.NewRestClient(c.token, c.cli),
		id,
	)
}

func (c *Client) Application(id string) *application.Resource {
	return application.NewApplicationResource(
		resources.NewRestClient(c.token, c.cli),
		id,
	)
}
