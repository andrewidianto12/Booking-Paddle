package handler

import (
	"database/sql"
	"errors"

	"github.com/andrewidianto/Paddle-Booking/entity"
)

func RegisterUser(db *sql.DB, fullname, password string, roleID int) (*entity.RegisterUser, error) {
	var user entity.RegisterUser

	result, err := db.Exec(`
		INSERT INTO users (full_name, password, role_id)
		VALUES (?, ?, ?)
	`, fullname, password, roleID)
	if err != nil {
		return nil, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.UserId = int(lastID)
	user.Fullname = fullname
	user.Password = password
	user.RoleID = roleID

	return &user, nil
}

func LoginUser(db *sql.DB, fullName, password string) (*entity.LoginUser, error) {
	var user entity.LoginUser

	err := db.QueryRow(`
		SELECT user_id, full_name, password, role_id
		FROM users
		WHERE full_name = ?
		LIMIT 1
	`, fullName).Scan(
		&user.UserId,
		&user.Fullname,
		&user.Password,
		&user.RoleID,
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
