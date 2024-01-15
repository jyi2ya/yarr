package storage

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	db *sql.DB
}

func New(path string) (*Storage, error) {
	db, err := sql.Open("pgx", path)
	if err != nil {
		return nil, err
	}

	if err = migrate(db); err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}
