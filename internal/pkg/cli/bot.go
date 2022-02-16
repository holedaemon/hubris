package cli

import (
	"context"

	"github.com/holedaemon/hubris/internal/bot"
)

type BotCmd struct {
	Token string `help:"The bot's OAuth2 token." required:"" short:"t" env:"HUBRIS_TOKEN"`
	Debug bool   `help:"Run in debug mode?" default:"false" short:"d"`
}

func (c *BotCmd) Run(ctx context.Context) error {
	b, err := bot.New(&bot.Options{
		Debug: c.Debug,
		Token: c.Token,
	})
	if err != nil {
		return err
	}

	return b.Connect(ctx)
}
