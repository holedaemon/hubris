package types

import "github.com/holedaemon/hubris/internal/discord/types/null"

type Emoji struct {
	ID            string      `json:"id"`
	Name          null.String `json:"name"`
	Roles         []string    `json:"roles,omitempty"`
	User          *User       `json:"user,omitempty"`
	RequireColons bool        `json:"require_colons,omitempty"`
	Managed       bool        `json:"managed,omitempty"`
	Animated      bool        `json:"animated,omitempty"`
	Available     bool        `json:"available,omitempty"`
}

type Reaction struct {
	Count int    `json:"count"`
	Me    bool   `json:"me"`
	Emoji *Emoji `json:"emoji"`
}
