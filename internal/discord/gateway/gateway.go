// Package gateway implements a websocket client for the Discord gateway.
package gateway

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/holedaemon/hubris/internal/discord/api"
	"github.com/holedaemon/hubris/internal/pkg/exp"
	"github.com/holedaemon/hubris/internal/pkg/heart"
	"github.com/holedaemon/hubris/internal/pkg/ws"
	"go.uber.org/atomic"
	"golang.org/x/sync/errgroup"
	"nhooyr.io/websocket"
)

const (
	version = "9"
)

var defaultBackoff = exp.New(time.Second*1, time.Second*120, 1.6, 0.2)

type Client struct {
	token string

	url string

	connected *atomic.Bool

	lastAck   *atomic.Int64
	sequence  *atomic.Int64
	sessionID string

	backoff *exp.Backoff
}

func New(t string) (*Client, error) {
	if t == "" {
		return nil, errors.New("discord/gateway: token required")
	}

	return &Client{
		token:     t,
		lastAck:   atomic.NewInt64(0),
		sequence:  atomic.NewInt64(0),
		connected: atomic.NewBool(false),
		backoff:   defaultBackoff,
	}, nil
}

func (c *Client) connect(pc context.Context) (*errgroup.Group, error) {
	ctx := context.Background()

	header := make(http.Header)
	header.Set("Accept-Encoding", "zlib")

	ws, err := ws.Dial(ctx, c.url, header)
	if err != nil {
		return nil, err
	}

	defer func() {
		c.connected.Store(false)

		if err != nil {
			ws.Close(websocket.StatusInternalError, "internal error")
		} else {
			ws.Close(websocket.StatusNormalClosure, "normal")
		}
	}()

	var h *hello
	if err := read(ctx, ws, &h); err != nil {
		return nil, fmt.Errorf("%w: reading initial Hello", err)
	}

	if c.sequence.Load() == 0 && c.sessionID == "" {
		if err := c.sendIdentify(ctx, ws); err != nil {
			return nil, err
		}

		if err := c.getReady(ctx, ws); err != nil {
			return nil, err
		}
	} else {
		if err := c.sendResume(ctx, ws); err != nil {
			return nil, err
		}
	}

	grp, ctx := errgroup.WithContext(ctx)

	grp.Go(func() error {
		<-pc.Done()
		return ctx.Err()
	})

	grp.Go(func() error {
		return heart.Pump(ctx, time.Millisecond*time.Duration(h.HeartbeatInterval), ws, c.beat)
	})

	grp.Go(func() error {
		return c.read(ctx, ws)
	})

	return grp, nil
}

func (c *Client) Connect(pc context.Context) error {
	if c.connected.Load() {
		return fmt.Errorf("discord/gateway: already connected")
	}

	if c.url == "" {
		cli, err := api.New(c.token)
		if err != nil {
			return err
		}

		u, err := cli.GetGateway(pc)
		if err != nil {
			return err
		}

		q := make(url.Values)
		q.Set("v", version)
		q.Set("encoding", "json")

		c.url = u + q.Encode()
	}

	grp, err := c.connect(pc)
	if err != nil {
		return err
	}

	err = grp.Wait()
	if shouldReconnect(err) {
		for i := 0; true; i++ {
			dur := c.backoff.Attempt(i)

			time.Sleep(dur)

			ctx := context.Background()
			grp, err = c.connect(ctx)
			if err != nil {
				continue
			}

			err = grp.Wait()
			if !shouldReconnect(err) {
				return err
			}

			continue
		}
	}

	return err
}

func (c *Client) beat(ctx context.Context, ws *ws.Conn) error {
	pay := &payload{
		Op: opHeartbeat,
		D:  json.RawMessage(c.sequence.String()),
	}

	return ws.Write(ctx, pay)
}
