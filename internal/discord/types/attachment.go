package types

import "github.com/holedaemon/hubris/internal/discord/types/null"

type Attachment struct {
	ID          string   `json:"id"`
	Filename    string   `json:"filename"`
	ContentType string   `json:"content_type,omitempty"`
	URL         string   `json:"url"`
	Height      null.Int `json:"height,omitempty"`
	Width       null.Int `json:"width,omitempty"`
}
