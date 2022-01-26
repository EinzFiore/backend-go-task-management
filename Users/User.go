package users

import "time"

type User struct {
	Id          int
	Fullname    string
	Description string
	Avatar      string
	Email       string
	Password    string
	RoleId      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
