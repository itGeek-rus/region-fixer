package database

import (
	"context"
	"fmt"
	"region-fixer/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConnection(ctx context.Context, cfg config.DBConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	if cfg.SSLMode != "" {
		dsn = dsn + fmt.Sprintf("?sslmode=%s", cfg.SSLMode)
	}

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to create db pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return pool, nil
}
