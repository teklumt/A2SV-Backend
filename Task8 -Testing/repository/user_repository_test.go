package repository

import (
	"clean_architecture_Testing/domain"
	"clean_architecture_Testing/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateUser(t *testing.T){
	mockCollaction := new(mocks.UserRepository)

	user := domain.User{
		Username:"Teklu Moges",
		Password:"Teklumo" ,
	}

	mockCollaction.On("CreateUser", user).Return(user, nil)

	result, err := mockCollaction.CreateUser(user)
	assert.NoError(t, err)
	assert.Equal(t, user.Username, result.Username)
	assert.Equal(t, user.Password, result.Password)

	mockCollaction.AssertExpectations(t) 

}


func TestDeleteUserID(t *testing.T){
	mockCollaction := new(mocks.UserRepository)
	userID := primitive.NewObjectID().Hex()
	userObjId, err := primitive.ObjectIDFromHex(userID)
	if err != nil{
		t.Errorf("Error converting task ID: %v", err)
		return 
	}
	user := domain.User{
		ID:		userObjId,
		Username: "Teklu Moges",
		Password: "teklumo",
	}

	mockCollaction.On("DeleteUserID", userID).Return(user, nil)
	result, err := mockCollaction.DeleteUserID(userID)

	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockCollaction.AssertExpectations(t)

}


func TestGetAllUsers(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	users := []domain.User{
		{Username: "Teklu Moges", Password: "teklumo"},
		{Username: "Yehone sew", Password: "yehonesew"},
	}

	mockRepo.On("GetAllUsers").Return(users, nil)

	result, err := mockRepo.GetAllUsers()

	assert.NoError(t, err)
	assert.ElementsMatch(t, users, result) 
	mockRepo.AssertExpectations(t)
}


func TestGetMyProfile(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	username := "Teklu Moges"
	user := domain.User{
		Username: username,
		Password: "teklumo",
	}

	mockRepo.On("GetMyProfile", username).Return(user, nil)

	result, err := mockRepo.GetMyProfile(username)

	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
}


func TestGetUserByID(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userID := primitive.NewObjectID().Hex()
	userObjId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		t.Errorf("Error converting user ID: %v", err)
		return
	}
	user := domain.User{
		ID:       userObjId,
		Username: "Teklu Moges",
		Password: "teklumo",
	}

	mockRepo.On("GetUserByID", userID).Return(user, nil)

	result, err := mockRepo.GetUserByID(userID)

	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
}


func TestLoginUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	username := "Teklu Moges"
	password := "teklumo"
	user := domain.User{
		Username: username,
		Password: password,
	}

	mockRepo.On("LoginUser", username, password).Return(user, nil)

	result, err := mockRepo.LoginUser(username, password)

	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
}
