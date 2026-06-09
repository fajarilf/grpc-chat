package models

import (
	"time"

	pb "github.com/fajarilf/grpc-chat-proto/proto"
)

type User struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func ToUserResponse(entity *User) *pb.UserResponse {
	return &pb.UserResponse{
		Id:       int32(entity.Id),
		Name:     entity.Name,
		Username: entity.Username,
	}
}

func ToUserLoginResponse(entity *User) *pb.UserLoginResponse {
	return &pb.UserLoginResponse{
		Username: entity.Username,
	}
}
