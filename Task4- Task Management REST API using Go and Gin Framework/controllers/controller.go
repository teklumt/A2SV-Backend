package controllers

import (
	"net/http"
	"task_manager/model"
	"task_manager/services"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(context *gin.Context) {
    context.IndentedJSON(http.StatusOK, gin.H{"tasks": services.FindAllData()})
}

func GetSpecificTask(context *gin.Context) {
    id := context.Param("id")
    task := services.FindByIndex(id)
    if task == (model.Task{}) {
        context.IndentedJSON(http.StatusNotFound, gin.H{
            "message": "Task not found",
            "error": "404 Not Found",

        })
        return
    }
    context.IndentedJSON(http.StatusOK, gin.H{"task": task})
}

func UpdateSpecificTask(context *gin.Context) {
    id := context.Param("id")
    var updatedTask model.Task
    if err := context.ShouldBindJSON(&updatedTask); err != nil {
        context.IndentedJSON(http.StatusBadRequest,gin.H{"message": "Invalid JSON", "error": err.Error()})
        return
    }
    updated := services.UpdateTask(id, updatedTask)
    if updated {
        context.IndentedJSON(http.StatusOK, gin.H{
            "message": "Successfully Updated",
            "task":    updatedTask,
        
        })
        return
    }
    context.IndentedJSON(http.StatusNotFound, gin.H{
        "message": "Resource Not Found",
        "error":   "404 Not Found",
    })
}

func DeleteSpecificTask(context *gin.Context) {
    id := context.Param("id")
    deleted := services.DeleteTask(id)
    if deleted {
        context.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully Deleted"})
        return
    }
    context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Resource Not Found"})
}

func AddSpecificTask(context *gin.Context) {
    var newTask model.Task
    if err := context.ShouldBindJSON(&newTask); err != nil {
        context.IndentedJSON(http.StatusBadRequest, err.Error())
        return
    }
    added := services.AddTask(newTask)
    if added {
        context.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully Added", "task": newTask})
        return
    }
    context.IndentedJSON(http.StatusNotAcceptable, gin.H{
        "message": "Not Added Successfully",
    })
}
