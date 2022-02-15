package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/holedaemon/hubris/internal/discord/gateway"
)

func main() {
	token := os.Getenv("TOKEN")
	if token == "" {
		fmt.Printf("$TOKEN is blank\n")
		return
	}

	c, err := gateway.New(token)
	if err != nil {
		fmt.Printf("Error creating gateway client: %s\n", err.Error())
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill)
	defer cancel()

	err = c.Connect(ctx)
	if err != nil {
		fmt.Printf("Error connecting to Discord: %s\n", err.Error())
		return
	}

	fmt.Printf("Closed without error")
}
