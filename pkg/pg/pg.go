package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgDb struct {
	*pgxpool.Pool
}

func Deal(dsn string) (*PgDb, error) {
	ctx, cancel := context.WithCancel(context.Background())

	if dsn == "" {
		return nil, fmt.Errorf("dsn empty")
	}

	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed create pg: %w", err)
	}

	if err := dbpool.Ping(ctx); err != nil {
		cancel()
		return nil, fmt.Errorf("failed ping postgres: %w", err)
	}

	return &PgDb{
		dbpool,
	}, nil

}