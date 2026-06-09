// Command migrate runs database migrations from the CLI, separately from the
// gRPC server. Run it in CI/CD or by hand:
//
//	go run ./cmd/migrate up        # apply all pending migrations
//	go run ./cmd/migrate down      # roll back the most recent migration
//	go run ./cmd/migrate version   # print current version and dirty state
//	go run ./cmd/migrate force 2   # mark version 2 as applied (fix dirty state)
//
// It reads the database URL from DB_URL (loaded from .env if present).
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/fajarilf/grpc-chat-server/repositories"
	"github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: migrate <up|down|version|force> [n]")
	}

	// Best-effort: load .env if it exists, otherwise rely on the real env.
	_ = godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatalf("DB_URL is not set")
	}

	m, err := repositories.NewMigrator(dbURL)
	if err != nil {
		log.Fatalf("migrator: %v", err)
	}
	defer m.Close()

	cmd := os.Args[1]
	switch cmd {
	case "up":
		err = m.Up()
	case "down":
		// Roll back a single step rather than everything, which is the safer
		// default when invoked by hand.
		err = m.Steps(-1)
	case "version":
		printVersion(m)
		return
	case "force":
		if len(os.Args) < 3 {
			log.Fatalf("usage: migrate force <version>")
		}
		v, convErr := strconv.Atoi(os.Args[2])
		if convErr != nil {
			log.Fatalf("invalid version %q: %v", os.Args[2], convErr)
		}
		err = m.Force(v)
	default:
		log.Fatalf("unknown command %q (want up|down|version|force)", cmd)
	}

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("%s: %v", cmd, err)
	}

	fmt.Printf("migrate %s: ok\n", cmd)
	printVersion(m)
}

func printVersion(m *migrate.Migrate) {
	version, dirty, err := m.Version()
	if errors.Is(err, migrate.ErrNilVersion) {
		fmt.Println("version: none (no migrations applied)")
		return
	}
	if err != nil {
		log.Fatalf("version: %v", err)
	}
	fmt.Printf("version: %d (dirty=%t)\n", version, dirty)
}
