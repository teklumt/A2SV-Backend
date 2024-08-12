package usecase

import (
	"clean_architecture_Testing/domain"
	"clean_architecture_Testing/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestRegisterUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewUserUsecase(mockRepo)

	user := domain.User{Username: "testuser", Password: "password123"}

	// Define expectations
	mockRepo.On("CreateUser", user).Return(user, nil)

	// Call the method under test
	err := uc.RegisterUser(user)

	// Assert the result
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRegisterUserWithMissingFields(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewUserUsecase(mockRepo)

	user := domain.User{Username: "", Password: ""}

	// Call the method under test
	err := uc.RegisterUser(user)

	// Assert the result
	assert.EqualError(t, err, "missing required fields")
	mockRepo.AssertExpectations(t)
}

func TestLoginUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewUserUsecase(mockRepo)

	user := domain.User{Username: "testuser", Password: "password123"}

	// Define expectations
	mockRepo.On("LoginUser", "testuser", "password123").Return(user, nil)

	// Call the method under test
	result, err := uc.LoginUser("testuser", "password123")

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
}

func TestLoginUserWithMissingFields(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewUserUsecase(mockRepo)

	// Call the method under test
	result, err := uc.LoginUser("", "")

	// Assert the result
	assert.EqualError(t, err, "missing required fields")
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestGetAllUsers(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewUserUsecase(mockRepo)

	objectID1, _ := primitive.ObjectIDFromHex("66b4da6f8d141a21dcc920bd")
	objectID2, _ := primitive.ObjectIDFromHex("76b4da6f8d141a21dcc920bd")

	users := []domain.User{
		{ID: objectID1, Username: "user1", Password: "password1"},
		{ID: objectID2, Username: "user2", Password: "password2"},
	}

	// Define expectations
	mockRepo.On("GetAllUsers").Return(users, nil)

	// Call the method under test
	result, err := uc.GetAllUsers()

	// Assert the results
	assert.NoError(t, err)
	assert.ElementsMatch(t, users, result) // ElementsMatch allows comparing slices regardless of order
	mockRepo.AssertExpectations(t)
}

func TestGetUserByID(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewUserUsecase(mockRepo)

	objectID, _ := primitive.ObjectIDFromHex("66b4da6f8d141a21dcc920bd")
	user := domain.User{ID: objectID, Username: "user1", Password: "password1"}

	// Define expectations
	mockRepo.On("GetUserByID", objectID.Hex()).Return(user, nil)

	// Call the method under test
	result, err := uc.GetUserByID(objectID.Hex())

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
}

func TestGetMyProfile(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewUserUsecase(mockRepo)

	user := domain.User{Username: "testuser", Password: "password123"}

	// Define expectations
	mockRepo.On("GetMyProfile", "testuser").Return(user, nil)

	// Call the method under test
	result, err := uc.GetMyProfile("testuser")

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUserID(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewUserUsecase(mockRepo)

	objectID, _ := primitive.ObjectIDFromHex("66b4da6f8d141a21dcc920bd")

	// Define expectations
	mockRepo.On("DeleteUserID", objectID.Hex()).Return(domain.User{}, nil)

	// Call the method under test
	_, err := uc.DeleteUserID(objectID.Hex())

	// Assert the result
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
