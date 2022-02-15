package gateway

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/holedaemon/hubris/internal/pkg/ws"
)

var errReconnect = errors.New("reconnect")

func (c *Client) read(ctx context.Context, ws *ws.Conn) error {
	for {
		reader, err := ws.Read(ctx)
		if err != nil {
			return fmt.Errorf("%w: reading from conn", err)
		}

		var p *payload
		if err := json.NewDecoder(reader).Decode(&p); err != nil {
			return fmt.Errorf("%w: decoding JSON", err)
		}

		if p.S.Valid {
			c.sequence.Store(int64(p.S.Value))
		}

		switch p.Op {
		case opDispatch:
			if err := c.dispatch(ctx, p.T.Value, p.D); err != nil {
				return err
			}
		case opHeartbeat:
			if err := c.beat(ctx, ws); err != nil {
				return err
			}
		case opReconnect:
			return errReconnect
		case opInvalidSession:
			var resumable bool

			if err := json.Unmarshal(p.D, &resumable); err != nil {
				return err
			}

			if resumable {
				dur := c.backoff.Attempt(1)

				time.Sleep(dur)

				if err := c.sendResume(ctx, ws); err != nil {
					return err
				}
			} else {
				if err := c.sendIdentify(ctx, ws); err != nil {
					return err
				}
			}
		case opHello:
			var h *hello

			if err := json.Unmarshal(p.D, &h); err != nil {
				return err
			}

			c.beater.Notify(time.Duration(h.HeartbeatInterval))

			if c.sequence.Load() == 0 && c.sessionID == "" {
				if err := c.sendIdentify(ctx, ws); err != nil {
					return fmt.Errorf("%w: sending identify payload", err)
				}
			} else {
				if err := c.sendResume(ctx, ws); err != nil {
					return fmt.Errorf("%w: sending resume payload", err)
				}
			}
		case opHeartbeatAck:
			c.lastAck.Store(time.Now().UnixNano())
		}
	}
}
