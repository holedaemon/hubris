package types

type MessageType int

const (
	MessageTypeDefault MessageType = iota
	MessageTypeRecipientAdd
	MessageTypeRecipientRemove
	MessageTypeCall
	MessageTypeChannelNameChange
	MessageTypeChannelIconChange
	MessageTypeChannelPinnedMessage
	MessageTypeGuildMemberJoin
	MessageTypeUserPremiumGuildSubscription
	MessageTypeUserPremiumGuildSubscriptionTier1
	MessageTypeUserPremiumGuildSubscriptionTier2
	MessageTypeUserPremiumGuildSubscriptionTier3
	MessageTypeChannelFollowAdd
	_
	MessageTypeGuildDiscoveryDisqualified
	MessageTypeGuildDiscoveryRequalified
	MessageTypeGuildDiscoveryGracePeriodInitialWarning
	MessageTypeGuildDiscoveryGracePeriodFinalWarning
	MessageTypeThreadCreated
	MessageTypeReply
	MessageTypeChatInputCommand
	MessageTypeThreadStarterMessage
	MessageTypeGuildInviteReminder
	MessageTypeContextMenuCommand
)

type Message struct {
	ID               string              `json:"id"`
	ChannelID        string              `json:"channel_id"`
	GuildID          string              `json:"guild_id"`
	Author           *User               `json:"author"`
	Member           *GuildMember        `json:"member,omitempty"`
	Content          string              `json:"content"`
	Timestamp        Time                `json:"timestamp"`
	EditedTimestamp  Time                `json:"edited_timestamp"`
	TTS              bool                `json:"tts"`
	MentionEveryone  bool                `json:"mention_everyone"`
	Mentions         []*MessageMention   `json:"mentions"`
	MentionRoles     []string            `json:"mention_roles"`
	MentionChannels  []*ChannelMention   `json:"mention_channels,omitempty"`
	Attachments      []*Attachment       `json:"attachments"`
	Embeds           []*Embed            `json:"embeds"`
	Reactions        []*Reaction         `json:"reaction,omitempty"`
	Pinned           bool                `json:"pinned"`
	WebhookID        string              `json:"webhook_id,omitempty"`
	Type             MessageType         `json:"type"`
	ApplicationID    string              `json:"application_id,omitempty"`
	MessageReference *MessageReference   `json:"message_reference,omitempty"`
	Interaction      *MessageInteraction `json:"interaction,omitempty"`
	Thread           *Channel            `json:"thread,omitempty"`
	Components       []*Component        `json:"components,omitempty"`
}

type MessageMention struct {
	*User
	Member *GuildMember `json:"member"`
}

type MessageInteraction struct {
	ID     string          `json:"id"`
	Type   InteractionType `json:"type"`
	Name   string          `json:"name"`
	User   *User           `json:"user"`
	Member *GuildMember    `json:"member,omitempty"`
}

type MessageReference struct {
	MessageID       string `json:"message_id,omitempty"`
	ChannelID       string `json:"channel_id,omitempty"`
	GuildID         string `json:"guild_id,omitempty"`
	FailIfNotExists bool   `json:"fail_if_not_exists,omitempty"`
}

type ChannelMention struct {
	ID      string      `json:"id"`
	GuildID string      `json:"guild_id"`
	Type    ChannelType `json:"type"`
	Name    string      `json:"name"`
}
