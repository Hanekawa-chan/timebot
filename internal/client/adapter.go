package client

import (
	"bot/internal/app"
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/rs/zerolog"
	"net/http"
)

type adapter struct {
	logger   zerolog.Logger
	client   *http.Client
	geocoder geo.Geocoder
	config   Config
}

func NewAdapter(logger zerolog.Logger, config Config) app.Client {
	return &adapter{
		logger:   logger,
		client:   &http.Client{Timeout: config.Timeout},
		geocoder: openstreetmap.Geocoder(),
		config:   config,
	}
}
