// Package migrations holds the versioned SQL migration files and exposes them
// as an embedded filesystem so they ship inside the compiled binary (no need
// to copy a migrations/ directory alongside the deployed server).
package migrations

import "embed"

// Files contains every *.sql migration in this directory. golang-migrate reads
// it through the iofs source driver.
//
//go:embed *.sql
var Files embed.FS
