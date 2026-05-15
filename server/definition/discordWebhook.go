package definition

type DiscordWebhook struct {
	Username string                `json:"username"`
	Embeds   []DiscordWebhookEmbed `json:"embeds"`
}

type DiscordWebhookEmbed struct {
	Title     string                     `json:"title"`
	Timestamp string                     `json:"timestamp"`
	Color     int                        `json:"color"`
	Author    *DiscordWebhookEmbedAuthor `json:"author"`
	Fields    []DiscordWebhookEmbedField `json:"fields"`
}

type DiscordWebhookEmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type DiscordWebhookEmbedAuthor struct {
	Name    string `json:"name"`
	IconUrl string `json:"icon_url"`
}
