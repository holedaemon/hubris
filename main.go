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

	var app cli.App
	ktx := kong.Parse(&app)

	ktx.BindTo(ctx, (*context.Context)(nil))
	err := ktx.Run(ctx)
	ktx.FatalIfErrorf(err)
}
