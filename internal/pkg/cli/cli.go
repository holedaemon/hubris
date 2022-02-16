package cli

import (
	"context"
	"fmt"

	"github.com/holedaemon/hubris/internal/pkg/version"
)

type App struct {
	Global

	Version VersionCmd `cmd:"" help:"Print the version." default:"1"`

	Run    RunCmd           `cmd:"" help:"Witness the hubris of man."`
	Add    ManagerAddCmd    `cmd:"" help:"Add a directory of JSON files as commands."`
	List   ManagerListCmd   `cmd:"" help:"List the application's commands."`
	Remove ManagerRemoveCmd `cmd:"" help:"Remove a command by its ID."`
}

type Global struct {
	Token         string `help:"The bot's OAuth2 token." required:"" short:"t" env:"HUBRIS_TOKEN"`
	ApplicationID string `help:"OAuth2 application snowflake." required:"" short:"a" env:"HUBRIS_APP_ID"`
}

type VersionCmd struct{}

func (c *VersionCmd) Run(ctx context.Context, g *Global) error {
	fmt.Printf("Running version %s of hubris\n", version.Version())
	return nil
}
