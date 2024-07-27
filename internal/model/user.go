package model

import "errors"

var (
	ErrInvalidPassword  = errors.New("invalid password")
	ErrUsernameNotFound = errors.New("username not found")
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) IsPasswordMatch(password string) bool {
	return u.Password == password
}

type IUserRepository interface {
	Create(user User) error
	Login(username string) (User, error)
	FindByUsername(username string) (User, error)
}

type IUserUsecase interface {
	Create(user User) error
	Login(username string, password string) (User, error)
	FindByUsername(username string) (User, error)
}
