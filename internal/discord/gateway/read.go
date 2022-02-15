package gateway

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/holedaemon/hubris/internal/pkg/ws"
)

var errReconnect = errors.New("reconnect")

func (c *Client) read(ctx context.Context, ws *ws.Conn) error {
	var p *payload
	if err := ws.Read(ctx, &p); err != nil {
		return err
	}

	if p.S.Valid {
		c.sequence.Store(int64(p.S.Value))
	}

	switch p.Op {
	case opDispatch:
		return c.dispatch(ctx, p.T.Value, p.D)
	case opHeartbeat:
		return c.beat(ctx, ws)
	case opReconnect:
		return errReconnect
	case opInvalidSession:
		var resumable bool
		if err := read(ctx, ws, &resumable); err != nil {
			return err
		}

		if resumable {
			if err := c.sendResume(ctx, ws); err != nil {
				return err
			}
		} else {
			time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)

			c.sequence.Store(0)
			c.sessionID = ""
			if err := c.sendIdentify(ctx, ws); err != nil {
				return err
			}
		}
	case opHello:
		// NO-OP
	case opHeartbeatAck:
		c.lastAck.Store(time.Now().UnixNano())
	}

	return nil
}
