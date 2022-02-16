package guild

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/holedaemon/hubris/internal/discord/api/resources"
	"github.com/holedaemon/hubris/internal/discord/types"
	"github.com/holedaemon/hubris/internal/discord/types/null"
)

type ModifyGuildMemberParams struct {
	Nick                       *null.String `json:"nick,omitempty"`
	Roles                      []string     `json:"roles,omitempty"`
	Mute                       *null.Bool   `json:"mute,omitempty"`
	Deaf                       *null.Bool   `json:"deaf,omitempty"`
	ChannelID                  *null.String `json:"channel_id,omitempty"`
	CommunicationDisabledUntil types.Time   `json:"communication_disabled_until,omitempty"`
}

func (r *Resource) ModifyGuildMember(ctx context.Context, id string, p *ModifyGuildMemberParams) (*types.GuildMember, error) {
	var gm *types.GuildMember

	raw, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return gm, r.rc.Patch(ctx, fmt.Sprintf("/guilds/%s/members/%s", r.id, id), &gm, resources.WithBody(raw))
}
