package repositories

import (
	"context"
	"log"

	pb "github.com/fajarilf/grpc-chat-proto/proto"
	"github.com/fajarilf/grpc-chat-server/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RoomRepository struct {
	pool *pgxpool.Pool
}

func NewRoomRepository(pool *pgxpool.Pool) *RoomRepository {
	return &RoomRepository{
		pool: pool,
	}
}

func (r *RoomRepository) Get(ctx context.Context, param *pb.RoomListRequest) ([]*models.Room, error) {
	var (
		rows pgx.Rows
		err  error
	)

	switch {
	case param.Cursor == 0:
		log.Print("case 1", param.Cursor, param.Size)
		rows, err = r.pool.Query(ctx,
			`SELECT id, name, description, background, background_type, created_at, updated_at 
			 FROM rooms
			 ORDER BY id DESC 
			 LIMIT $1`,
			param.Size,
		)
	case param.Forward:
		log.Print("case 2", param.Cursor, param.Size)
		rows, err = r.pool.Query(ctx,
			`SELECT id, name, description, background, background_type, created_at, updated_at 
			 FROM rooms
			 WHERE id > $1 
			 ORDER BY id DESC 
			 LIMIT $2`,
			param.Cursor, param.Size,
		)
	default:
		log.Print("case 3", param.Cursor, param.Size)
		rows, err = r.pool.Query(ctx,
			`SELECT id, name, description, background, background_type, created_at, updated_at 
			 FROM rooms
			 WHERE id < $1 
			 ORDER BY id ASC 
			 LIMIT $2`,
			param.Cursor, param.Size,
		)
	}
	if err != nil {
		return nil, err
	}

	rooms, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.Room])
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r *RoomRepository) GetById(ctx context.Context, id int) (*models.Room, error) {
	rows, _ := r.pool.Query(ctx,
		`SELECT id, name, description, background, background_type, created_at, updated_at 
		FROM rooms WHERE id = $1`, id,
	)

	room, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[models.Room])
	if err != nil {
		return nil, err
	}

	return room, nil
}
