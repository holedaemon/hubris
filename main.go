package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/alecthomas/kong"
	"github.com/holedaemon/hubris/internal/pkg/cli"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill)
	defer cancel()

	app := cli.App{
		Global: cli.Global{},
	}

	ktx := kong.Parse(&app)

	ktx.Bind(&app.Global)
	ktx.BindTo(ctx, (*context.Context)(nil))

	err := ktx.Run(ctx, cli.Global{})
	ktx.FatalIfErrorf(err)
}
