package interaction

import (
	"context"
	"fmt"

	"github.com/holedaemon/hubris/internal/discord/types"
)

func (r *Resource) CreateInteractionResponse(ctx context.Context, res *types.InteractionResponse) error {
	return r.rc.Post(
		ctx,
		fmt.Sprintf("/interactions/%s/%s/callback", r.id, r.token),
		nil,
	)
}

func (r *Resource) GetOriginalInteractionResponse(ctx context.Context) (*types.Message, error) {
	var m *types.Message

	return m, r.rc.Get(
		ctx,
		fmt.Sprintf("/interactions/%s/%s/messages/@original", r.appID, r.token),
		&m,
	)
}
