package usecase

import (
	"clean_architecture/domain"
	"errors"
)

type UserUsecase struct {
	UserRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) *UserUsecase {
	return &UserUsecase{UserRepo: userRepo}
}

func (uc *UserUsecase) RegisterUser(user domain.User) error {
	if user.Username == "" || user.Password == "" {
		return errors.New("missing required fields")
	}
	_, err := uc.UserRepo.CreateUser(user)
	return err
}
