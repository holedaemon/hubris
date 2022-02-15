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
	ID            string          `json:"id"`
	ApplicationID string          `json:"application_id"`
	Type          InteractionType `json:"type"`
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
