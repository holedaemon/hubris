package gateway

import (
	"context"
	"encoding/json"
	"time"

	"github.com/holedaemon/hubris/internal/discord/types"
)

type handler interface {
	Handle(context.Context, interface{})
}

func (c *Client) handle(ctx context.Context, ev string, v interface{}) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	h, ok := c.handlers[ev]
	if ok {
		go h.Handle(ctx, v)
	}
}

func (c *Client) registerEvent(ev string, h handler) {
	if h == nil {
		panic("discord/gateway: registering nil handler")
	}

	c.mu.Lock()
	c.handlers[ev] = h
	c.mu.Unlock()
}

func (c *Client) dispatch(ctx context.Context, ev string, data json.RawMessage) error {
	var v interface{}

	switch ev {
	case eventReady:
		var r *Ready

		if err := json.Unmarshal(data, &r); err != nil {
			return err
		}

		v = r
	case eventGuildCreate:
		var g *types.Guild

		if err := json.Unmarshal(data, &g); err != nil {
			return err
		}

		v = g
	case eventInteractionCreate:
		var i *types.Interaction

		if err := json.Unmarshal(data, &i); err != nil {
			return err
		}

		v = i
	case eventMessageCreate:
		var m *types.Message

		if err := json.Unmarshal(data, &m); err != nil {
			return err
		}

		v = m
	}

	c.handle(ctx, ev, v)
	return nil
}
