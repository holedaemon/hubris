package types

import "github.com/holedaemon/hubris/internal/discord/types/null"

type ChannelType int

const (
	ChannelTypeGuildText ChannelType = iota
	ChannelTypeDM
	ChannelTypeGuildVoice
	ChannelTypeGuildCategory
	ChannelTypeGuildNews
	ChannelTypeGuildStore
	ChannelTypeGuildNewsThread
	ChannelTypeGuildPublicThread
	ChannelTypePrivateThread
	ChannelTypeStageVoice
)

type VideoQualityMode int

const (
	VideoQualityModeAuto VideoQualityMode = iota + 1
	VideoQualityModeFull
)

type Channel struct {
	ID                         string           `json:"id"`
	Type                       ChannelType      `json:"type"`
	GuildID                    string           `json:"guild_id,omitempty"`
	Position                   int              `json:"position,omitempty"`
	PermissionOverwrites       []*Overwrite     `json:"permission_overwrites,omitempty"`
	Name                       string           `json:"name,omitempty"`
	Topic                      null.String      `json:"topic,omitempty"`
	NSFW                       bool             `json:"nsfw,omitempty"`
	LastMessageID              null.String      `json:"last_message_id"`
	Bitrate                    int              `json:"bitrate,omitempty"`
	UserLimit                  int              `json:"user_limit,omitempty"`
	RateLimitPerUser           int              `json:"rate_limit_per_user,omitempty"`
	Recipients                 []*User          `json:"recipients,omitempty"`
	Icon                       null.String      `json:"icon,omitempty"`
	OwnerID                    string           `json:"owner_id,omitempty"`
	ApplicationID              string           `json:"application_id,omitempty"`
	ParentID                   null.String      `json:"parent_id,omitempty"`
	LastPinTimestamp           Time             `json:"last_pin_timestamp,omitempty"`
	RTCRegion                  null.String      `json:"rtc_region,omitempty"`
	VideoQualityMode           VideoQualityMode `json:"video_quality_mode,omitempty"`
	MessageCount               int              `json:"message_count,omitempty"`
	MemberCount                int              `json:"member_count,omitempty"`
	ThreadMetadata             *ThreadMetadata  `json:"thread_metadata,omitempty"`
	Member                     *ThreadMember    `json:"thread_member,omitempty"`
	DefaultAutoArchiveDuration int              `json:"default_auto_archive_duration,omitempty"`
	Permissions                string           `json:"permissions,omitempty"`
}

type Overwrite struct {
	ID    string `json:"id"`
	Type  int    `json:"type"`
	Allow string `json:"allow"`
	Deny  string `json:"deny"`
}

type ThreadMetadata struct {
	Archived            bool `json:"archived"`
	AutoArchiveDuration int  `json:"auto_archive_duration"`
	ArchiveTimestamp    Time `json:"archive_timestamp"`
	Locked              bool `json:"locked"`
	Invitable           bool `json:"invitable,omitempty"`
	CreateTimestamp     Time `json:"create_timestamp,omitempty"`
}

type ThreadMember struct {
	ID            string `json:"id,omitempty"`
	UserID        string `json:"user_id,omitempty"`
	JoinTimestamp Time   `json:"join_timestamp,omitempty"`
	Flags         int    `json:"flags"`
}
