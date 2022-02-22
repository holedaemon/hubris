package bot

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/holedaemon/hubris/internal/discord/api"
	"github.com/holedaemon/hubris/internal/discord/gateway"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

// Bot is a Discord API/Gateway client.
type Bot struct {
	appCommands map[string]ApplicationCommand

	logger  *zap.Logger
	api     *api.Client
	gateway *gateway.Client
	db      *sql.DB
}

// Options are used to configure a Bot client.
type Options struct {
	Debug bool
	Token string

	DB *sql.DB
}

// New creates a new Bot from opts.
func New(opts *Options) (*Bot, error) {
	if opts.Token == "" {
		return nil, fmt.Errorf("bot: token is blank")
	}

	b := &Bot{
		appCommands: make(map[string]ApplicationCommand),
		logger:      ctxlog.New(opts.Debug),
		db:          opts.DB,
	}

	a, err := api.New(opts.Token)
	if err != nil {
		return nil, err
	}

	b.api = a

	g, err := gateway.New(opts.Token, gateway.Debug())
	if err != nil {
		return nil, err
	}

	g.OnReady(b.handleReady)
	g.OnGuildCreate(b.handleGuildCreate)
	g.OnMessageCreate(b.handleMessageCreate)
	g.OnInteractionCreate(b.handleInteractionCreate)

	b.gateway = g

	return b, nil
}

// Connect has the internal gateway client connect to Discord.
func (b *Bot) Connect(ctx context.Context) error {
	defer func() {
		b.logger.Info("Connect() has finished")
	}()
	return b.gateway.Connect(ctx)
}
