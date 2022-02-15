package application

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/holedaemon/hubris/internal/discord/api/resources"
	"github.com/holedaemon/hubris/internal/discord/types"
)

type CreateApplicationCommandParams struct {
	Name              string                            `json:"name"`
	Description       string                            `json:"description"`
	Options           []*types.ApplicationCommandOption `json:"options,omitempty"`
	DefaultPermission bool                              `json:"default_permission,omitempty"`
	Type              types.ApplicationCommandType      `json:"type,omitempty"`
}

func (r *Resource) CreateGuildApplicationCommand(ctx context.Context, id string, p *CreateApplicationCommandParams) (*types.ApplicationCommand, error) {
	var ap *types.ApplicationCommand

	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return ap, r.rc.Post(ctx,
		fmt.Sprintf("/applications/%s/guilds/%s/commands", r.id, id),
		&ap,
		resources.WithBody(data),
	)
}
