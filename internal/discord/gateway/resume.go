package gateway

import (
	"context"

	"nhooyr.io/websocket"
)

type resume struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Seq       int64  `json:"seq"`
}

func (c *Client) sendResume(ctx context.Context, ws *websocket.Conn) error {
	r := &resume{
		Token:     c.token,
		SessionID: c.sessionID,
		Seq:       c.sequence.Load(),
	}

	return write(ctx, ws, opResume, r)
}
