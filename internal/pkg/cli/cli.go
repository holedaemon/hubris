package cli

type App struct {
	Global

	Run    RunCmd           `cmd:"" help:"Witness the hubris of man."`
	Add    ManagerAddCmd    `cmd:"" help:"Add a directory of JSON files as commands."`
	List   ManagerListCmd   `cmd:"" help:"List the application's commands."`
	Remove ManagerRemoveCmd `cmd:"" help:"Remove a command by its ID."`
}

type Global struct {
	Token         string `help:"The bot's OAuth2 token." required:"" short:"t" env:"HUBRIS_TOKEN"`
	ApplicationID string `help:"OAuth2 application snowflake." required:"" short:"a" env:"HUBRIS_APP_ID"`
}
