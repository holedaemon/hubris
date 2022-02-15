package gateway

import (
	"github.com/holedaemon/hubris/internal/discord/types"
)

type Ready struct {
	V         int                       `json:"v"`
	User      *types.User               `json:"user"`
	Guilds    []*types.UnavailableGuild `json:"guilds"`
	SessionID string                    `json:"session_id,omitempty"`
	Shard     [2]int                    `json:"shard,omitempty"`
}
