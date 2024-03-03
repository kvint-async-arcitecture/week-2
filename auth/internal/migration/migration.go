package migration

import (
	"context"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	"auth/internal/log"
)

var (
	ErrOpenDBWithDriver = errors.New("open db error")
	ErrDBPing           = errors.New("ping db error")
	ErrDBMigration      = errors.New("migration db error")
)

type gooseLogger struct{}

func (l gooseLogger) Fatalf(format string, v ...interface{}) {
	log.Default().Error().Msgf(format, v...)
}

func (l gooseLogger) Printf(format string, v ...interface{}) {
	log.Default().Info().Msgf(format, v...)
}

func Up(ctx context.Context, migrationURL string) error {
	log.Default().Info().Msg("started db migration")
	defer log.Default().Info().Msg("finished db migration")

	goose.SetLogger(new(gooseLogger))

	db, err := goose.OpenDBWithDriver("postgres", migrationURL)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrOpenDBWithDriver, err.Error())
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Default().Err(err).Msg("can't close migration db connection")
		}
	}()

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("%w: %s", ErrDBPing, err.Error())
	}

	if err := goose.UpContext(ctx, db, "scripts/db/migrations"); err != nil {
		return fmt.Errorf("%w: %s", ErrDBMigration, err.Error())
	}

	return nil
}
