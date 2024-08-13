package repository

import (
	"clean_architecture_Testing/domain"
	"clean_architecture_Testing/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRepositorySuite struct {
	suite.Suite
	mockRepo *mocks.TaskRepository
}

func (suite *TaskRepositorySuite) SetupTest() {
	suite.mockRepo = new(mocks.TaskRepository)
}

func (suite *TaskRepositorySuite) TestCreateTask() {
	task := domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "Pending",
		CreaterID:   "user1",
	}

	suite.mockRepo.On("CreateTask", task).Return(task, nil)

	result, err := suite.mockRepo.CreateTask(task)

	suite.NoError(err)
	suite.Equal(task.Title, result.Title)
	suite.Equal(task.Description, result.Description)
	suite.Equal(task.Status, result.Status)
	suite.Equal(task.CreaterID, result.CreaterID)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskRepositorySuite) TestGetTasks() {
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
	suite.mockRepo.On("GetTasks").Return(tasks, nil)

	result, err := suite.mockRepo.GetTasks()
	suite.NoError(err)
	suite.Equal(tasks, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskRepositorySuite) TestGetTaskByID() {
	taskID := primitive.NewObjectID().Hex()
	taskIDObj, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		suite.T().Errorf("Error converting task ID: %v", err)
		return
	}
	task := domain.Task{
		ID:          taskIDObj,
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "Pending",
		CreaterID:   "user1",
	}
	suite.mockRepo.On("GetTaskByID", taskID, "user1", "user").Return(task, nil)

	result, err := suite.mockRepo.GetTaskByID(taskID, "user1", "user")
	suite.NoError(err)
	suite.Equal(task, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskRepositorySuite) TestGetMyTasks() {
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
	suite.mockRepo.On("GetMyTasks", username).Return(tasks, nil)

	result, err := suite.mockRepo.GetMyTasks(username)
	suite.NoError(err)
	suite.Equal(tasks, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskRepositorySuite) TestDeleteTask() {
	taskID := primitive.NewObjectID().Hex()
	taskIDObj, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		suite.T().Errorf("Error converting task ID: %v", err)
		return
	}
	task := domain.Task{
		ID:          taskIDObj,
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "Pending",
		CreaterID:   "user1",
	}
	suite.mockRepo.On("DeleteTask", taskID).Return(task, nil)

	result, err := suite.mockRepo.DeleteTask(taskID)
	suite.NoError(err)
	suite.Equal(task, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskRepositorySuite) TestUpdateTask() {
	taskID := primitive.NewObjectID().Hex()
	taskIDObj, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		suite.T().Errorf("Error converting task ID: %v", err)
		return
	}
	updatedTask := domain.Task{
		ID:          taskIDObj,
		Title:       "Updated Task",
		Description: "Updated Description",
		Status:      "Completed",
		CreaterID:   "user1",
	}
	suite.mockRepo.On("UpdateTask", taskID, updatedTask).Return(updatedTask, nil)

	result, err := suite.mockRepo.UpdateTask(taskID, updatedTask)
	suite.NoError(err)
	suite.Equal(updatedTask, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestTaskRepositorySuite(t *testing.T) {
	suite.Run(t, new(TaskRepositorySuite))
}
