package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/OzkrOssa/idktd/pkg/config"
	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

/**
 * DB is a wrapper for PostgreSQL database connection
 * that uses pgxpool as database driver.
 * It also holds a reference to squirrel.StatementBuilderType
 * which is used to build SQL queries that compatible with PostgreSQL syntax
 */
type DB struct {
	*pgxpool.Pool
	*squirrel.StatementBuilderType
	url string
}

// New creates a new PostgreSQL database instance
func New(ctx context.Context, config *config.DB) (*DB, error) {
	url := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.Connection,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	cfg.ConnConfig.Tracer = otelpgx.NewTracer()

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return &DB{}, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return &DB{}, err
	}

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	db := &DB{
		pool,
		&psql,
		url,
	}

	return db, nil
}

// ErrorCode returns the error code of the given error
func (db *DB) ErrorCode(err error) string {
	var pgErr *pgconn.PgError
	errors.As(err, &pgErr)
	return pgErr.Code
}

// Close closes the database connection
func (db *DB) Close() {
	db.Pool.Close()
}
