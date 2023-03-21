PROJECT_NAME := bot
PROJECT := github.com/Hanekawa-chan/auth

build:
	CGO_ENABLED=0 go build -o ./bin/$(PROJECT_NAME) ./cmd/$(PROJECT_NAME)