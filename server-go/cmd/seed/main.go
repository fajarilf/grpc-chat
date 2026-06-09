// Command seed populates the database with sample users and rooms for local
// development. Run migrations first, then:
//
//	go run ./cmd/seed
//
// It reads the database URL from DB_URL (loaded from .env if present) and is
// destructive: it truncates users, rooms and room_users before inserting.
package main

import (
	"context"
	"log"
	"os"

	"github.com/fajarilf/grpc-chat-server/repositories"
	"github.com/fajarilf/grpc-chat-server/seeders"
	"github.com/joho/godotenv"
)

func main() {
	// Best-effort: load .env if present, otherwise rely on the real env.
	_ = godotenv.Load()

	if os.Getenv("DB_URL") == "" {
		log.Fatalf("DB_URL is not set")
	}

	ctx := context.Background()

	pool, err := repositories.Connect(ctx)
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	defer pool.Close()

	if err := seeders.Seed(ctx, pool); err != nil {
		log.Fatalf("seed: %v", err)
	}

	log.Println("seed: ok")
}
