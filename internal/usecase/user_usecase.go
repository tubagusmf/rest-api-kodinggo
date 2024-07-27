package usecase

import (
	"golang-rest-api-articles/internal/model"

	"github.com/sirupsen/logrus"
)

type UserUsecase struct {
	userRepo model.IUserRepository
}

func NewUserUsecase(
	userRepo model.IUserRepository,
) model.IUserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (u *UserUsecase) Create(user model.User) error {
	log := logrus.WithFields(logrus.Fields{
		"user": user,
	})

	err := u.userRepo.Create(user)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (u *UserUsecase) Login(username string, password string) (model.User, error) {
	log := logrus.WithFields(logrus.Fields{
		"username": username,
		"password": password,
	})

	user, err := u.userRepo.Login(username)

	if err != nil {
		log.Error(err)
		return model.User{}, model.ErrUsernameNotFound
	}

	if !user.IsPasswordMatch(password) {
		log.Error(model.ErrInvalidPassword)
		return model.User{}, model.ErrInvalidPassword
	}

	return user, nil

}

func (u *UserUsecase) FindByUsername(username string) (model.User, error) {
	log := logrus.WithFields(logrus.Fields{
		"username": username,
	})

	user, err := u.userRepo.FindByUsername(username)
	if err != nil {
		log.Error(err)
		return model.User{}, err
	}

	return user, nil

}
