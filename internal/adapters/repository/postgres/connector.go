package postgres

import (
	"context"
	"fmt"

	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateDBPool(
	ctx context.Context,
	cfg Config,
	serviceName string,
) (*pgxpool.Pool, error) {
	pgxConfig, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("create connection pool: %w", err)
	}

	pgxConfig.ConnConfig.Tracer = otelpgx.NewTracer(otelpgx.WithIncludeQueryParameters())

	pgxConfig.ConnConfig.RuntimeParams["timezone"] = cfg.Timezone
	pgxConfig.ConnConfig.RuntimeParams["application_name"] = serviceName
	pgxConfig.MaxConns = cfg.PoolConfig.MaxConnections
	pgxConfig.MinConns = cfg.PoolConfig.MinConnections

	pool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	if err := otelpgx.RecordStats(pool); err != nil {
		return nil, fmt.Errorf("unable to record database stats: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
