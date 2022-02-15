package gateway

import (
	"context"
	"encoding/json"

	"github.com/holedaemon/hubris/internal/discord/types"
)

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
	}

	return nil
}
