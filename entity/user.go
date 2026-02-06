package entity

import "time"

type User struct {
	ID        int
	FullName  string
	RoleID    int
	RoleName  string
	CreatedAt time.Time
}

type LoginUser struct {
	UserId   int
	Fullname string
	Password string
	RoleID   int
}

type RegisterUser struct {
	UserId   int
	Fullname string
	Password string
	RoleID   int
}
