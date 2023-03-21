package bot

import (
	"bot/internal/app"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
)

type adapter struct {
	logger  zerolog.Logger
	config  Config
	bot     *tgbotapi.BotAPI
	service app.Service
}

func NewAdapter(logger zerolog.Logger, config Config, service app.Service) (app.Bot, error) {
	bot, err := tgbotapi.NewBotAPI(config.TelegramToken)
	if err != nil {
		logger.Error().Err(err).Msg("Bot create")
		return nil, err
	}

	a := &adapter{
		logger:  logger,
		config:  config,
		service: service,
		bot:     bot,
	}

	return a, nil
}
