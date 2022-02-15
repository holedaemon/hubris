package types

type ApplicationCommandType int

const (
	ApplicationCommandTypeChatInput ApplicationCommandType = iota + 1
	ApplicationCommandTypeUser
	ApplicationCommandTypeMessage
)

type ApplicationCommand struct {
	ID            string                 `json:"id"`
	Type          ApplicationCommandType `json:"type"`
	ApplicationID string                 `json:"application_id"`
	GuildID       string                 `json:"guild_id"`
	Name          string                 `json:"name"`
	Description   string                 `json:"description"`

	DefaultPermission bool   `json:"default_permission,omitempty"`
	Version           string `json:"version"`
}

type ApplicationCommandOptionType int

const (
	ApplicationCommandOptionTypeSubCommand ApplicationCommandOptionType = iota + 1
	ApplicationCommandOptionTypeSubCommandGroup
	ApplicationCommandOptionTypeString
	ApplicationCommandOptionTypeInteger
	ApplicationCommandOptionTypeBoolean
	ApplicationCommandOptionTypeUser
	ApplicationCommandOptionTypeChannel
	ApplicationCommandOptionTypeRole
	ApplicationCommandOptionTypeMentionable
	ApplicationCommandOptionTypeNumber
	ApplicationCommandOptionTypeAttachment
)

type ApplicationCommandOption struct {
	Type         ApplicationCommandOptionType      `json:"type"`
	Name         string                            `json:"name"`
	Description  string                            `json:"description"`
	Required     bool                              `json:"required,omitempty"`
	Choices      []*ApplicationCommandOptionChoice `json:"choices,omitempty"`
	Options      []*ApplicationCommandOption       `json:"options,omitempty"`
	ChannelTypes []ChannelType                     `json:"channel_types,omitempty"`
	MinValue     float64                           `json:"min_value,omitempty"`
	MaxValue     float64                           `json:"max_value,omitempty"`
	Autocomplete bool                              `json:"autocomplete,omitempty"`
}

type ApplicationCommandOptionChoice struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

func (c *ApplicationCommandOptionChoice) String() string {
	s, ok := c.Value.(string)
	if !ok {
		return ""
	}

	return s
}

func (c *ApplicationCommandOptionChoice) Integer() int {
	i, ok := c.Value.(int)
	if !ok {
		return 0
	}
	return i
}

func (c *ApplicationCommandOptionChoice) Double() float64 {
	i, ok := c.Value.(float64)
	if !ok {
		return 0
	}

	return i
}

type ApplicationCommandInteractionDataOption struct {
	Name    string                                     `json:"name"`
	Type    ApplicationCommandOptionType               `json:"type"`
	Value   interface{}                                `json:"value,omitempty"`
	Options []*ApplicationCommandInteractionDataOption `json:"options,omitempty"`
	Focused bool                                       `json:"focused,omitempty"`
}
