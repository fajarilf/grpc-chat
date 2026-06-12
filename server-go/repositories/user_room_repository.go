package repositories

import (
	"context"
	"fmt"

	"github.com/fajarilf/grpc-chat-server/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRoomRepository struct {
	pool *pgxpool.Pool
}

func NewUserRoomRepository(pool *pgxpool.Pool) *UserRoomRepository {
	return &UserRoomRepository{
		pool: pool,
	}
}

// CreateBatch inserts one membership row per userID for a single room in one
// round trip. Postgres expands the int[] via unnest into one row each, pairing
// every user_id with the constant room_id. ON CONFLICT DO NOTHING makes it
// idempotent so re-adding an existing member is a no-op instead of a PK error.
func (r *UserRoomRepository) CreateBatch(ctx context.Context, roomID int, userIDs []int) error {
	if len(userIDs) == 0 {
		return nil
	}

	_, err := r.pool.Exec(ctx,
		`INSERT INTO room_users (room_id, user_id)
		 SELECT $1, unnest($2::int[])
		 ON CONFLICT DO NOTHING`,
		roomID, userIDs,
	)
	if err != nil {
		return fmt.Errorf("UserRoomRepo CreateBatch Error: %v", err)
	}

	return nil
}

func (r *UserRoomRepository) GetList(ctx context.Context) ([]*models.UserRoom, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT user_id, room_id FROM room_users`,
	)
	if err != nil {
		return nil, fmt.Errorf("UserRoomRepo Error: %v", rows)
	}

	userRoom, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.UserRoom])
	if err != nil {
		return nil, fmt.Errorf("UserRoomRepo Error: %v", rows)
	}

	return userRoom, err
}
