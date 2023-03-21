package app

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

type service struct {
	logger   zerolog.Logger
	config   Config
	db       Database
	client   *http.Client
	geocoder geo.Geocoder
}

func NewService(logger zerolog.Logger, config Config, database Database) Service {
	return &service{
		logger:   logger,
		config:   config,
		db:       database,
		client:   &http.Client{Timeout: 30 * time.Second},
		geocoder: openstreetmap.Geocoder(),
	}
}
