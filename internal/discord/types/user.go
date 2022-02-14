package types

import "github.com/holedaemon/hubris/internal/discord/types/null"

type PremiumType int

const (
	PremiumTypeNone PremiumType = iota
	PremiumTypeNitroClassic
	PremiumTypeNitro
)

type UserFlag int

const (
	UserFlagNone                  UserFlag = 0
	UserFlagStaff                 UserFlag = 1 << 0
	UserFlagPartner               UserFlag = 1 << 1
	UserFlagHypesquad             UserFlag = 1 << 2
	UserFlagBugHunterLevel1       UserFlag = 1 << 3
	UserFlagHypesquadOnlineHouse1 UserFlag = 1 << 6
	UserFlagHypesquadOnlineHouse2 UserFlag = 1 << 7
	UserFlagHypesquadOnlineHouse3 UserFlag = 1 << 8
	UserFlagPremiumEarlySupporter UserFlag = 1 << 9
	UserFlagTeamPseudoUser        UserFlag = 1 << 10
	UserFlagBugHunterLevel2       UserFlag = 1 << 14
	UserFlagVerifiedBot           UserFlag = 1 << 16
	UserFlagVerifiedDeveloper     UserFlag = 1 << 17
	UserFlagCertifiedModerator    UserFlag = 1 << 18
	UserFlagBotHTTPInteractions   UserFlag = 1 << 19
)

type User struct {
	ID            string      `json:"id"`
	Username      string      `json:"username"`
	Discriminator string      `json:"discriminator"`
	Avatar        null.String `json:"avatar"`
	Bot           bool        `json:"bot,omitempty"`
	System        bool        `json:"system,omitempty"`
	MFAEnabled    bool        `json:"mfa_enabled"`
	Banner        null.String `json:"banner,omitempty"`
	AccentColor   null.Int    `json:"accent_color,omitempty"`
	Locale        string      `json:"locale,omitempty"`
	Verified      bool        `json:"verified,omitempty"`
	Email         null.String `json:"email,omitempty"`
	Flags         int         `json:"flags,omitempty"`
	PremiumType   PremiumType `json:"premium_type,omitempty"`
	PublicFlags   UserFlag    `json:"public_flags,omitempty"`
}
