package repository

import (
	"clean_architecture_Testing/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepositoryImpl struct {
	collection mongo.Collection
}

func NewTaskRepositoryImpl(coll mongo.Collection) domain.TaskRepository {
	return &TaskRepositoryImpl{
		collection: coll,
	}
}

func (tr *TaskRepositoryImpl) CreateTask(task domain.Task) (domain.Task, error) {
	_, err := tr.collection.InsertOne(context.Background(), task)
	return task, err
}




func (tr *TaskRepositoryImpl) GetTasks() ([]domain.Task, error) {
	cursor, err := tr.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var tasks []domain.Task
	for cursor.Next(context.Background()) {
		var task domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (tr *TaskRepositoryImpl) GetTaskByID(id string,  creter string, Role_ string) (domain.Task, error) {
	var task domain.Task
	newId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Task{}, err
	}
	err = tr.collection.FindOne(context.Background(), bson.M{"_id": newId}).Decode(&task)
	if err != nil {
		return domain.Task{}, err
	}
	if Role_ != "admin" &&  task.CreaterID != creter {
		return domain.Task{}, errors.New("you are not the creater of this task")
	} 
	return task, err
}


func (tr *TaskRepositoryImpl) GetMyTasks(username string) ([]domain.Task, error) {
	var task []domain.Task
	// err := tr.collection.FindOne(context.Background(), bson.M{"creater_id": username}).Decode(&task)
	cursor, err := tr.collection.Find(context.Background(), bson.M{"creater_id": username})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var t domain.Task
		if err := cursor.Decode(&t); err != nil {
			return nil, err
		}
		task = append(task, t)
	}


	return task, err
}


func (tr *TaskRepositoryImpl) DeleteTask(id string) (domain.Task, error) {
	var task domain.Task
	newID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Task{}, err
	}
	err = tr.collection.FindOneAndDelete(context.Background(), bson.M{"_id": newID}).Decode(&task)
	return task, err
}


func (tr *TaskRepositoryImpl) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	newID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Task{}, err
	}
	_, err = tr.collection.UpdateOne(context.Background(), bson.M{"_id": newID}, bson.M{"$set": task})
	return task, err
}
