package repositories

import (
	"context"

	"github.com/fajarilf/grpc-chat-server/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}

func (r *UserRepository) Get(ctx context.Context) ([]*models.User, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, name, username, password, created_at, updated_at
		 FROM users`,
	)
	if err != nil {
		return nil, err
	}

	users, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.User])
	if err != nil {
		return nil, err
	}

	return users, nil
}
