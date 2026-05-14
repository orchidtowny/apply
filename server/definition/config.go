package definition

type Config struct {
	Port              int               `json:"port"`
	ServerIp          string            `json:"server_ip"`
	DiscordWebhookUrl string            `json:"discord_webhook_url"`
	ApiKey            string            `json:"api_key"`
	Rules             map[string]string `json:"rules"`
}
