package config

import (
	"bot/internal/bot"
	"bot/internal/client"
	"bot/internal/database"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Logger LoggerConfig
	DB     database.Config
	Bot    bot.Config
	Client client.Config
}

type LoggerConfig struct {
	LogLevel string `envconfig:"LOG_LEVEL" default:"debug"`
}

func Parse() (*Config, error) {
	cfg := Config{}
	logger := LoggerConfig{}
	db := database.Config{}
	telegramBot := bot.Config{}
	clientConfig := client.Config{}
	project := "BOT"

	err := envconfig.Process(project, &logger)
	if err != nil {
		log.Err(err).Msg("logger config error")
		return nil, err
	}

	err = envconfig.Process(project, &db)
	if err != nil {
		log.Err(err).Msg("db config error")
		return nil, err
	}

	err = envconfig.Process(project, &telegramBot)
	if err != nil {
		log.Err(err).Msg("auth config error")
		return nil, err
	}

	err = envconfig.Process(project, &clientConfig)
	if err != nil {
		log.Err(err).Msg("auth config error")
		return nil, err
	}

	cfg.Bot = telegramBot
	cfg.DB = db
	cfg.Logger = logger
	cfg.Client = clientConfig

	return &cfg, nil
}
