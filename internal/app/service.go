package app

import (
	"github.com/rs/zerolog"
)

type service struct {
	logger zerolog.Logger
	db     Database
	client Client
}

func NewService(logger zerolog.Logger, database Database, client Client) Service {
	return &service{
		logger: logger,
		db:     database,
		client: client,
	}
}
