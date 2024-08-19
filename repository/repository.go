package repository

import (
	user "coding-interview-agustus-1/proto"
	"database/sql"
)

type Repository interface {
	GetUserByEmail(email string) (*user.User, error)
	GetUserById(uid int) (*user.User, error)
	GetUserRoleById(id int) (*user.Role, error)
	GetAllUser() ([]*user.User, error)
	CreateUser(*user.User) error
	UpdateUserById(id int, data *user.User) error
	DeleteUserById(id int) error
}

type (
	repository struct {
		sql.DB
	}
)
