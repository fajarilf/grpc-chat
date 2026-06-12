package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/fajarilf/grpc-chat-proto/proto"
	server "github.com/fajarilf/grpc-chat-server/grpc-server"
	"github.com/fajarilf/grpc-chat-server/repositories"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	_ = godotenv.Load()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := repositories.Connect(ctx)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	defer pool.Close()
	log.Println("connected to PostgreSQL")

	grpcServer := grpc.NewServer()

	// repository
	roomRepo := repositories.NewRoomRepository(pool)
	userRepo := repositories.NewUserRepository(pool)
	userRoomRepo := repositories.NewUserRoomRepository(pool)

	// register grpc server
	roomServer := server.NewRoomServer(roomRepo, userRoomRepo)
	pb.RegisterRoomServiceServer(grpcServer, roomServer)

	userServer := server.NewUserServer(userRepo)
	pb.RegisterUserServiceServer(grpcServer, userServer)

	log.Println("Chat server started on :50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
