package storage

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

func GetConn(logger *slog.Logger, connStr string) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, errors.Wrap(err, "config parse")
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, errors.Wrap(err, "newWithConfig")
	}
	if err = pool.Ping(context.Background()); err != nil {
		return nil, errors.Wrap(err, "ping")
	}
	logger.Info("Successfully connected to database")
	return pool, nil
}
