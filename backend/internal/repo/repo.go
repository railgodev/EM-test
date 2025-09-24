package repo

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	conn *pgxpool.Pool
	log  *slog.Logger
}

func New(conn *pgxpool.Pool, log *slog.Logger) *Repo {
	return &Repo{conn: conn, log: log}
}
