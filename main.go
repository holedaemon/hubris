package main

import (
	"context"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/alecthomas/kong"
	"github.com/holedaemon/hubris/internal/pkg/cli"
)

func main() {
	rand.Seed(time.Now().UnixNano())

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
