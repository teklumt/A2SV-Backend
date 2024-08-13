package controllers

import (
	"bytes"
	"clean_architecture_Testing/domain"
	"clean_architecture_Testing/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func setupRouter(taskUsecase domain.TaskUsecase) *gin.Engine {
	r := gin.Default()
	taskController := NewTaskController(taskUsecase)
	r.POST("/tasks", taskController.CreateTask)
	r.GET("/tasks", taskController.GetTasks)
	r.GET("/tasks/:id", taskController.GetTaskByID)
	r.GET("/me", taskController.GetMyTasks)
	r.DELETE("/tasks/:id", taskController.DeleteTask)
	r.PUT("/tasks/:id", taskController.UpdateTask)
	return r
}

func TestCreateTask(t *testing.T) {
    mockUsecase := new(mocks.TaskUsecase)
    mockController := NewTaskController(mockUsecase)

   
    task := domain.Task{
        Title:       "Test Task",
        Description: "Test Description",
        Status:      "Pending",
    }

    mockUsecase.On("CreateTask", task).Return(nil)

    router := gin.Default()
    router.POST("/tasks", mockController.CreateTask)

    jsonData, _ := json.Marshal(task)
    req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonData))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("username", "user1") 
    rr := httptest.NewRecorder()

    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
    assert.JSONEq(t, `{"message":"task created successfully"}`, rr.Body.String())

    mockUsecase.AssertExpectations(t)
}


func TestGetTasks(t *testing.T) {
	mockUsecase := new(mocks.TaskUsecase)
	tasks := []domain.Task{
		{Title: "Task 1"},
		{Title: "Task 2"},
	}

	mockUsecase.On("GetTasks").Return(tasks, nil)

	router := setupRouter(mockUsecase)

	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string][]domain.Task
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Error unmarshalling response: %v", err)
	}

	assert.Len(t, response["tasks"], 2)
	mockUsecase.AssertExpectations(t)
}


func TestGetTaskByID(t *testing.T) {
    mockUsecase := new(mocks.TaskUsecase)
    taskController := NewTaskController(mockUsecase)

    taskID := primitive.NewObjectID().Hex()
    taskIDObj, err := primitive.ObjectIDFromHex(taskID)
    if err != nil {
        t.Fatalf("Error converting task ID: %v", err)
    }
    
    task := domain.Task{
        ID:          taskIDObj,
        Title:       "Test Task",
        Description: "Test Description",
        Status:      "Pending",
        CreaterID:   "user1",
    }

    mockUsecase.On("GetTaskByID", taskID, "user1", "user").Return(task, nil)

    router := gin.Default()
    router.GET("/tasks/:id", func(c *gin.Context) {
        c.Set("username", "user1")
        c.Set("role", "user")
        taskController.GetTaskByID(c)
    })

    req, _ := http.NewRequest("GET", "/tasks/"+taskID, nil)
    rr := httptest.NewRecorder()

    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
    expected := `{"task": {"_id":"` + taskID + `", "title":"Test Task", "description":"Test Description", "status":"Pending", "creater_id":"user1"}}`
    assert.JSONEq(t, expected, rr.Body.String())

    mockUsecase.AssertExpectations(t)
}



func TestDeleteTask(t *testing.T) {
	mockUsecase := new(mocks.TaskUsecase)
    taskController := NewTaskController(mockUsecase)

    taskID := primitive.NewObjectID().Hex()
    taskIDObj, err := primitive.ObjectIDFromHex(taskID)
    if err != nil {
        t.Fatalf("Error converting task ID: %v", err)
    }
	task := domain.Task{
		ID:         taskIDObj,
		Title:       "Task to Delete",
		Description: "To be deleted",
		Status:      "Pending",
		CreaterID:   "user1",
	}

	mockUsecase.On("GetTaskByID", taskID, "user1", "user").Return(task, nil)
	mockUsecase.On("DeleteTask", taskID).Return(nil)

	router := gin.Default()
	router.DELETE("/tasks/:id", func(c *gin.Context) {
		c.Set("username", "user1")
		c.Set("role", "user")
		taskController.DeleteTask(c)
	})
	req, _ := http.NewRequest("DELETE", "/tasks/"+taskID, nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"message":"task deleted successfully", "task": {"_id":"`+taskID+`", "title":"Task to Delete", "description":"To be deleted", "status":"Pending", "creater_id":"user1"}}`, rr.Body.String())

	mockUsecase.AssertExpectations(t)

}


func TestUpdateTask(t *testing.T) {
	mockUsecase := new(mocks.TaskUsecase)
    taskController := NewTaskController(mockUsecase)

    taskID := primitive.NewObjectID().Hex()
    taskIDObj, err := primitive.ObjectIDFromHex(taskID)
    if err != nil {
        t.Fatalf("Error converting task ID: %v", err)
    }
	task := domain.Task{
		ID:         taskIDObj,
		Title:       "Task to Delete",
		Description: "To be deleted",
		Status:      "Pending",
		CreaterID:   "user1",
	}

	mockUsecase.On("UpdateTask", taskID, task).Return(nil)
	
	router := gin.Default()
	router.PUT("/tasks/:id", func(c *gin.Context) {
		c.Set("username", "user1")
		c.Set("role", "user")
		taskController.UpdateTask(c)
	})

	jsonData, _ := json.Marshal(task)
	req, _ := http.NewRequest("PUT", "/tasks/"+taskID, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("username", "user1")
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)


	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"message":"task updated successfully"}`, rr.Body.String())

	mockUsecase.AssertExpectations(t)

}