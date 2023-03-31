package main

import (
	"bot/internal/app"
	"bot/internal/bot"
	"bot/internal/client"
	"bot/internal/config"
	"bot/internal/database"
	"github.com/rs/zerolog"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("Loading bot")

	// Parse all configs form env
	cfg, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	// Parse log level
	level, err := zerolog.ParseLevel(cfg.Logger.LogLevel)
	_ = err
	if err != nil {
		log.Fatal(err)
	}

	// Initializations
	logger := zerolog.New(os.Stdout).Level(level)

	db, err := database.NewAdapter(logger, cfg.DB)
	if err != nil {
		logger.Fatal().Err(err).Msg("Database init")
	}

	clientAdapter := client.NewAdapter(logger, cfg.Client)

	service := app.NewService(logger, db, clientAdapter)

	telegramBot, err := bot.NewAdapter(logger, cfg.Bot, service)
	if err != nil {
		logger.Fatal().Err(err).Msg("Bot init")
	}

	logger.Info().Msg("Initialized everything")

	// Channels for errors and os signals
	stop := make(chan error, 1)
	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig, syscall.SIGINT, syscall.SIGTERM)

	// Receive errors form start bot func into error channel
	go func(stop chan<- error) {
		stop <- telegramBot.Start()
	}(stop)

	// Blocking select
	select {
	case sig := <-osSig:
		logger.Info().Msgf("Received os syscall signal %v", sig)
	case err := <-stop:
		logger.Error().Err(err).Msg("Received Error signal")
	}

	// Shutdown code
	logger.Info().Msg("Shutting down...")

	telegramBot.Shutdown()

	logger.Info().Msg("Shutdown - success")
}
