package bot

import (
	"github.com/holedaemon/hubris/internal/discord/api"
	"github.com/holedaemon/hubris/internal/discord/gateway"
)

type Bot struct {
	api     *api.Client
	gateway *gateway.Client
}
