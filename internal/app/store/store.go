package store

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
)

// DB ...
type DB struct {
	config Config
	db     *sql.DB
}

// New ...
func New(config Config) DB {
	return DB{
		config: config,
	}
}

// Open ...
func (d DB) Open() error {
	db, err := sql.Open("postgres", d.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	d.db = db
	d.config.Logger.Info("Database opened!")
	return nil
}

// Close ...
func (d DB) Close() {
	d.db.Close()
}
