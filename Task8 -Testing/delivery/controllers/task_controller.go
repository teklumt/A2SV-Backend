package controllers

import (
	"clean_architecture_Testing/domain"
	"clean_architecture_Testing/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskUsecase *usecase.TaskUsecase
}

func NewTaskController(taskUsecase *usecase.TaskUsecase) *TaskController {
	return &TaskController{TaskUsecase: taskUsecase}
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.CreaterID = c.GetString("username")
	fmt.Println(task.CreaterID)

	if err := tc.TaskUsecase.CreateTask(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task created successfully"})
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.TaskUsecase.GetTasks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}


func (tc *TaskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	crater := c.GetString("username")
	Role_ := c.GetString("role")
	task, err := tc.TaskUsecase.GetTaskByID(id, crater, Role_ )
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func (tc *TaskController) GetMyTasks(c *gin.Context) {
	username := c.GetString("username")
	tasks, err := tc.TaskUsecase.GetMyTasks(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}


func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.TaskUsecase.GetTaskByID(id, c.GetString("username"), c.GetString("role"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if task.CreaterID != c.GetString("username") && c.GetString("role") != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you are not the creater of this task"})
		return
	}

	if err := tc.TaskUsecase.DeleteTask(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully", "task": task})
}


func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.CreaterID = c.GetString("username")
	fmt.Println(task.CreaterID)

	if err := tc.TaskUsecase.UpdateTask(id, task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task updated successfully"})
}