package repositories

import (
	"context"
	"slices"

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

// Create inserts a room using the given executor, which may be the pool or a
// transaction so the caller can compose it with the membership insert atomically.
func (r *RoomRepository) Create(ctx context.Context, db DBTX, param *pb.RoomCreateRequest) (*models.Room, error) {
	rows, err := db.Query(ctx,
		`INSERT INTO rooms (name, description, background, background_type)
		 VALUES ($1, $2, $3, $4)
		 RETURNING *`,
		param.Name, param.Description, param.Background, param.BackgroundType,
	)
	if err != nil {
		return nil, err
	}

	room, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[models.Room])
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (r *RoomRepository) Get(ctx context.Context, param *pb.RoomListRequest) (*models.RoomCursor, error) {
	var (
		size     int
		rows     pgx.Rows
		err      error
		backward bool
	)

	size = int(param.Size)

	switch {
	case param.Cursor == 0:
		rows, err = r.pool.Query(ctx,
			`SELECT id, name, description, background, background_type, created_at, updated_at 
			 FROM rooms
			 ORDER BY id DESC 
			 LIMIT $1`,
			param.Size+1,
		)
	case param.Forward:
		rows, err = r.pool.Query(ctx,
			`SELECT id, name, description, background, background_type, created_at, updated_at 
			 FROM rooms
			 WHERE id < $1 
			 ORDER BY id DESC 
			 LIMIT $2`,
			param.Cursor, param.Size+1,
		)
	default:
		backward = true
		rows, err = r.pool.Query(ctx,
			`SELECT id, name, description, background, background_type, created_at, updated_at 
			 FROM rooms
			 WHERE id > $1 
			 ORDER BY id ASC 
			 LIMIT $2`,
			param.Cursor, param.Size+1,
		)
	}
	if err != nil {
		return nil, err
	}

	rooms, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.Room])
	if err != nil {
		return nil, err
	}

	hasMore := len(rooms) > size
	if hasMore {
		rooms = rooms[:size]
	}

	if backward {
		slices.Reverse(rooms)
	}

	result := &models.RoomCursor{
		Rooms:   rooms,
		HasMore: hasMore,
	}

	if len(rooms) > 0 {
		result.NextCursor = rooms[len(rooms)-1].Id
		result.PrevCursor = rooms[0].Id
	}

	return result, nil
}

// MembersByRoomIDs returns the members of each room, keyed by room id, in a
// single query — avoids an N+1 of one membership lookup per room.
func (r *RoomRepository) MembersByRoomIDs(ctx context.Context, roomIDs []int) (map[int][]*models.User, error) {
	out := make(map[int][]*models.User)
	if len(roomIDs) == 0 {
		return out, nil
	}

	rows, err := r.pool.Query(ctx,
		`SELECT ru.room_id, u.id, u.name, u.username
		 FROM room_users ru
		 JOIN users u ON u.id = ru.user_id
		 WHERE ru.room_id = ANY($1)
		 ORDER BY ru.room_id, u.id`,
		roomIDs, // pgx encodes []int as a Postgres integer[] for ANY()
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			roomID int
			user   models.User
		)
		if err := rows.Scan(&roomID, &user.Id, &user.Name, &user.Username); err != nil {
			return nil, err
		}
		member := user // copy before taking its address inside the loop
		out[roomID] = append(out[roomID], &member)
	}

	return out, rows.Err()
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
