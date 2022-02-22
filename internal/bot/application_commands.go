package bot

import (
	"context"
	"fmt"

	"github.com/holedaemon/hubris/internal/discord/api/resources/interaction"
	"github.com/holedaemon/hubris/internal/discord/types"
)

type CommandContext struct {
	ApplicationID string
	GuildID       string
	ChannelID     string

	Options  []*types.ApplicationCommandInteractionDataOption
	Resolved *types.ResolvedData

	User   *types.User
	Member *types.GuildMember

	i *interaction.Resource
}

func NewCommandContext(i *types.Interaction) *CommandContext {
	return &CommandContext{
		ApplicationID: i.ApplicationID,
		GuildID:       i.GuildID,
		ChannelID:     i.ChannelID,
		Options:       i.Data.Options,
		Resolved:      i.Data.Resolved,
		User:          i.User,
		Member:        i.Member,
	}
}

func (c *CommandContext) Option(n int) *types.ApplicationCommandInteractionDataOption {
	if len(c.Options) == 0 {
		return nil
	}

	if len(c.Options) < n {
		return nil
	}

	return c.Options[n]
}

func (c *CommandContext) Reply(ctx context.Context, msg string, args ...interface{}) error {
	r := interaction.NewMessageResponse(
		&types.Message{
			Content: fmt.Sprintf(msg, args...),
		},
	)

	return c.i.CreateInteractionResponse(ctx, r)
}

func (c *CommandContext) DeferredReply(ctx context.Context, msg string, args ...interface{}) error {
	r := interaction.NewMessageResponse(
		&types.Message{
			Content: fmt.Sprintf(msg, args...),
		},
	)

	return c.i.CreateInteractionResponse(ctx, r)
}

type ApplicationCommand func(context.Context, *CommandContext) error
