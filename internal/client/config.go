package client

import "time"

type Config struct {
	TimeToken string        `envconfig:"CLIENT_TIME_TOKEN"`
	TimeURL   string        `envconfig:"CLIENT_TIME_URL"`
	Timeout   time.Duration `envconfig:"CLIENT_HTTP_CLIENT_TIMEOUT" default:"2m"`
}
