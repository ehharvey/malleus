package pg

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDbPool(connectionString string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, connectionString)

	if err != nil {
		return nil, err
	} else {
		return pool, err
	}
}
