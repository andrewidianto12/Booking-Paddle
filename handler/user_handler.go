package handler

import (
	"database/sql"
	"errors"

	"github.com/andrewidianto/Paddle-Booking/entity"
)

func LoginUser(db *sql.DB, fullName, password string) (*entity.LoginUser, error) {
	var user entity.LoginUser

	err := db.QueryRow(`
		SELECT full_name, password
		FROM users
		WHERE full_name = ?
		LIMIT 1
	`, fullName).Scan(
		&user.Fullname,
		&user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}
