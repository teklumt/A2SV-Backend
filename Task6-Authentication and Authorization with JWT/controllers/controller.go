package controllers

import (
	"jwt_task_manegnment/model"
	"jwt_task_manegnment/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
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
    userClaims, exists := context.Get("user")
    if !exists {
        context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
        return
    }
    user, ok := userClaims.(jwt.MapClaims)
    if !ok {
        context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Invalid user claims"})
        return
    }

    Role_, ok := user["role"].(string)
    userName, ok := user["username"].(string)
    if !ok || Role_ == "" || userName == "" {
        context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Invalid Role in claims"})

        return
    }

    if Role_ != "admin" {
        task_ := services.FindByIndex(id)
        if task_ == (model.Task{}) {
            context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Resource Not Found"})
            return
        }
        if task_.CreaterID != user["username"] {
            context.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized Access"})
            return
        }
    }



    updatedTask.CreaterID = userName
    updated := services.UpdateTask(id, updatedTask)
    if updated {
        updatedTask.ID = id
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
    userClaims, exists := context.Get("user")

    if !exists {
        context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
        return
    }
    
    user, ok := userClaims.(jwt.MapClaims)
    if !ok {
        context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Invalid user claims"})
        return
    }

    Role_, ok := user["role"].(string)

    if !ok || Role_ == "" {
        context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Invalid Role in claims"})
        return
    }

    if Role_ != "admin" {
        task_ := services.FindByIndex(id)

        if task_ == (model.Task{}) {
            context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Resource Not Found"})
            return
        }

        if task_.CreaterID != user["username"] {
            context.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized Access"})
            return
        }
    }

        




    deleted := services.DeleteTask(id)
    if deleted {
        context.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully Deleted"})
        return
    }
    context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Resource Not Found"})
}

func AddSpecificTask(c *gin.Context) {
    var newTask model.Task

  
    userClaims, exists := c.Get("user")
    if !exists {
        c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
        return
    }

    user, ok := userClaims.(jwt.MapClaims)
    if !ok {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Invalid user claims"})
        return
    }


    username, ok := user["username"].(string)
    // fmt.Println(username)
    if !ok || username == "" {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Invalid username in claims"})
        return
    }

 
    if err := c.ShouldBindJSON(&newTask); err != nil {
        c.IndentedJSON(http.StatusBadRequest, err.Error())
        return
    }


    added, newID := services.AddTask(newTask, username)

    if added {
        newTask.ID = newID
        newTask.CreaterID = username
        c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully Added", "task": newTask})
        return
    }

    c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Not Added Successfully"})
}


func GetAllTasksForUser(c *gin.Context) {
    userClaims, exists := c.Get("user")
    if !exists {
        c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
        return
    }

    user, ok := userClaims.(jwt.MapClaims)
    if !ok {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Invalid user claims"})
        return
    }

    username, ok := user["username"].(string)
    if !ok || username == "" {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Invalid username in claims"})
        return
    }
    data, err := services.FindAllDataForUser(username)  
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"tasks": data})
}