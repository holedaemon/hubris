package gateway

import (
	"context"

	"github.com/holedaemon/hubris/internal/pkg/ws"
)

type resume struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Seq       int64  `json:"seq"`
}

func (c *Client) sendResume(ctx context.Context, ws *ws.Conn) error {
	r := &resume{
		Token:     c.token,
		SessionID: c.sessionID,
		Seq:       c.sequence.Load(),
	}

	return ws.Write(ctx, r)
}
