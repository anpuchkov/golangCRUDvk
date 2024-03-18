package psql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

func InitPostgresConnection(ctx context.Context, cfg pgconn.Config) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s search_path=filmoteka", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to the database")
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to ping the database")
	}

	return pool, nil
}
