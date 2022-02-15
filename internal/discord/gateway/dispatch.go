package gateway

import (
	"context"
	"encoding/json"
)

func (c *Client) dispatch(ctx context.Context, ev string, data json.RawMessage) error {
	switch ev {
	case eventGuildCreate:
		return nil
	}

	return nil
}
