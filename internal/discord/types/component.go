package types

type ComponentType int

const (
	ComponentTypeActionRow ComponentType = iota + 1
	ComponentTypeButton
	ComponentTypeSelectMenu
	ComponentTypeTextInput
)

type Component struct {
	Type        ComponentType   `json:"type"`
	CustomID    string          `json:"custom_id,omitempty"`
	Disabled    bool            `json:"disabled,omitempty"`
	Style       ButtonStyle     `json:"style,omitempty"`
	Label       string          `json:"label,omitempty"`
	Emoji       *Emoji          `json:"emoji,omitempty"`
	URL         string          `json:"url,omitempty"`
	Options     []*SelectOption `json:"options,omitempty"`
	Placeholder string          `json:"placeholder,omitempty"`
	MinValues   int             `json:"min_values,omitempty"`
	MaxValues   int             `json:"max_values,omitempty"`
	Components  []Component     `json:"components,omitempty"`
	MinLength   int             `json:"min_length,omitempty"`
	MaxLength   int             `json:"max_length,omitempty"`
	Required    bool            `json:"required,omitempty"`
	Value       string          `json:"value,omitempty"`
}
