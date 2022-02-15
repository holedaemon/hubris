package types

import "github.com/holedaemon/hubris/internal/discord/types/null"

type Role struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Color        int         `json:"color"`
	Hoist        bool        `json:"hoist"`
	Icon         null.String `json:"icon,omitempty"`
	UnicodeEmoji null.String `json:"unicode_emoji"`
	Position     int         `json:"position"`
	Permissions  string      `json:"permissions"`
	Managed      bool        `json:"managed"`
	Mentionable  bool        `json:"mentionable"`
	Tags         *RoleTag    `json:"tags,omitempty"`
}

type RoleTag struct {
	BotID         string `json:"bot_id,omitempty"`
	IntegrationID string `json:"integration_id"`
}
