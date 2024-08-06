package db

import (
	"context"
	"errors"
	"jwt_task_manegnment/model"

	"go.mongodb.org/mongo-driver/bson"
)



func LogUser(username string) (bool,string, error) {
	var user model.User
	var Role_ string
	// err := UserCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	err := UserCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	Role_ = user.Role
	if err != nil {
		return false,Role_, err
	}
	return true,Role_, nil
}

func CreateUser(user model.User) error {
	if res, _ , _:= LogUser(user.Username); res {
		return errors.New("user exists")

	}
	user.Role = "user"
	_, err := UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}


func FindAllDataForUser(username string) ([]model.Task, error) {
	var tasks []model.Task
	cursor, err := TaskCollection.Find(context.Background(), bson.M{"creater_id": username})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.Background(), &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
