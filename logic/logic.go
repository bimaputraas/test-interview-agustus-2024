package logic

import (
	"coding-interview-agustus-1/pkg"
	user "coding-interview-agustus-1/proto"
	"coding-interview-agustus-1/repository"
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type (
	Logic struct {
		repo repository.Repository
	}
	AuthParams struct {
		Token        string
		XLinkService string
	}
)

func NewLogic(repo repository.Repository) *Logic {
	return &Logic{repo: repo}
}

func (logic *Logic) GetAllUsers() ([]*user.User, error) {
	users, err := logic.repo.GetAllUser()
	if err != nil {
		return nil, err
	}

	return users, nil
}
func (logic *Logic) CreateUser(data *user.User) error {
	err := logic.repo.CreateUser(data)
	if err != nil {
		return err
	}

	return nil
}
func (logic *Logic) UpdateUserById(id int, data *user.User) error {
	err := logic.repo.UpdateUserById(id, data)
	if err != nil {
		return err
	}

	return nil
}
func (logic *Logic) DeleteUserById(id int) error {
	err := logic.repo.DeleteUserById(id)
	if err != nil {
		return err
	}

	return nil
}

func (logic *Logic) Login(email string, password string) (string, error) {
	if email == "" || password == "" {
		return "", errors.New("email or password are required")
	}

	user, err := logic.repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if user.Password != password {
		return "", errors.New("wrong email or password")
	}

	secret := []byte("example")
	token, err := pkg.GenerateJWT(jwt.MapClaims{"user_id": user.Id}, secret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (logic *Logic) Auth(auth AuthParams) (user.User, error) {
	if auth.Token == "" {
		return user.User{}, errors.New("no token")
	}

	if auth.XLinkService == "" {
		return user.User{}, errors.New("unauthorized user")
	}

	if auth.XLinkService != "be" {
		return user.User{}, errors.New("unauthorized user")
	}

	jwt, err := pkg.ParseJWT(auth.Token, []byte("example"))
	if err != nil {
		return user.User{}, err
	}

	uid, ok := jwt["user_id"].(float64)
	if !ok {
		return user.User{}, errors.New("invalid token")
	}

	u, err := logic.repo.GetUserById(int(uid))
	if err != nil {
	}
	if err == sql.ErrNoRows {
		return user.User{}, errors.New("unauthorized user")
	}

	if err != nil {
		return user.User{}, nil
	}

	return *u, nil
}

func (logic *Logic) AuthRead(u user.User) (bool, error) {
	role, err := logic.repo.GetUserRoleById(int(u.RoleId))
	if err == sql.ErrNoRows {
		return false, errors.New("unauthorized user")
	}
	if err != nil {
		return false, err
	}

	return role.RRead == 1, nil

}
func (logic *Logic) AuthCreate(u user.User) (bool, error) {
	role, err := logic.repo.GetUserRoleById(int(u.RoleId))
	if err == sql.ErrNoRows {
		return false, errors.New("unauthorized user")
	}
	if err != nil {
		return false, err
	}

	return role.RCreate == 1, nil
}
func (logic *Logic) AuthDelete(u user.User) (bool, error) {
	role, err := logic.repo.GetUserRoleById(int(u.RoleId))
	if err == sql.ErrNoRows {
		return false, errors.New("unauthorized user")
	}
	if err != nil {
		return false, err
	}

	return role.RDelete == 1, nil
}
func (logic *Logic) AuthUpdate(u user.User) (bool, error) {
	role, err := logic.repo.GetUserRoleById(int(u.RoleId))
	if err == sql.ErrNoRows {
		return false, errors.New("unauthorized user")
	}
	if err != nil {
		return false, err
	}

	return role.RUpdate == 1, nil
}
