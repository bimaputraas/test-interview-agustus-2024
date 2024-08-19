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

func (logic *Logic) Authentication(token string) (*user.User, error) {
	if token == "" {
		return &user.User{}, errors.New("no token")
	}

	jwt, err := pkg.ParseJWT(token, []byte("example"))
	if err != nil {
		return &user.User{}, err
	}

	uid, ok := jwt["user_id"].(float64)
	if !ok {
		return &user.User{}, errors.New("invalid token")
	}

	u, err := logic.repo.GetUserById(int(uid))
	if err != nil {
	}
	if err == sql.ErrNoRows {
		return &user.User{}, errors.New("unauthorized user")
	}

	if err != nil {
		return &user.User{}, nil
	}

	return u, nil
}

type (
	ActionType int
)

const (
	Read   ActionType = 1
	Update ActionType = 2
	Delete ActionType = 3
	Create ActionType = 4
)

func (logic *Logic) Authorize(u *user.User, xLinkService string, act ActionType) error {
	if xLinkService != "be" {
		return errors.New("unauthorized user")
	}

	role, err := logic.repo.GetUserRoleById(int(u.RoleId))
	if err == sql.ErrNoRows {
		return errors.New("unauthorized user")
	}
	if err != nil {
		return err
	}

	if role.RRead != 1 {
		return errors.New("unauthorized user")
	}

	var isValid bool
	switch act {
	case Read:
		isValid = role.RRead == 1
	case Update:
		isValid = role.RUpdate == 1
	case Delete:
		isValid = role.RDelete == 1
	case Create:
		isValid = role.RCreate == 1
	default:
		return errors.New("invalid action")
	}

	if !isValid {
		return errors.New("unauthorized user")
	}

	return nil

}
