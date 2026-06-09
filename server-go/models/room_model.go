package models

import (
	"time"

	pb "github.com/fajarilf/grpc-chat-proto/proto"
)

type Room struct {
	Id             int       `db:"id"`
	Name           string    `db:"name"`
	Description    string    `db:"description"`
	Background     string    `db:"background"`
	BackgroundType int16     `db:"background_type"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

type RoomCursor struct {
	Rooms      []*Room
	NextCursor int
	PrevCursor int
	HasMore    bool
}

func GetStringBgType(num int) string {
	if name, ok := pb.BgTypes_name[int32(num)]; ok {
		return name
	}

	return pb.BgTypes_name[int32(pb.BgTypes_COLOR)]
}

func ToRoomResponse(entity *Room) *pb.RoomResponse {
	return &pb.RoomResponse{
		Id:             int32(entity.Id),
		Name:           entity.Name,
		Description:    entity.Description,
		Background:     entity.Background,
		BackgroundType: GetStringBgType(int(entity.BackgroundType)),
	}
}

func ToRoomResponseWithUser(entity *Room, members []*User) *pb.RoomResponseWithUser {
	users := make([]*pb.UserResponse, 0, len(members))
	for _, m := range members {
		users = append(users, ToUserResponse(m))
	}

	return &pb.RoomResponseWithUser{
		Id:             int32(entity.Id),
		Name:           entity.Name,
		Description:    entity.Description,
		Background:     entity.Background,
		BackgroundType: GetStringBgType(int(entity.BackgroundType)),
		Users:          users,
	}
}
