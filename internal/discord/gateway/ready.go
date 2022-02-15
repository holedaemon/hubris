package gateway

import (
	"context"

	"github.com/holedaemon/hubris/internal/discord/types"
	"github.com/holedaemon/hubris/internal/pkg/ws"
)

type Ready struct {
	V         int                       `json:"v"`
	User      *types.User               `json:"user"`
	Guilds    []*types.UnavailableGuild `json:"guilds"`
	SessionID string                    `json:"session_id,omitempty"`
	Shard     [2]int                    `json:"shard,omitempty"`
}

func (c *Client) getReady(ctx context.Context, ws *ws.Conn) error {
	var r *Ready
	if err := read(ctx, ws, &r); err != nil {
		return err
	}

	c.sessionID = r.SessionID
	c.connected.Store(true)
	return nil
}
