package cli

import (
	"context"

	"github.com/holedaemon/hubris/internal/bot"
	"github.com/holedaemon/hubris/internal/database/dbx"
)

type RunCmd struct {
	Debug bool `help:"Run in debug mode?" default:"false" short:"d"`
}

func (c *RunCmd) Run(ctx context.Context, g *Global) error {
	db, err := dbx.Open(ctx, g.DSN)
	if err != nil {
		return err
	}

	b, err := bot.New(&bot.Options{
		Debug: c.Debug,
		Token: g.Token,
		DB:    db,
	})
	if err != nil {
		return err
	}

	return b.Connect(ctx)
}
