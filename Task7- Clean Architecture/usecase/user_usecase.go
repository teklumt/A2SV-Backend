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

func (uc *UserUsecase) LoginUser(username string, password string) (domain.User, error) {
	if username == "" || password == "" {
		return domain.User{}, errors.New("missing required fields")
	}
	user, err := uc.UserRepo.LoginUser(username, password)
	return user, err
}

func (uc *UserUsecase) GetAllUsers() ([]domain.User, error) {
	users, err := uc.UserRepo.GetAllUsers()
	return users, err
}

func (uc *UserUsecase) DeleteUserID(id string) (domain.User, error) {
	user, err := uc.UserRepo.DeleteUserID(id)
	return user, err
}

func (uc *UserUsecase) GetUserByID(id string) (domain.User, error) {
	user, err := uc.UserRepo.GetUserByID(id)
	return user, err
}

func (uc *UserUsecase) GetMyProfile(username string) (domain.User, error) {
	user, err := uc.UserRepo.GetMyProfile(username)
	return user, err
}