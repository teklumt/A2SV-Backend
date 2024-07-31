package services

import (
	"task_manager/db"
	"task_manager/model"
)

func FindAllData()[]model.Task{
	return db.Database
}


func FindByIndex(id string) model.Task {
	for i, task := range db.Database {
		if task.ID == id {
			return db.Database[i]
		}
	}
	return model.Task{}
}



func UpdateTask(id string, updatedTask model.Task) bool {
	for i, task := range db.Database {
		if task.ID == id {
			db.Database[i] = updatedTask
			return true
		}
	}
	return false
}


func DeleteTask(id string) bool {
	for i, task := range db.Database {
		if task.ID == id {
			db.Database = append(db.Database[:i], db.Database[i+1:]...)
			return true
		}
	}
	return false
}



func AddTask(task model.Task) bool {

	db.Database = append(db.Database, task)
	return true
}
