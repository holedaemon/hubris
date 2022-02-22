package types

type InteractionType int

const (
	InteractionTypePing InteractionType = iota + 1
	InteractionTypeApplicationCommand
	InteractionTypeMessageComponent
	InteractionTypeApplicationCommandAutocomplete
	InteractionTypeModalSubmit
)

type Interaction struct {
	ID            string           `json:"id"`
	ApplicationID string           `json:"application_id"`
	Type          InteractionType  `json:"type"`
	Data          *InteractionData `json:"data,omitempty"`
	GuildID       string           `json:"guild_id,omitempty"`
	ChannelID     string           `json:"channel_id,omitempty"`
	Member        *GuildMember     `json:"member,omitempty"`
	User          *User            `json:"user,omitempty"`
	Token         string           `json:"token"`
	Version       int              `json:"version"`
	Message       *Message         `json:"message,omitempty"`
	Locale        string           `json:"locale,omitempty"`
	GuildLocale   string           `json:"guild_locale"`
}

type InteractionData struct {
	ID            string                                     `json:"id"`
	Name          string                                     `json:"name"`
	Type          ApplicationCommandType                     `json:"type"`
	Resolved      *ResolvedData                              `json:"resolved,omitempty"`
	Options       []*ApplicationCommandInteractionDataOption `json:"options,omitempty"`
	CustomID      string                                     `json:"custom_id,omitempty"`
	ComponentType ComponentType                              `json:"component_type,omitempty"`
	Values        []*SelectOption                            `json:"values,omitempty"`
	TargetID      string                                     `json:"target_id,omitempty"`
	Components    []*Component                               `json:"components,omitempty"`
}

type ResolvedData struct {
	Users    map[string]*User        `json:"users,omitempty"`
	Members  map[string]*GuildMember `json:"members,omitempty"`
	Roles    map[string]*Role        `json:"roles,omitempty"`
	Channels map[string]*Channel     `json:"channels,omitempty"`
	Messages map[string]*Message     `json:"messages,omitempty"`
	// Attachments
}

type InteractionCallbackType int

const (
	InteractionCallbackTypePong InteractionCallbackType = iota + 1
	_
	_
	InteractionCallbackTypeChannelMessageWithSource
	InteractionCallbackTypeDeferredChannelMessageWithSource
	InteractionCallbackTypeDeferredUpdateMessages
	InteractionCallbackTypeUpdateMessage
	InteractionCallbackTypeApplicationCommandAutocompleteResult
	InteractionCallbackTypeModal
)

type InteractionResponse struct {
	Type InteractionCallbackType `json:"type"`
	Data interface{}             `json:"data,omitempty"`
}
