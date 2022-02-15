package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/holedaemon/hubris/internal/discord/gateway"
	"github.com/holedaemon/hubris/internal/discord/types"
)

func main() {
	token := os.Getenv("TOKEN")
	if token == "" {
		fmt.Printf("$TOKEN is blank\n")
		return
	}

	c, err := gateway.New(token, gateway.Debug())
	if err != nil {
		fmt.Printf("Error creating gateway client: %s\n", err.Error())
		return
	}

	c.OnReady(handleReady)
	c.OnMessageCreate(handleMessage)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill)
	defer cancel()

	err = c.Connect(ctx)
	if err != nil {
		fmt.Printf("Error connecting to Discord: %s\n", err.Error())
		return
	}

	fmt.Printf("Closed without error")
}

func handleReady(ctx context.Context, r *gateway.Ready) {
	fmt.Println(r)
}

func handleMessage(ctx context.Context, m *types.Message) {
	fmt.Println(m.Content)
}
