package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a adapter) Start() error {
	a.logger.Info().Msg("Bot started")

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = a.config.Timeout

	updates := a.bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID
			switch update.Message.Command() {
			case "time":
				args := update.Message.CommandArguments()

				timeInCity, err := a.service.GetTimeByCity(chatID, args)
				if err != nil {
					a.logger.Error().Err(err).Msg("Get time by place")
					msg := tgbotapi.NewMessage(chatID, "Maybe this place doesn't exist?")
					if _, err = a.bot.Send(msg); err != nil {
						a.logger.Error().Err(err).Msg("Message send")
					}
					break
				}

				msg := tgbotapi.NewMessage(chatID, timeInCity)
				if _, err = a.bot.Send(msg); err != nil {
					a.logger.Error().Err(err).Msg("Message send")
					break
				}
			case "location":
				args := update.Message.CommandArguments()

				stats, err := a.service.GetLocation(chatID, args)
				if err != nil {
					a.logger.Error().Err(err).Msg("Get stats")
					msg := tgbotapi.NewMessage(chatID, "Maybe this place doesn't exist?")
					if _, err = a.bot.Send(msg); err != nil {
						a.logger.Error().Err(err).Msg("Message send")
					}
					break
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, stats)
				if _, err = a.bot.Send(msg); err != nil {
					a.logger.Error().Err(err).Msg("Message send")
					break
				}
			case "stats":
				stats, err := a.service.GetStats(chatID)
				if err != nil {
					a.logger.Error().Err(err).Msg("Get stats")
					msg := tgbotapi.NewMessage(chatID, "No stats to send, sorry!")
					if _, err = a.bot.Send(msg); err != nil {
						a.logger.Error().Err(err).Msg("Message send")
					}
					break
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, stats)
				if _, err = a.bot.Send(msg); err != nil {
					a.logger.Error().Err(err).Msg("Message send")
					break
				}
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "There are 3 commands:\n"+
					"/time [place]\n"+
					"\tDescription: shows time in specified place\n"+
					"/location [place]\n"+
					"\tDescription: shows location of specified place in latitude and longitude\n"+
					"/stats\n"+
					"\tDescription: shows yours statistics")
				if _, err := a.bot.Send(msg); err != nil {
					a.logger.Error().Err(err).Msg("Message send")
					break
				}
			}
		}
	}

	return nil
}

func (a adapter) Shutdown() {
	a.bot.StopReceivingUpdates()
}
