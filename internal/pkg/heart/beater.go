package heart

import (
	"context"
	"time"

	"github.com/zikaeroh/ctxlog"
	"nhooyr.io/websocket"
)

// Beat is called when Pump ticks.
type Beat func(context.Context, *websocket.Conn) error

type Beater struct {
	ch chan time.Duration
}

func NewBeater() *Beater {
	return &Beater{
		ch: make(chan time.Duration),
	}
}

func (b *Beater) Notify(nd time.Duration) {
	b.ch <- nd
}

func (b *Beater) Pump(ctx context.Context, ws *websocket.Conn, fn Beat) error {
	var tkr *time.Ticker

	select {
	case d := <-b.ch:
		tkr = time.NewTicker(d * time.Millisecond)
	case <-ctx.Done():
		if tkr != nil {
			tkr.Stop()
		}

		return ctx.Err()
	}

	for {
		select {
		case <-tkr.C:
			ctxlog.Debug(ctx, "heartbeater has ticked")
			if err := fn(ctx, ws); err != nil {
				return err
			}
		case <-ctx.Done():
			tkr.Stop()
			return ctx.Err()
		}
	}
}
