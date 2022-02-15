package types

import "github.com/holedaemon/hubris/internal/discord/types/null"

type ActivityType int

const (
	ActivityTypeGame ActivityType = iota
	ActivityTypeStreaming
	ActivityTypeListening
	ActivityTypeWatching
	ActivityTypeCustom
	ActivityTypeCompeting
)

type Activity struct {
	Name string       `json:"name"`
	Type ActivityType `json:"type"`
	URL  null.String  `json:"url,omitempty"`
}
