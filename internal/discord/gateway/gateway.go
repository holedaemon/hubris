// Package gateway implements a websocket client for the Discord gateway.
package gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/holedaemon/hubris/internal/discord/api"
	"github.com/holedaemon/hubris/internal/pkg/exp"
	"github.com/holedaemon/hubris/internal/pkg/heart"
	"github.com/holedaemon/hubris/internal/pkg/ws"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

const (
	version = "9"
)

var defaultBackoff = exp.New(time.Second*1, time.Second*120, 1.6, 0.2)

type Client struct {
	mu    sync.Mutex
	token string
	debug bool

	url string

	connected *atomic.Bool

	lastAck   *atomic.Int64
	sequence  *atomic.Int64
	sessionID string

	handlers map[string]handler

	backoff *exp.Backoff
	beater  *heart.Beater
	logger  *zap.Logger
}

func New(t string, opts ...Option) (*Client, error) {
	if t == "" {
		return nil, errors.New("discord/gateway: token required")
	}

	c := &Client{
		token:     t,
		lastAck:   atomic.NewInt64(0),
		sequence:  atomic.NewInt64(0),
		handlers:  make(map[string]handler),
		connected: atomic.NewBool(false),
		beater:    heart.NewBeater(),
		backoff:   defaultBackoff,
	}

	for _, o := range opts {
		o(c)
	}

	c.logger = ctxlog.New(c.debug)
	return c, nil
}

func (c *Client) connect(pc context.Context) error {
	ctx := context.Background()
	ctx = ctxlog.WithLogger(ctx, c.logger)

	header := make(http.Header)
	header.Set("Accept-Encoding", "zlib")

	ws, err := ws.Dial(ctx, c.url, header)
	if err != nil {
		return err
	}

	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(func() error {
		return c.beater.Pump(ctx, ws, c.beat)
	})

	grp.Go(func() error {
		return c.read(ctx, ws)
	})

	return grp.Wait()
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

		c.url = u + "?" + q.Encode()
	}

	err := c.connect(pc)
	if err != nil {
		return err
	}

	if shouldReconnect(err) {
		for i := 0; true; i++ {
			c.logger.Info("disconnected, attempting to reconnect", zap.Int("attempt", i))
			dur := c.backoff.Attempt(i)

			time.Sleep(dur)

			ctx := context.Background()
			err = c.connect(ctx)
			if err != nil {
				continue
			}

			if !shouldReconnect(err) {
				return err
			}

			continue
		}
	}

	return err
}

func (c *Client) beat(ctx context.Context, ws *ws.Conn) error {
	ctxlog.Debug(ctx, "sending heartbeat payload")
	return write(ctx, ws, opHeartbeat, c.sequence.Load())
}
