package bot

import (
	"context"

	"github.com/holedaemon/hubris/internal/database/models"
	"github.com/holedaemon/hubris/internal/discord/gateway"
	"github.com/holedaemon/hubris/internal/discord/types"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

func (b *Bot) handleReady(ctx context.Context, m *gateway.Ready) {
	b.logger.Info("connected to Discord")
}

func (b *Bot) handleGuildCreate(ctx context.Context, g *types.Guild) {
	l := b.logger.With(zap.String("guild_id", g.ID))
	l.Info("received GUILD_CREATE")

	exists, err := models.Guilds(qm.Where("guild_id = ?", g.ID)).Exists(ctx, b.db)
	if err != nil {
		l.Error("error checking for guild in database", zap.Error(err))
		return
	}

	if !exists {
		gd := &models.Guild{
			GuildID: g.ID,
		}

		if err := gd.Insert(ctx, b.db, boil.Infer()); err != nil {
			l.Error("error creating new record for guild", zap.Error(err))
			return
		}
	}
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

func (b *Bot) handleInteractionCreate(ctx context.Context, i *types.Interaction) {

}
