package types

type GatewayIntent int

const (
	GatewayIntentGuilds                 GatewayIntent = 1 << 0
	GatewayIntentGuildMembers           GatewayIntent = 1 << 1
	GatewayIntentGuildBans              GatewayIntent = 1 << 2
	GatewayIntentGuildEmojisAndStickers GatewayIntent = 1 << 3
	GatewayIntentGuildIntegrations      GatewayIntent = 1 << 4
	GatewayIntentGuildWebhooks          GatewayIntent = 1 << 5
	GatewayIntentGuildInvites           GatewayIntent = 1 << 6
	GatewayIntentGuildVoiceStates       GatewayIntent = 1 << 7
	GatewayIntentGuildPresences         GatewayIntent = 1 << 8
	GatewayIntentGuildMessages          GatewayIntent = 1 << 9
	GatewayIntentGuildMessageReactions  GatewayIntent = 1 << 10
	GatewayIntentGuildMessageTyping     GatewayIntent = 1 << 11
	GatewayIntentGuildScheduledEvents   GatewayIntent = 1 << 16

	GatewayIntentDirectMessages         GatewayIntent = 1 << 12
	GatewayIntentDirectMessageReactions GatewayIntent = 1 << 13
	GatewayIntentDirectMessageTyping    GatewayIntent = 1 << 14
)

const (
	GatewayIntentDMsAll GatewayIntent = GatewayIntentDirectMessages | GatewayIntentDirectMessageReactions | GatewayIntentDirectMessageTyping

	GatewayIntentPrivilegedGuildsAll GatewayIntent = GatewayIntentGuildPresences | GatewayIntentGuildMembers

	GatewayIntentUnprivilegedGuildsAll GatewayIntent = GatewayIntentGuildBans | GatewayIntentGuilds | GatewayIntentGuildEmojisAndStickers | GatewayIntentGuildIntegrations | GatewayIntentGuildWebhooks | GatewayIntentGuildInvites | GatewayIntentGuildVoiceStates | GatewayIntentGuildPresences | GatewayIntentGuildMessages | GatewayIntentGuildMessageReactions | GatewayIntentGuildMessageTyping | GatewayIntentGuildScheduledEvents

	GatewayIntentGuildsAll GatewayIntent = GatewayIntentPrivilegedGuildsAll | GatewayIntentUnprivilegedGuildsAll
)
