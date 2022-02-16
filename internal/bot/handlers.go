package bot

import (
	"context"
	"fmt"

	"github.com/holedaemon/hubris/internal/discord/api/resources/channel"
	"github.com/holedaemon/hubris/internal/discord/types"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

func (b *Bot) handleMessageCreate(ctx context.Context, m *types.Message) {
	if m.Content == "" {
		return
	}

	ch := b.api.Channel(m.ChannelID)

	fmt.Println(regexWaste.MatchString(m.Content), m.Content)

	for re, at := range actions {
		if re.MatchString(m.Content) {
			if at == ActionTypeTimeout && m.Type == types.MessageTypeReply {
				err := b.TimeoutUser(ctx, m.GuildID, m.ReferencedMessage.Author.ID)
				if err != nil {
					_, err = ch.CreateMessage(ctx,
						channel.WithMessageContent("Sorry boss, it ain't workin'"),
					)

					if err != nil {
						ctxlog.Error(ctx, "error sending reply", zap.Error(err))
					}

					return
				}

				_, err = ch.CreateMessage(ctx,
					channel.WithMessageContent("ay jeez do i's looks like some cock weasels to youse? da jobs done boss"),
				)
				if err != nil {
					ctxlog.Error(ctx, "error sending reply", zap.Error(err))
				}
			}
		}
	}
}
