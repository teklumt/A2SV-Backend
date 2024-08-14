package repository

import (
	"clean_architecture_Testing/domain"
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepositorySuite struct {
	suite.Suite
	client     *mongo.Client
	collection *mongo.Collection
	repo       domain.TaskRepository
}

func (suite *TaskRepositorySuite) SetupTest() {
	// Initialize in-memory MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	suite.NoError(err)

	suite.client = client
	suite.collection = client.Database("testdb").Collection("tasks")
	suite.repo = NewTaskRepositoryImpl(*suite.collection)
}

func (suite *TaskRepositorySuite) TearDownTest() {
	// Clean up
	suite.client.Disconnect(context.Background())
}

func (suite *TaskRepositorySuite) TestCreateTask() {
	task := domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "Pending",
		CreaterID:   "user1",
	}

	result, err := suite.repo.CreateTask(task)
	suite.NoError(err)
	suite.Equal(task.Title, result.Title)
	suite.Equal(task.Description, result.Description)
	suite.Equal(task.Status, result.Status)
	suite.Equal(task.CreaterID, result.CreaterID)
}

func (suite *TaskRepositorySuite) TestGetTasks() {
	// Add a task to the collection
	task := domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "Pending",
		CreaterID:   "user1",
	}
	_, err := suite.collection.InsertOne(context.Background(), task)
	suite.NoError(err)

	tasks, err := suite.repo.GetTasks()
	suite.NoError(err)
	suite.Equal(task.Title, tasks[0].Title)
}

func (suite *TaskRepositorySuite) TestGetTaskByID() {
	// Add a task to the collection
	taskID := primitive.NewObjectID()
	task := domain.Task{
		ID:          taskID,
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "Pending",
		CreaterID:   "user1",
	}
	_, err := suite.collection.InsertOne(context.Background(), task)
	suite.NoError(err)

	result, err := suite.repo.GetTaskByID(taskID.Hex(), "user1", "admin")
	suite.NoError(err)
	suite.Equal(task.ID, result.ID)
	suite.Equal(task.Title, result.Title)
}

func (suite *TaskRepositorySuite) TestDeleteTask() {
	// Add a task to the collection
	taskID := primitive.NewObjectID()
	task := domain.Task{
		ID:          taskID,
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "Pending",
		CreaterID:   "user1",
	}
	_, err := suite.collection.InsertOne(context.Background(), task)
	suite.NoError(err)

	deletedTask, err := suite.repo.DeleteTask(taskID.Hex())
	suite.NoError(err)
	suite.Equal(task.ID, deletedTask.ID)
}

func (suite *TaskRepositorySuite) TestUpdateTask() {
	// Add a task to the collection
	taskID := primitive.NewObjectID()
	task := domain.Task{
		ID:          taskID,
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "Pending",
		CreaterID:   "user1",
	}
	_, err := suite.collection.InsertOne(context.Background(), task)
	suite.NoError(err)

	updatedTask := domain.Task{
		ID:          taskID,
		Title:       "Updated Task",
		Description: "Updated Description",
		Status:      "Completed",
		CreaterID:   "user1",
	}
	result, err := suite.repo.UpdateTask(taskID.Hex(), updatedTask)
	suite.NoError(err)
	suite.Equal(updatedTask.Title, result.Title)
	suite.Equal(updatedTask.Description, result.Description)
	suite.Equal(updatedTask.Status, result.Status)
}

func TestTaskRepositorySuite(t *testing.T) {
	suite.Run(t, new(TaskRepositorySuite))
}
