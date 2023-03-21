package bot

type Config struct {
	TelegramToken string `envconfig:"TELEGRAM_TOKEN"`
	Timeout       int    `envconfig:"TIMEOUT" default:"30"`
}
