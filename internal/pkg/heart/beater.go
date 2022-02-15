package heart

import (
	"context"
	"time"

	"github.com/holedaemon/hubris/internal/pkg/ws"
)

// Beat is called when Pump ticks.
type Beat func(context.Context, *ws.Conn) error

// Pump beats the heart once every interval.
func Pump(ctx context.Context, interval time.Duration, c *ws.Conn, fn Beat) error {
	tkr := time.NewTicker(interval)

	for {
		select {
		case <-ctx.Done():
			tkr.Stop()
			return ctx.Err()
		case <-tkr.C:
			if err := fn(ctx, c); err != nil {
				return err
			}
		}
	}
}
