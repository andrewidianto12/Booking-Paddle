package entity

type User struct {
	UserId   int
	Username string
	Email    string
	Password string
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
