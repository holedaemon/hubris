package channel

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/holedaemon/hubris/internal/discord/api/resources"
	"github.com/holedaemon/hubris/internal/discord/types"
)

func (r *Resource) Get(ctx context.Context) (*types.Channel, error) {
	var ch *types.Channel

	return ch, r.rc.Get(ctx,
		fmt.Sprintf("/channels/%s", r.id),
		&ch,
	)
}

type createMessageParams struct {
	Content          string                  `json:"content,omitempty"`
	Embeds           []*types.Embed          `json:"embeds,omitempty"`
	MessageReference *types.MessageReference `json:"message_reference,omitempty"`
}

type CreateMessageOption func(*createMessageParams)

func WithMessageContent(c string) CreateMessageOption {
	return func(cmp *createMessageParams) {
		cmp.Content = c
	}
}

func WithEmbed(e *types.Embed) CreateMessageOption {
	return func(cmp *createMessageParams) {
		if cmp.Embeds == nil {
			cmp.Embeds = make([]*types.Embed, 0)
		}

		cmp.Embeds = append(cmp.Embeds, e)
	}
}

func WithMessageReference(r string) CreateMessageOption {
	return func(cmp *createMessageParams) {
		cmp.MessageReference = &types.MessageReference{
			MessageID: r,
		}
	}
}

func (r *Resource) CreateMessage(ctx context.Context, opts ...CreateMessageOption) (*types.Message, error) {
	cmo := new(createMessageParams)

	for _, o := range opts {
		o(cmo)
	}

	var m *types.Message

	raw, err := json.Marshal(cmo)
	if err != nil {
		return nil, err
	}

	return m, r.rc.Post(ctx,
		fmt.Sprintf("/channels/%s/messages", r.id),
		&m,
		resources.WithBody(raw),
	)
}
