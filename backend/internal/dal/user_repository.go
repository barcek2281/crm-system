package dal

import (
	"crmsystem/internal/model"
	"database/sql"

	_ "github.com/lib/pq"
)

type User struct {
	db *sql.DB
}

func (u *User) Register(user model.RegisterRequest) error {
	query := `INSERT INTO "user" (email, phone_number, password_hash) VALUES ($1, $2, $3)`
	row := u.db.QueryRow(query, user.Email, user.PhoneNumber, user.Password)

	if row.Err() != nil {
		return row.Err()
	}
	return nil
}
