package store

import (
	"context"

	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5/pgxpool"

	"auth/internal/config"
)

func NewDBPool(ctx context.Context) (*pgxpool.Pool, error) {
	cfg := config.Default()

	cfg2, err := pgxpool.ParseConfig(cfg.AuthDB.URL)
	if err != nil {
		return nil, err
	}

	cfg2.MaxConns = cfg.AuthDB.PoolMaxConnections
	cfg2.HealthCheckPeriod = cfg.AuthDB.HealthCheckPeriod
	cfg2.ConnConfig.Tracer = otelpgx.NewTracer(otelpgx.WithIncludeQueryParameters())

	pool, err := pgxpool.NewWithConfig(ctx, cfg2)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
