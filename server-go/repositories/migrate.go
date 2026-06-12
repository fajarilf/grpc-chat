package repositories

import (
	"errors"
	"strings"

	"github.com/fajarilf/grpc-chat-server/migrations"
	"github.com/fajarilf/grpc-chat-server/utils"
	"github.com/golang-migrate/migrate/v4"

	// Registers the "pgx5" database URL scheme for golang-migrate.
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

// NewMigrator builds a *migrate.Migrate backed by the embedded SQL files and
// the given Postgres URL. Callers own the returned instance and must Close it.
//
// dbURL is the same DB_URL used by Connect (e.g. postgres://user:pass@host/db);
// the scheme is rewritten to pgx5:// so golang-migrate routes it through the
// pgx/v5 driver instead of pulling in a separate Postgres driver.
func NewMigrator(dbURL string) (*migrate.Migrate, error) {
	src, err := iofs.New(migrations.Files, ".")
	if err != nil {
		return nil, utils.WrapError("load embedded migrations", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", src, toPgxURL(dbURL))
	if err != nil {
		return nil, utils.WrapError("init migrator", err)
	}
	return m, nil
}

// Migrate applies all pending up migrations. It is safe to call on startup:
// ErrNoChange (nothing to apply) is treated as success.
func Migrate(dbURL string) error {
	m, err := NewMigrator(dbURL)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return utils.WrapError("apply migrations", err)
	}
	return nil
}

// toPgxURL swaps a postgres://postgresql:// scheme for the pgx5:// scheme that
// golang-migrate's pgx/v5 driver registers.
func toPgxURL(dbURL string) string {
	for _, prefix := range []string{"postgresql://", "postgres://"} {
		if strings.HasPrefix(dbURL, prefix) {
			return "pgx5://" + strings.TrimPrefix(dbURL, prefix)
		}
	}
	return dbURL
}
