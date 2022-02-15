package types

import "github.com/holedaemon/hubris/internal/discord/types/null"

type StatusType string

const (
	StatusTypeOnline    StatusType = "online"
	StatusTypeDND       StatusType = "dnd"
	StatusTypeIdle      StatusType = "idle"
	StatusTypeInvisible StatusType = "invisible"
	StatusTypeOffline   StatusType = "offline"
)

type Presence struct {
	Since      null.Int    `json:"since"`
	Activities []*Activity `json:"activities"`
	Status     StatusType  `json:"status"`
	AFK        bool        `json:"afk"`
}
