package repository

import (
	"clean_architecture_Testing/domain"
	"clean_architecture_Testing/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateTask(t *testing.T) {
	mockCollaction := new(mocks.TaskRepository)

	task := domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "Pending",
		CreaterID:   "user1",
	}


	mockCollaction.On("CreateTask", task).Return(task, nil)


	result, err := mockCollaction.CreateTask(task)
	
	
	assert.NoError(t, err)
	assert.Equal(t, task.Title, result.Title)
	assert.Equal(t, task.Description, result.Description)
	assert.Equal(t, task.Status, result.Status)
	assert.Equal(t, task.CreaterID, result.CreaterID)


	mockCollaction.AssertExpectations(t)
}



func TestGetTasks(t *testing.T) {
	mockColl := new(mocks.TaskRepository)
	tasks := []domain.Task{
		{
			ID:          primitive.NewObjectID(),
			Title:       "Test Task 1",
			Description: "Test Description 1",
			Status:      "Pending",
			CreaterID:   "user1",
		},
		{
			ID:          primitive.NewObjectID(),
			Title:       "Test Task 2",
			Description: "Test Description 2",
			Status:      "Completed",
			CreaterID:   "user2",
		},
	}
	mockColl.On("GetTasks").Return(tasks, nil)

	result, err := mockColl.GetTasks()
	assert.NoError(t, err)
	assert.Equal(t, tasks, result)
	mockColl.AssertExpectations(t)
}

func TestGetTaskByID(t *testing.T) {
	mockColl := new(mocks.TaskRepository)
	taskID := primitive.NewObjectID().Hex()
	taskIDObj, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		t.Errorf("Error converting task ID: %v", err)
		return
	}
	
	task := domain.Task{
		ID:          taskIDObj,
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "Pending",
		CreaterID:   "user1",
	}
	mockColl.On("GetTaskByID", taskID, "user1", "user").Return(task, nil)


	result, err := mockColl.GetTaskByID(taskID, "user1", "user")
	assert.NoError(t, err)
	assert.Equal(t, task, result)
	mockColl.AssertExpectations(t)
}

func TestGetMyTasks(t *testing.T) {
	mockColl := new(mocks.TaskRepository)
	username := "user1"
	tasks := []domain.Task{
		{
			ID:          primitive.NewObjectID(),
			Title:       "Test Task 1",
			Description: "Test Description 1",
			Status:      "Pending",
			CreaterID:   username,
		},
	}
	mockColl.On("GetMyTasks", username).Return(tasks, nil)


	result, err := mockColl.GetMyTasks(username)
	assert.NoError(t, err)
	assert.Equal(t, tasks, result)
	mockColl.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	mockColl := new(mocks.TaskRepository)
	taskID := primitive.NewObjectID().Hex()
	taskIDObj, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		t.Errorf("Error converting task ID: %v", err)
		return
	}
	task := domain.Task{
		ID:          taskIDObj,
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "Pending",
		CreaterID:   "user1",
	}
	mockColl.On("DeleteTask", taskID).Return(task, nil)


	result, err := mockColl.DeleteTask(taskID)
	assert.NoError(t, err)
	assert.Equal(t, task, result)
	mockColl.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	mockColl := new(mocks.TaskRepository)
	taskID := primitive.NewObjectID().Hex()
	taskIDObj, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		t.Errorf("Error converting task ID: %v", err)
		return
	}
	updatedTask := domain.Task{
		ID:          taskIDObj,
		Title:       "Updated Task",
		Description: "Updated Description",
		Status:      "Completed",
		CreaterID:   "user1",
	}
	mockColl.On("UpdateTask", taskID, updatedTask).Return(updatedTask, nil)


	result, err := mockColl.UpdateTask(taskID, updatedTask)
	assert.NoError(t, err)
	assert.Equal(t, updatedTask, result)
	mockColl.AssertExpectations(t)
}
