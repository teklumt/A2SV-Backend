package services

import (
	"context"
	"log"
	"task_manager_DB/db"
	"task_manager_DB/model"

	"github.com/google/uuid"
)

func FindAllData() []model.Task {
    cursor, err := db.GetAllData()
    if err != nil {
        log.Println("Error retrieving data:", err)
        return []model.Task{}
    }
    defer cursor.Close(context.Background())

    var tasks []model.Task
    for cursor.Next(context.Background()) {
        var task model.Task
        if err := cursor.Decode(&task); err != nil {
            log.Println("Error decoding task:", err)
            continue
        }
        tasks = append(tasks, task)
    }

    if err := cursor.Err(); err != nil {
        log.Println("Cursor error:", err)
        return []model.Task{}
    }

    return tasks
}

func FindByIndex(id string) model.Task {
    newTask, err := db.GetTaskByID(id)
    if err != nil {
        return model.Task{}
    }
    return *newTask
}

func UpdateTask(id string, updatedTask model.Task) bool {
	updatedTask.ID = id
    _, err := db.UpdateTask(id, updatedTask)
    if err != nil {
        return false
    }
    return true
}

func DeleteTask(id string) bool {
   
    _, err := db.DeleteTask(id)
    if err != nil {
        return false
    }
    return true
}

func AddTask(task model.Task) (bool, string)  {
    if  task.Title == "" || task.Description == "" {
        return false,""
    }

    exist, _ := db.GetTaskByID(task.ID)
    if exist != nil {
        return false,""
    }
    
  
    task.ID = uuid.New().String() 
    _, err := db.CreateTask(task)
    if err != nil {
        return false, ""
    }
    return true, task.ID



}
