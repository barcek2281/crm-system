package dal

import (
	"crmsystem/internal/config"
	"database/sql"
	"fmt"
	"log/slog"
	_ "github.com/lib/pq"
)

const (
	POSTGRES = "postgres"
)

type Store struct {
	db   *sql.DB
	user *User
}

func NewStore(cnf config.Config) (*Store, error) {
	db, err := sql.Open(POSTGRES, fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable",
		POSTGRES,
		cnf.DB.DBuser,
		cnf.DB.DBpassword,
		cnf.DB.DBhost,
		cnf.DB.DBport,
		cnf.DB.DBname,
	))
	if err != nil {
		slog.Error("error to open db", "error", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		slog.Error("error to ping db", "error", err)
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}

func (s *Store) User() *User {
	if s.user == nil {
		s.user = &User{
			db: s.db,
		}
	}
	return s.user
}
