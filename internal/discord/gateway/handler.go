package gateway

import (
	"context"

	"github.com/holedaemon/hubris/internal/discord/types"
)

type readyHandler func(context.Context, *Ready)

func (h readyHandler) Handle(ctx context.Context, v interface{}) {
	h(ctx, v.(*Ready))
}

func (c *Client) OnReady(f func(context.Context, *Ready)) {
	c.registerEvent(eventReady, readyHandler(f))
}

type messageCreateHandler func(context.Context, *types.Message)

func (h messageCreateHandler) Handle(ctx context.Context, v interface{}) {
	h(ctx, v.(*types.Message))
}

func (c *Client) OnMessageCreate(f func(context.Context, *types.Message)) {
	c.registerEvent(eventMessageCreate, messageCreateHandler(f))
}

type interactionCreateHandler func(context.Context, *types.Interaction)

func (h interactionCreateHandler) Handle(ctx context.Context, v interface{}) {
	h(ctx, v.(*types.Interaction))
}

func (c *Client) OnInteractionCreate(f func(context.Context, *types.Interaction)) {
	c.registerEvent(eventInteractionCreate, interactionCreateHandler(f))
}

type guildCreateHandler func(context.Context, *types.Guild)

func (h guildCreateHandler) Handle(ctx context.Context, v interface{}) {
	h(ctx, v.(*types.Guild))
}

func (c *Client) OnGuildCreate(f func(context.Context, *types.Guild)) {
	c.registerEvent(eventGuildCreate, guildCreateHandler(f))
}
