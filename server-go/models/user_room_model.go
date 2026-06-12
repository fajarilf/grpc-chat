package models

import "time"

type UserRoom struct {
	RoomId   int       `db:"room_id"`
	UserId   int       `db:"user_id"`
	JoinedAt time.Time `db:"joined_at"`
}

type UserRoomCreateRequest struct {
	RoomId int
	UserId int
}
