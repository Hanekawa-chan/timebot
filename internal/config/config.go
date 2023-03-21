package config

import (
	"bot/internal/app"
	"bot/internal/bot"
	"bot/internal/database"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Logger  LoggerConfig
	DB      database.Config
	Bot     bot.Config
	Service app.Config
}

type LoggerConfig struct {
	LogLevel string `envconfig:"LOG_LEVEL" default:"debug"`
}

func Parse() (*Config, error) {
	cfg := Config{}
	logger := LoggerConfig{}
	db := database.Config{}
	telegramBot := bot.Config{}
	service := app.Config{}
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

	err = envconfig.Process(project, &service)
	if err != nil {
		log.Err(err).Msg("auth config error")
		return nil, err
	}

	cfg.Bot = telegramBot
	cfg.DB = db
	cfg.Logger = logger
	cfg.Service = service

	return &cfg, nil
}
