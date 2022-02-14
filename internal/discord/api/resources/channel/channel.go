package channel

import (
	"context"
	"fmt"

	"github.com/holedaemon/hubris/internal/discord/types"
)

func (r *Resource) Get(ctx context.Context) (*types.Channel, error) {
	var ch *types.Channel

	return ch, r.rc.Get(ctx,
		fmt.Sprintf("/channels/%s", r.id),
		&ch,
	)
}
