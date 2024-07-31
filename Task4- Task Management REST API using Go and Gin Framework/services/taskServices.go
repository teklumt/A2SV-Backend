package services

import (
	"task_manager/db"
	"task_manager/model"

	"github.com/google/uuid"
)

func FindAllData() []model.Task {
    return db.Database
}

func FindByIndex(id string) model.Task {
    for _, task := range db.Database {
        if task.ID == id {
            return task
        }
    }
    return model.Task{}
}

func UpdateTask(id string, updatedTask model.Task) bool {
	updatedTask.ID = id
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
    if  task.Title == "" || task.Description == "" {
        return false
    }
    for _, t := range db.Database {
        if t.Title == task.Title && t.Description == task.Description {
            return false
        }
    }


    task.ID = uuid.New().String() 
    db.Database = append(db.Database, task)
    return true
}
