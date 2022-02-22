package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/holedaemon/hubris/internal/discord/api"
	"github.com/holedaemon/hubris/internal/discord/types"
	"github.com/rodaine/table"
)

var typeString = map[types.ApplicationCommandType]string{
	types.ApplicationCommandTypeChatInput: "CHAT_INPUT",
	types.ApplicationCommandTypeMessage:   "MESSAGE",
	types.ApplicationCommandTypeUser:      "USER",
}

type ManagerAddCmd struct {
	GuildID   string `help:"GuildID to apply commands to. Global if blank." short:"g"`
	Directory string `help:"Path of JSON files to use for upload." required:"" type:"existingdir" arg:""`
}

func (c *ManagerAddCmd) Run(ctx context.Context, g *Global) error {
	dir, err := os.ReadDir(c.Directory)
	if err != nil {
		return err
	}

	files := make([]string, 0)

	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		if !strings.HasSuffix(strings.ToLower(file.Name()), ".json") {
			continue
		}

		files = append(files, filepath.Join(c.Directory, file.Name()))
	}

	ap := make([]*types.ApplicationCommand, 0, len(files))

	for _, f := range files {
		dat, err := os.Open(f)
		if err != nil {
			return fmt.Errorf("reading JSON file: %w", err)
		}

		var apc *types.ApplicationCommand
		if err := json.NewDecoder(dat).Decode(&apc); err != nil {
			return err
		}

		ap = append(ap, apc)
	}

	cli, err := api.New(g.Token)
	if err != nil {
		return err
	}

	app := cli.Application(g.ApplicationID)

	if c.GuildID == "" {
		ap, err = app.BulkOverwriteGlobalApplicationCommands(ctx, ap)
	} else {
		ap, err = app.BulkOverwriteGuildApplicationCommands(ctx, c.GuildID, ap)
	}

	if err != nil {
		return err
	}

	tbl := table.New("Name", "ID", "Type", "Description")

	for _, a := range ap {
		tbl.AddRow(a.Name, a.ID, typeString[a.Type], a.Description)
	}

	tbl.Print()
	return nil
}

type ManagerListCmd struct {
	GuildID string `help:"GuildID to pull commands from." short:"g"`
}

func (c *ManagerListCmd) Run(ctx context.Context, g *Global) error {
	var (
		ap  []*types.ApplicationCommand
		err error
	)

	cli, err := api.New(g.Token)
	if err != nil {
		return err
	}

	app := cli.Application(g.ApplicationID)

	if c.GuildID == "" {
		ap, err = app.GetGlobalApplicationCommands(ctx)
	} else {
		ap, err = app.GetGuildApplicationCommands(ctx, c.GuildID)
	}

	if err != nil {
		return err
	}

	if len(ap) == 0 {
		fmt.Println("No commands.")
		return nil
	}

	tbl := table.New("Name", "ID", "Type", "Description")

	for _, a := range ap {
		tbl.AddRow(a.Name, a.ID, typeString[a.Type], a.Description)
	}

	tbl.Print()
	return nil
}

type ManagerRemoveCmd struct {
	GuildID string `help:"GuildID to remove command from." short:"g"`
	ID      string `help:"ID of command to remove." arg:""`
}

func (c *ManagerRemoveCmd) Run(ctx context.Context, g *Global) error {
	cli, err := api.New(g.Token)
	if err != nil {
		return err
	}

	app := cli.Application(g.ApplicationID)

	if c.GuildID == "" {
		return app.DeleteGlobalApplicationCommand(ctx, c.ID)
	} else {
		return app.DeleteGuildApplicationCommand(ctx, c.GuildID, c.ID)
	}
}
