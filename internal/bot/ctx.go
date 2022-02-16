package bot

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/holedaemon/hubris/internal/discord/api/resources/channel"
	"github.com/holedaemon/hubris/internal/discord/api/resources/guild"
	"github.com/holedaemon/hubris/internal/discord/types"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

var errNotReply = errors.New("not a reply")

type Context struct {
	Message *types.Message

	Channel *channel.Resource
	Guild   *guild.Resource
}

func (b *Bot) FromMessage(m *types.Message) *Context {
	return &Context{
		Message: m,

		Channel: b.api.Channel(m.ChannelID),
		Guild:   b.api.Guild(m.GuildID),
	}
}

func (c *Context) Reply(ctx context.Context, msg string, args ...interface{}) {
	_, err := c.Channel.CreateMessage(ctx,
		channel.WithMessageContent(
			fmt.Sprintf(msg, args...),
		),
		channel.WithMessageReference(
			c.Message.ID,
		),
	)

	if err != nil {
		ctxlog.Error(ctx, "error sending reply to message", zap.String("message_id", c.Message.ID), zap.Error(err))
	}
}

func (c *Context) Timeout(ctx context.Context) error {
	if c.Message.Type != types.MessageTypeReply {
		return errNotReply
	}

	f := rand.Intn(10)
	t := time.Now().Add(time.Duration(f) * time.Minute)

	ctx = ctxlog.With(ctx, zap.String("user_id", c.Message.ReferencedMessage.Author.ID), zap.String("guild_id", c.Message.GuildID))
	ctxlog.Info(ctx, "attempting to time out user", zap.Time("until", t))

	_, err := c.Guild.ModifyGuildMember(ctx,
		c.Message.ReferencedMessage.Author.ID,
		&guild.ModifyGuildMemberParams{
			CommunicationDisabledUntil: types.TimeToTime(t),
		},
	)
	if err != nil {
		ctxlog.Error(ctx, "error timing out user", zap.String("user_id", c.Message.Author.ID), zap.Error(err))
	}
	return err
}
