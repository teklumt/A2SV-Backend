package usecase

import (
	"clean_architecture_Testing/domain"
	"clean_architecture_Testing/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecaseSuite struct {
	suite.Suite
	mockRepo *mocks.TaskRepository
	uc       *TaskUsecase
}

func (suite *TaskUsecaseSuite) SetupTest() {
	suite.mockRepo = new(mocks.TaskRepository)
	suite.uc = NewTaskUsecase(suite.mockRepo)
}

func (suite *TaskUsecaseSuite) TestCreateTask() {
	task := domain.Task{Title: "Test Task", Description: "Test Description"}

	suite.mockRepo.On("CreateTask", task).Return(task, nil)

	err := suite.uc.CreateTask(task)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestCreateTaskWithMissingFields() {
	task := domain.Task{Title: "", Description: ""}

	err := suite.uc.CreateTask(task)

	suite.EqualError(err, "missing required fields")
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestGetTasks() {
	objectID1, _ := primitive.ObjectIDFromHex("66b4da6f8d141a21dcc920bd")
	objectID2, _ := primitive.ObjectIDFromHex("76b4da6f8d141a21dcc920bd")

	tasks := []domain.Task{
		{ID: objectID1, Title: "Task 1", Description: "Description 1"},
		{ID: objectID2, Title: "Task 2", Description: "Description 2"},
	}

	suite.mockRepo.On("GetTasks").Return(tasks, nil)

	result, err := suite.uc.GetTasks()

	suite.NoError(err)
	suite.ElementsMatch(tasks, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestGetTaskByID() {
	objectID, _ := primitive.ObjectIDFromHex("66b4da6f8d141a21dcc920bd")
	task := domain.Task{ID: objectID, Title: "Task 1", Description: "Description 1"}

	suite.mockRepo.On("GetTaskByID", objectID.Hex(), "creator", "role").Return(task, nil)

	result, err := suite.uc.GetTaskByID(objectID.Hex(), "creator", "role")

	suite.NoError(err)
	suite.Equal(task, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestUpdateTask() {
	objectID, _ := primitive.ObjectIDFromHex("66b4da6f8d141a21dcc920bd")
	task := domain.Task{Title: "Updated Task", Description: "Updated Description"}

	suite.mockRepo.On("UpdateTask", objectID.Hex(), task).Return(task, nil)

	err := suite.uc.UpdateTask(objectID.Hex(), task)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestUpdateTaskWithMissingFields() {
	objectID, _ := primitive.ObjectIDFromHex("66b4da6f8d141a21dcc920bd")
	task := domain.Task{Title: "", Description: ""}

	err := suite.uc.UpdateTask(objectID.Hex(), task)

	suite.EqualError(err, "missing required fields")
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestDeleteTask() {
	objectID, _ := primitive.ObjectIDFromHex("66b4da6f8d141a21dcc920bd")

	suite.mockRepo.On("DeleteTask", objectID.Hex()).Return(domain.Task{}, nil)

	err := suite.uc.DeleteTask(objectID.Hex())

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestTaskUsecaseSuite(t *testing.T) {
	suite.Run(t, new(TaskUsecaseSuite))
}
