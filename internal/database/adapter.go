package database

import (
	"bot/internal/app"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type adapter struct {
	logger zerolog.Logger
	config Config
	db     *sqlx.DB
}

func NewAdapter(logger zerolog.Logger, config Config) (app.Database, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Name, config.Password)
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		logger.Err(err).Msg("db connect")
		return nil, err
	}

	instance, err := postgres.WithInstance(db.DB, &postgres.Config{DatabaseName: config.Name, SchemaName: "public"})
	if err != nil {
		logger.Err(err).Msg("db instance")
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		config.MigrationsURL, config.Name, instance)
	if err != nil {
		logger.Err(err).Msg("db migration create")
		return nil, err
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Err(err).Msg("db migrate up")
		return nil, err
	}

	a := &adapter{
		logger: logger,
		config: config,
		db:     db,
	}
	return a, nil
}
