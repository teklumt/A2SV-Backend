package usecase

import (
	"clean_architecture_Testing/domain"
	"clean_architecture_Testing/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateTask(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)
	uc := NewTaskUsecase(mockRepo)

	task := domain.Task{Title: "Test Task", Description: "Test Description"}

	mockRepo.On("CreateTask", task).Return(task, nil)

	err := uc.CreateTask(task)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateTaskWithMissingFields(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)
	uc := NewTaskUsecase(mockRepo)

	task := domain.Task{Title: "", Description: ""}

	err := uc.CreateTask(task)

	assert.EqualError(t, err, "missing required fields")
	mockRepo.AssertExpectations(t)
}

func TestGetTasks(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)
	uc := NewTaskUsecase(mockRepo)

	objectID1, _ := primitive.ObjectIDFromHex("66b4da6f8d141a21dcc920bd")
	objectID2, _ := primitive.ObjectIDFromHex("76b4da6f8d141a21dcc920bd")

	tasks := []domain.Task{
		{ID: objectID1, Title: "Task 1", Description: "Description 1"},
		{ID: objectID2, Title: "Task 2", Description: "Description 2"},
	}


	mockRepo.On("GetTasks").Return(tasks, nil)


	result, err := uc.GetTasks()


	assert.NoError(t, err)
	assert.ElementsMatch(t, tasks, result) 
	mockRepo.AssertExpectations(t)
}

func TestGetTaskByID(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)
	uc := NewTaskUsecase(mockRepo)

	objectID, _ := primitive.ObjectIDFromHex("66b4da6f8d141a21dcc920bd")
	task := domain.Task{ID: objectID, Title: "Task 1", Description: "Description 1"}


	mockRepo.On("GetTaskByID", objectID.Hex(), "creator", "role").Return(task, nil)

	
	result, err := uc.GetTaskByID(objectID.Hex(), "creator", "role")


	assert.NoError(t, err)
	assert.Equal(t, task, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)
	uc := NewTaskUsecase(mockRepo)

	objectID, _ := primitive.ObjectIDFromHex("66b4da6f8d141a21dcc920bd")
	task := domain.Task{Title: "Updated Task", Description: "Updated Description"}

	mockRepo.On("UpdateTask", objectID.Hex(), task).Return(task, nil)

	err := uc.UpdateTask(objectID.Hex(), task)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTaskWithMissingFields(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)
	uc := NewTaskUsecase(mockRepo)

	objectID, _ := primitive.ObjectIDFromHex("66b4da6f8d141a21dcc920bd")
	task := domain.Task{Title: "", Description: ""}

	err := uc.UpdateTask(objectID.Hex(), task)

	assert.EqualError(t, err, "missing required fields")
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)
	uc := NewTaskUsecase(mockRepo)

	objectID, _ := primitive.ObjectIDFromHex("66b4da6f8d141a21dcc920bd")

	mockRepo.On("DeleteTask", objectID.Hex()).Return(domain.Task{}, nil)

	err := uc.DeleteTask(objectID.Hex())

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
