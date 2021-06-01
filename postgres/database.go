package postgres

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	sqldb *sql.DB
}

func New(connString string) (*Database, error) {
	sqldb, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	return &Database{
		sqldb: sqldb,
	}, nil
}

func (db *Database) Get(ctx context.Context) (*sql.DB, error) {
	if err := db.sqldb.PingContext(ctx); err != nil {
		return nil, err
	}
	return db.sqldb, nil
}

func (db *Database) Close() error {
	return db.sqldb.Close()
}
