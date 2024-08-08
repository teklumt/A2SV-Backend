package services

import (
	"context"
	"fmt"
	"jwt_task_manegnment/db"
	"jwt_task_manegnment/model"
	"log"

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

func AddTask(task model.Task, userName string) (bool, string) {
    if task.Title == "" || task.Description == "" {
        return false, ""
    }

    exist, err := db.GetTaskByID(task.ID)
    if err != nil || exist != nil {
        fmt.Println("Task already exists")
      
    }

    
    task.ID = uuid.New().String()
    task.CreaterID = userName
  
    _, err = db.CreateTask(task)
    if err != nil {
        fmt.Println("Error creating task:", err)
        return false, ""
    }
    return true, task.ID
}
