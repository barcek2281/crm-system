package dal

import (
	"crmsystem/internal/model"
	"database/sql"

	_ "github.com/lib/pq"
)

type User struct {
	db *sql.DB
}

func (u *User) Register(user model.RegisterRequest) (string, error) {
	query := `INSERT INTO "user" (first_name, last_name, password_hash) VALUES ($1, $2, $3, $4, $5) RETURNING user_id`
	id := ""
	if err := u.db.QueryRow(query, user.FirstName, user.LastName, user.Password).Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

func (u *User) Exist(user model.LoginRequest) error {
	q := `SELECT user_id FROM "user" WHERE login = $1`
	return u.db.QueryRow(q, user.Login).Err()
}

func (u *User) Get(login string) (model.User, error) {
	q := `SELECT user_id, first_name, last_name, password_hash FROM "user" WHERE login = $1`
	user := model.User{}
	if err := u.db.QueryRow(q, login).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,

		&user.Password); err != nil {
		return model.User{}, err
	}

	return user, nil
}
