// Package seeders fills the database with deterministic sample data for local
// development. It assumes the schema already exists (run cmd/migrate first).
package seeders

import (
	"context"
	"fmt"
	"math/rand/v2"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

// defaultPassword is the plaintext every seeded user shares; it is bcrypt
// hashed before insert so the rows look like real records.
const defaultPassword = "password"

// roomCount is how many rooms to seed. Bump it if you need a longer list.
const roomCount = 12

type seedUser struct {
	name     string
	username string
}

var seedUsers = []seedUser{
	{"Aiko Tanaka", "aiko"},
	{"Budi Santoso", "budi"},
	{"Chen Wei", "chenw"},
	{"Dewi Lestari", "dewi"},
	{"Eko Prabowo", "eko"},
	{"Fajar Ilham", "fajar"},
}

// background_type values mirror the BgTypes proto enum (0 = COLOR, 1 = IMAGE).
const (
	bgColor int16 = 0
	bgImage int16 = 1
)

var colors = []string{
	"#475569", "#0284c7", "#059669", "#d97706",
	"#e11d48", "#7c3aed", "#14110d", "#0d9488",
}

var imgs = []string{
	"https://picsum.photos/seed/mountains/640/480",
	"https://picsum.photos/seed/city/640/480",
	"https://picsum.photos/seed/ocean/640/480",
	"https://picsum.photos/seed/forest/640/480",
	"https://picsum.photos/seed/desert/640/480",
	"https://picsum.photos/seed/aurora/640/480",
}

// Seed truncates and repopulates users, rooms and room_users in a single
// transaction so a partial failure leaves the database untouched. Truncating
// with RESTART IDENTITY keeps ids deterministic (1..N) across runs.
func Seed(ctx context.Context, pool *pgxpool.Pool) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(defaultPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	// Rolled back automatically unless we reach Commit; a no-op after commit.
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx,
		`TRUNCATE room_users, rooms, users RESTART IDENTITY CASCADE`,
	); err != nil {
		return fmt.Errorf("truncate tables: %w", err)
	}

	userIDs := make([]int, 0, len(seedUsers))
	for _, u := range seedUsers {
		var id int
		if err := tx.QueryRow(ctx,
			`INSERT INTO users (name, username, password)
			 VALUES ($1, $2, $3)
			 RETURNING id`,
			u.name, u.username, string(hash),
		).Scan(&id); err != nil {
			return fmt.Errorf("insert user %q: %w", u.username, err)
		}
		userIDs = append(userIDs, id)
	}

	for i := range roomCount {
		bgVal := imgs[rand.IntN(len(imgs))]
		bgType := bgImage
		if i%2 == 0 { // even rooms get a color, odd rooms an image
			bgVal = colors[rand.IntN(len(colors))]
			bgType = bgColor
		}

		var roomID int
		if err := tx.QueryRow(ctx,
			`INSERT INTO rooms (name, description, background, background_type)
			 VALUES ($1, $2, $3, $4)
			 RETURNING id`,
			fmt.Sprintf("Room %d", i+1),
			fmt.Sprintf("Seeded room %d", i+1),
			bgVal, bgType,
		).Scan(&roomID); err != nil {
			return fmt.Errorf("insert room %d: %w", i+1, err)
		}

		for _, uid := range pickMembers(userIDs, 3) {
			if _, err := tx.Exec(ctx,
				`INSERT INTO room_users (room_id, user_id) VALUES ($1, $2)`,
				roomID, uid,
			); err != nil {
				return fmt.Errorf("insert membership room=%d user=%d: %w", roomID, uid, err)
			}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit: %w", err)
	}
	return nil
}

// pickMembers returns up to n distinct ids drawn at random from the pool.
func pickMembers(ids []int, n int) []int {
	n = min(n, len(ids))
	perm := rand.Perm(len(ids))
	out := make([]int, n)
	for i := range n {
		out[i] = ids[perm[i]]
	}
	return out
}
