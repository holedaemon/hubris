package heart

import (
	"context"
	"time"

	"github.com/holedaemon/hubris/internal/pkg/ws"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

// Beat is called when Pump ticks.
type Beat func(context.Context, *ws.Conn) error

// Pump beats the heart once every interval.
func Pump(ctx context.Context, interval time.Duration, c *ws.Conn, fn Beat) error {
	tkr := time.NewTicker(interval)
	ctxlog.Debug(ctx, "creating new heartbeat ticker", zap.Duration("every", interval))

	for {
		select {
		case <-ctx.Done():
			ctxlog.Debug(ctx, "heartbeater received cancellation signal")
			tkr.Stop()
			return ctx.Err()
		case <-tkr.C:
			ctxlog.Debug(ctx, "heartbeater has ticked")
			if err := fn(ctx, c); err != nil {
				return err
			}
		}
	}
}
