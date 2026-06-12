package repositories

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/fajarilf/grpc-chat-server/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"%s", os.Getenv("DB_URL"),
	)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, utils.WrapError("Parse config error", err)
	}

	config.MaxConns = 25
	config.MinConns = 5
	config.MaxConnLifetime = 5 * time.Minute
	config.MaxConnIdleTime = 1 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, utils.WrapError("Create pool error", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, utils.WrapError("Ping database error", err)
	}

	return pool, nil
}
