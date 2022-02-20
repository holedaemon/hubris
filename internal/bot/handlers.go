package bot

import (
	"context"

	"github.com/holedaemon/hubris/internal/discord/gateway"
	"github.com/holedaemon/hubris/internal/discord/types"
)

func (b *Bot) handleReady(ctx context.Context, m *gateway.Ready) {
	b.logger.Info("connected to Discord")
}

func (b *Bot) handleMessageCreate(ctx context.Context, m *types.Message) {
	if m.Content == "" {
		return
	}

	c := b.FromMessage(m)

	for ar, fn := range actions {
		if ar.MatchString(m.Content) {
			fn(ctx, c)
			return
		}
	}
}
