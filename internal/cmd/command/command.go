package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/holedaemon/hubris/internal/discord/api"
	"github.com/holedaemon/hubris/internal/discord/api/resources/application"
)

const dataDir = "data"

func main() {
	token := os.Getenv("TOKEN")
	id := os.Getenv("APPLICATION_ID")
	gid := os.Getenv("GUILD_ID")

	if token == "" {
		fmt.Printf("$TOKEN is blank.\n")
		return
	}

	if id == "" {
		fmt.Printf("$APPLICATION_ID is blank.\n")
		return
	}

	if gid == "" {
		fmt.Printf("$GUILD_ID is blank.\n")
		return
	}

	dir, err := os.ReadDir(dataDir)
	if err != nil {
		fmt.Printf("Error reading commands directory: %s\n", err.Error())
		return
	}

	params := make([]*application.CreateApplicationCommandParams, 0, len(dir))

	for _, f := range dir {
		if !strings.HasSuffix(f.Name(), ".json") {
			continue
		}

		fr, err := os.Open(filepath.Join(dataDir, f.Name()))
		if err != nil {
			fmt.Printf("Error reading JSON file: %s: %s\n", f.Name(), err.Error())
			return
		}

		var p *application.CreateApplicationCommandParams
		if err := json.NewDecoder(fr).Decode(&p); err != nil {
			fmt.Printf("Error decoding JSON file: %s\n", err.Error())
			return
		}

		params = append(params, p)
	}

	cli, err := api.New(token)
	if err != nil {
		fmt.Printf("Error creating API client: %s\n", err.Error())
		return
	}

	ar := cli.Application(id)

	ctx := context.Background()

	for _, p := range params {
		_, err := ar.CreateGuildApplicationCommand(ctx, gid, p)
		if err != nil {
			fmt.Printf("Error creating command: %s\n", err.Error())
			return
		}
	}
}
