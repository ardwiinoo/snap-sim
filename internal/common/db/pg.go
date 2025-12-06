package db

import (
	"context"
	"fmt"

	"github.com/ardwiinoo/snap-sim/internal/common/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPgPool(cfg config.DBConfig) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SSLMode,
	)

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
