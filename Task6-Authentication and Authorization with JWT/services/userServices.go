package services

import (
	"errors"
	"jwt_task_manegnment/db"
	"jwt_task_manegnment/model"
)

func Login(user *model.User) (string, error) {
	var Role_ string
	if user.Username == "" || user.Password == "" {
		return Role_,errors.New("invalid user data")
	}

	_,Role_, err := db.LogUser(user.Username)
	if err != nil {
		return Role_, err
	}
	return Role_,nil
}

func CreateUser(user model.User) error {
	if user.Username == "" || user.Password == "" {
		return errors.New("invalid user data")
	}

	err := db.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}


func FindAllDataForUser(username string) ([]model.Task, error) {
	tasks, err := db.FindAllDataForUser(username)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}