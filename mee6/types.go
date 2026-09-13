package mee6

type Response struct {
	Page              int             `json:"page"`
	Guild             Guild           `json:"guild"`
	XpRate            float64         `json:"xp_rate"`
	XpPerMessage      []int           `json:"xp_per_message"`
	RoleRewards       []RoleRewards   `json:"role_rewards"`
	MonetizeOptions   MonetizeOptions `json:"monetize_options"`
	Players           []PlayerType    `json:"players"`
	Player            interface{}     `json:"player"`
	BannerURL         interface{}     `json:"banner_url"`
	IsMember          bool            `json:"is_member"`
	Admin             bool            `json:"admin"`
	UserGuildSettings interface{}     `json:"user_guild_settings"`
	Country           string          `json:"country"`
}

type Guild struct {
	ID                         string `json:"id"`
	Icon                       string `json:"icon"`
	Name                       string `json:"name"`
	Premium                    bool   `json:"premium"`
	AllowJoin                  bool   `json:"allow_join"`
	LeaderboardURL             string `json:"leaderboard_url"`
	InviteLeaderboard          bool   `json:"invite_leaderboard"`
	CommandsPrefix             string `json:"commands_prefix"`
	ApplicationCommandsEnabled bool   `json:"application_commands_enabled"`
}

type RoleRewards struct {
	Rank int  `json:"rank"`
	Role Role `json:"role"`
}

type Role struct {
	Color        int    `json:"color"`
	Hoist        bool   `json:"hoist"`
	Icon         string `json:"icon"`
	ID           string `json:"id"`
	Managed      bool   `json:"managed"`
	Mentionable  bool   `json:"mentionable"`
	Name         string `json:"name"`
	Permissions  int64  `json:"permissions"`
	Position     int    `json:"position"`
	UnicodeEmoji string `json:"unicode_emoji"`
}

type MonetizeOptions struct {
	DisplayPlans        bool `json:"display_plans"`
	ShowcaseSubscribers bool `json:"showcase_subscribers"`
}

type PlayerType struct {
	Avatar               string `json:"avatar"`
	Discriminator        string `json:"discriminator"`
	GuildID              string `json:"guild_id"`
	ID                   string `json:"id"`
	MessageCount         int    `json:"message_count"`
	MonetizeXpBoost      int    `json:"monetize_xp_boost"`
	Username             string `json:"username"`
	Xp                   int    `json:"xp"`
	IsMonetizeSubscriber bool   `json:"is_monetize_subscriber"`
	DetailedXp           []int  `json:"detailed_xp"`
	Level                int    `json:"level"`
}
