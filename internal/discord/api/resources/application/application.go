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

func (r *Resource) GetGlobalApplicationCommands(ctx context.Context) ([]*types.ApplicationCommand, error) {
	var ap []*types.ApplicationCommand

	return ap, r.rc.Get(ctx,
		fmt.Sprintf("/applications/%s/commands", r.id),
		&ap,
	)
}

func (r *Resource) GetGuildApplicationCommands(ctx context.Context, gid string) ([]*types.ApplicationCommand, error) {
	var ap []*types.ApplicationCommand

	return ap, r.rc.Get(ctx,
		fmt.Sprintf("/applications/%s/guilds/%s/commands", r.id, gid),
		&ap,
	)
}

func (r *Resource) DeleteGuildApplicationCommand(ctx context.Context, gid, cid string) error {
	return r.rc.Delete(ctx,
		fmt.Sprintf("/applications/%s/guilds/%s/commands/%s", r.id, gid, cid),
		nil,
	)
}

func (r *Resource) DeleteGlobalApplicationCommand(ctx context.Context, cid string) error {
	return r.rc.Delete(ctx,
		fmt.Sprintf("/applications/%s/commands/%s", r.id, cid),
		nil,
	)
}

func (r *Resource) BulkOverwriteGlobalApplicationCommands(ctx context.Context, ac []*types.ApplicationCommand) ([]*types.ApplicationCommand, error) {
	raw, err := json.Marshal(ac)
	if err != nil {
		return nil, err
	}

	var rc []*types.ApplicationCommand
	return rc, r.rc.Put(ctx,
		fmt.Sprintf("/applications/%s/commands", r.id),
		&rc,
		resources.WithBody(raw),
	)
}

func (r *Resource) BulkOverwriteGuildApplicationCommands(ctx context.Context, gid string, ac []*types.ApplicationCommand) ([]*types.ApplicationCommand, error) {
	raw, err := json.Marshal(ac)
	if err != nil {
		return nil, err
	}

	var rc []*types.ApplicationCommand
	return rc, r.rc.Put(ctx,
		fmt.Sprintf("/applications/%s/guilds/%s/commands", r.id, gid),
		&rc,
		resources.WithBody(raw),
	)
}
