package cli

import (
	"context"

	"github.com/holedaemon/hubris/internal/bot"
)

type RunCmd struct {
	Debug bool `help:"Run in debug mode?" default:"false" short:"d"`
}

func (c *RunCmd) Run(ctx context.Context, g *Global) error {
	b, err := bot.New(&bot.Options{
		Debug: c.Debug,
		Token: g.Token,
	})
	if err != nil {
		return err
	}

	return b.Connect(ctx)
}
