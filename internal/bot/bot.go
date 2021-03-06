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

type Bot struct {
	logger *zap.Logger

	api     *api.Client
	gateway *gateway.Client

	db *sql.DB
}

type Options struct {
	Debug bool
	Token string

	DB *sql.DB
}

func New(opts *Options) (*Bot, error) {
	if opts.Token == "" {
		return nil, fmt.Errorf("bot: token is blank")
	}

	b := &Bot{
		logger: ctxlog.New(opts.Debug),
		db:     opts.DB,
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

	b.gateway = g

	return b, nil
}

func (b *Bot) Connect(ctx context.Context) error {
	defer func() {
		b.logger.Info("Connect() has finished")
	}()
	return b.gateway.Connect(ctx)
}
