package dbx

import (
	"context"
	"database/sql"
	"time"

	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"

	// Driver import
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	// Driver is the PostgreSQL driver used by Hubris.
	Driver = "pgx"
	// MaxAttempts are the number of retries that can be made.
	MaxAttempts = 20
	// MaxTimeout is the number of seconds to wait between connection attempts.
	MaxTimeout = time.Second * 10
)

// Open opens a connection to a database.
func Open(ctx context.Context, dsn string) (*sql.DB, error) {
	var (
		db  *sql.DB
		err error
	)

	connected := false
	for i := 0; i < MaxAttempts; i++ {
		db, err = sql.Open(Driver, dsn)
		if err != nil {
			ctxlog.Error(ctx, "error opening database connection", zap.Error(err))
			continue
		}

		if err := db.PingContext(ctx); err != nil {
			ctxlog.Error(ctx, "error pinging database", zap.Error(err))
			continue
		}

		connected = true
	}

	if !connected {
		return nil, err
	}

	return db, nil
}
