package types

import "github.com/holedaemon/hubris/internal/discord/types/null"

type VerificationLevel int

const (
	VerificationLevelNone VerificationLevel = iota
	VerificationLevelLow
	VerificationLevelMedium
	VerificationLevelHigh
	VerificationLevelVeryHigh
)

type DefaultMessageNotificationLevel int

const (
	DefaultMessageNotificationLevelAllMessages DefaultMessageNotificationLevel = iota
	DefaultMessageNotificationLevelOnlyMentions
)

type ExplicitContentFilterLevel int

const (
	ExplicitContentFilterLevelDisabled ExplicitContentFilterLevel = iota
	ExplicitContentFilterLevelMembersWithoutRoles
	ExplicitContentFilterLevelAllMembers
)

type MFALevel int

const (
	MFALevelNone MFALevel = iota
	MFALevelElevated
)

type GuildNSFWLevel int

const (
	GuildNSFWLevelDefault GuildNSFWLevel = iota
	GuildNSFWLevelExplicit
	GuildNSFWLevelSafe
	GuildNSFWLevelAgeRestricted
)

type PremiumTier int

const (
	PremiumTierNone PremiumTier = iota
	PremiumTier1
	PremiumTier2
	PremiumTier3
)

type SystemChannelFlag int

const (
	SystemChannelFlagSuppressJoinNotifications          SystemChannelFlag = 1 << 0
	SystemChannelFlagSuppressPremiumSubscriptions       SystemChannelFlag = 1 << 1
	SystemChannelFlagSuppressGuildReminderNotifications SystemChannelFlag = 1 << 2
	SystemChannelFlagSuppressJoinNotificationReplies    SystemChannelFlag = 1 << 3
)

type GuildFeature string

const (
	GuildFeatureAnimatedIcon                  = "ANIMATED_ICON"
	GuildFeatureBanner                        = "BANNER"
	GuildFeatureCommerce                      = "COMMERCE"
	GuildFeatureCommunity                     = "COMMUNITY"
	GuildFeatureDiscoverable                  = "DISCOVERABLE"
	GuildFeatureFeaturable                    = "FEATURABLE"
	GuildFeatureInviteSplash                  = "INVITE_SPLASH"
	GuildFeatureMemberVerificationGateEnabled = "MEMBER_VERIFICATION_GATE_ENABLED"
	GuildFeatureMonetizationEnabled           = "MONETIZATION_ENABLED"
	GuildFeatureMoreStickers                  = "MORE_STICKERS"
	GuildFeatureNews                          = "NEWS"
	GuildFeaturePartnered                     = "PARTNERED"
	GuildFeaturePreviewEnabled                = "PREVIEW_ENABLED"
	GuildFeaturePrivateThreads                = "PRIVATE_THREADS"
	GuildFeatureRoleIcons                     = "ROLE_ICONS"
	GuildFeatureSevenDayThreadArchive         = "SEVEN_DAY_THREAD_ARCHIVE"
	GuildFeatureThreeDayThreadArchive         = "THREE_DAY_THREAD_ARCHIVE"
	GuildFeatureTicketedEventsEnabled         = "TICKETED_EVENTS_ENABLED"
	GuildFeatureVanityURL                     = "VANITY_URL"
	GuildFeatureVerified                      = "VERIFIED"
	GuildFeatureVIPRegions                    = "VIP_REGIONS"
	GuildFeatureWelcomeScreenEnabled          = "WELCOME_SCREEN_ENABLED"
)

type Guild struct {
	ID                          string                          `json:"id"`
	Name                        string                          `json:"name"`
	Icon                        null.String                     `json:"icon"`
	IconHash                    null.String                     `json:"icon_hash,omitempty"`
	Splash                      null.String                     `json:"splash"`
	DiscoverySplash             null.String                     `json:"discovery_splash"`
	Owner                       bool                            `json:"owner,omitempty"`
	OwnerID                     string                          `json:"owner_id,omitempty"`
	Permissions                 string                          `json:"permissions,omitempty"`
	Region                      null.String                     `json:"region,omitempty"`
	AFKChannelID                null.String                     `json:"afk_channel_id"`
	AFKTimeout                  int                             `json:"afk_timeout"`
	WidgetEnabled               bool                            `json:"widget_enabled,omitempty"`
	WidgetChannelID             null.String                     `json:"widget_channe_id,omitempty"`
	VerificationLevel           VerificationLevel               `json:"verification_level"`
	DefaultMessageNotifications DefaultMessageNotificationLevel `json:"default_message_notifications"`
	ExplicitContentFilter       ExplicitContentFilterLevel      `json:"explicit_content_filter"`
	// Roles []*Role `json:"roles"`
	// Emojis []*Emoji `json:"emojis"`
	Features           []GuildFeature    `json:"features"`
	MFALevel           MFALevel          `json:"mfa_level"`
	ApplicationID      null.String       `json:"application_id"`
	SystemChannelID    null.String       `json:"system_channel_id"`
	SystemChannelFlags SystemChannelFlag `json:"system_channel_flags"`
	RulesChannelID     null.String       `json:"rules_channel_id"`
	JoinedAt           Time              `json:"joined_at,omitempty"`
	Large              bool              `json:"large,omitempty"`
	Unavailable        bool              `json:"unavailable,omitempty"`
	MemberCount        int               `json:"member_count,omitempty"`
	// VoiceStates []*VoiceState `json:"voice_states,omitempty"`
	Members                  []*GuildMember `json:"members,omitempty"`
	Channels                 []*Channel     `json:"channels,omitempty"`
	Threads                  []*Channel     `json:"threads,omitempty"`
	Presences                []*Presence    `json:"presences,omitempty"`
	MaxPresences             null.Int       `json:"max_presences,omitempty"`
	MaxMembers               int            `json:"max_members,omitempty"`
	VanityURLCode            null.String    `json:"vanity_url_code"`
	Description              null.String    `json:"description"`
	Banner                   null.String    `json:"banner"`
	PremiumTier              PremiumTier    `json:"premium_tier,omitempty"`
	PremiumSubscriptionCount int            `json:"premium_subscription_count"`
	PreferredLocale          string         `json:"preferred_locale,omitempty"`
	PublicUpdatesChannelID   null.String    `json:"public_updates_channel_id"`
	MaxVideoChannelUsers     int            `json:"max_video_channel_users,omitempty"`
	ApproximateMemberCount   int            `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount int            `json:"approximate_presence_count,omitempty"`
	NSFWLevel                GuildNSFWLevel `json:"nsfw_level"`
}

type GuildMember struct {
	User                       *User       `json:"user,omitempty"`
	Nick                       null.String `json:"nick,omitempty"`
	Avatar                     null.String `json:"avatar,omitempty"`
	Roles                      []string    `json:"roles"`
	JoinedAt                   Time        `json:"joined_at"`
	PremiumSince               Time        `json:"premium_since"`
	Deaf                       bool        `json:"deaf"`
	Mute                       bool        `json:"mute"`
	Pending                    bool        `json:"pending,omitempty"`
	Permissions                string      `json:"permissions,omitempty"`
	CommunicationDisabledUntil Time        `json:"communication_disabled_until,omitempty"`
}
