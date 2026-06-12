package services

import (
	"context"
	"errors"

	pb "github.com/fajarilf/grpc-chat-proto/proto"
	"github.com/fajarilf/grpc-chat-server/models"
	"github.com/fajarilf/grpc-chat-server/repositories"
	"github.com/fajarilf/grpc-chat-server/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
)

type RoomService struct {
	pool         *pgxpool.Pool
	repo         *repositories.RoomRepository
	userRoomRepo *repositories.UserRoomRepository
}

func NewRoomService(pool *pgxpool.Pool, repo *repositories.RoomRepository, userRoomRepo *repositories.UserRoomRepository) *RoomService {
	return &RoomService{
		pool:         pool,
		repo:         repo,
		userRoomRepo: userRoomRepo,
	}
}

func (rs *RoomService) CreateRoom(ctx context.Context, param *pb.RoomCreateRequest) (*pb.RoomResponseWithUser, error) {
	userIDs := make([]int, len(param.UserIds))
	for i, id := range param.UserIds {
		userIDs[i] = int(id)
	}

	// Room insert and membership insert must commit together, so run both on a
	// single transaction. The deferred Rollback is a no-op once Commit succeeds.
	tx, err := rs.pool.Begin(ctx)
	if err != nil {
		return nil, utils.WrapGRPCError(codes.Internal, "Create Room Error", err)
	}
	defer tx.Rollback(ctx)

	result, err := rs.repo.Create(ctx, tx, param)
	if err != nil {
		return nil, utils.WrapGRPCError(codes.Internal, "Create Room Error", err)
	}

	if err := rs.userRoomRepo.CreateBatch(ctx, tx, result.Id, userIDs); err != nil {
		return nil, utils.WrapGRPCError(codes.Internal, "Create Room Error", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, utils.WrapGRPCError(codes.Internal, "Create Room Error", err)
	}

	membersByRoom, err := rs.repo.MembersByRoomIDs(ctx, []int{result.Id})
	if err != nil {
		return nil, utils.WrapGRPCError(codes.Internal, "Create Room Error", err)
	}

	return models.ToRoomResponseWithUser(result, membersByRoom[result.Id]), nil
}

func (rs *RoomService) GetRoomById(ctx context.Context, param *pb.RoomId) (*pb.RoomResponseWithUser, error) {
	result, err := rs.repo.GetById(ctx, int(param.Id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, utils.WrapGRPCError(codes.NotFound, "Get Room by Id Error", err)
		}
		return nil, utils.WrapGRPCError(codes.Internal, "Get Room by Id Error", err)
	}

	membersByRoom, err := rs.repo.MembersByRoomIDs(ctx, []int{result.Id})
	if err != nil {
		return nil, utils.WrapGRPCError(codes.Internal, "Get Room by Id Error", err)
	}

	return models.ToRoomResponseWithUser(result, membersByRoom[result.Id]), nil
}

func (rs *RoomService) GetListRoom(ctx context.Context, param *pb.RoomListRequest) (*pb.RoomListResponse, error) {
	result, err := rs.repo.Get(ctx, param)
	if err != nil {
		return nil, utils.WrapGRPCError(codes.Internal, "Get List Room Error", err)
	}

	ids := make([]int, len(result.Rooms))
	for i, val := range result.Rooms {
		ids[i] = val.Id
	}

	membersByRoom, _ := rs.repo.MembersByRoomIDs(ctx, ids)

	rooms := make([]*pb.RoomResponseWithUser, 0, len(result.Rooms))
	for _, val := range result.Rooms {
		room := models.ToRoomResponseWithUser(val, membersByRoom[val.Id])
		rooms = append(rooms, room)
	}

	return &pb.RoomListResponse{
		Rooms:      rooms,
		NextCursor: int32(result.NextCursor),
		PrevCursor: int32(result.PrevCursor),
		HasMore:    result.HasMore,
	}, nil
}
