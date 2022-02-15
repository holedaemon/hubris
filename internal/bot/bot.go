package bot

import (
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
}

type Options struct {
	Logger *zap.Logger
	Debug  bool
	Token  string
}

func New(opts *Options) (*Bot, error) {
	if opts.Token == "" {
		return nil, fmt.Errorf("bot: token is blank")
	}

	b := &Bot{
		logger: opts.Logger,
	}

	if b.logger == nil {
		b.logger = ctxlog.New(opts.Debug)
	}

	a, err := api.New(opts.Token)
	if err != nil {
		return nil, err
	}

	b.api = a

	g, err := gateway.New(opts.Token)
	if err != nil {
		return nil, err
	}

	b.gateway = g

	return b, nil
}
