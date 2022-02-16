package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/holedaemon/hubris/internal/bot"
)

func logErr(msg string, args ...interface{}) {
	if !strings.HasSuffix(msg, "\n") {
		msg += "\n"
	}

	fmt.Fprintf(os.Stderr, msg, args...)
}

func main() {
	token := os.Getenv("TOKEN")
	if token == "" {
		logErr("$TOKEN is not set")
		return
	}

	b, err := bot.New(&bot.Options{
		Debug: true,
		Token: token,
	})
	if err != nil {
		logErr("Error creating bot: %s", err)
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill)
	defer cancel()

	err = b.Connect(ctx)
	if err != nil {
		logErr("Error during Discord connection: %s", err)
		return
	}
}
