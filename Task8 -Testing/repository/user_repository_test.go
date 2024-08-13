package repository

import (
	"clean_architecture_Testing/domain"
	"clean_architecture_Testing/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepositorySuite struct {
	suite.Suite
	mockRepo *mocks.UserRepository
}

func (suite *UserRepositorySuite) SetupTest() {
	suite.mockRepo = new(mocks.UserRepository)
}

func (suite *UserRepositorySuite) TestCreateUser() {
	user := domain.User{
		Username: "Teklu Moges",
		Password: "Teklumo",
	}

	suite.mockRepo.On("CreateUser", user).Return(user, nil)

	result, err := suite.mockRepo.CreateUser(user)
	suite.NoError(err)
	suite.Equal(user.Username, result.Username)
	suite.Equal(user.Password, result.Password)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestDeleteUserID() {
	userID := primitive.NewObjectID().Hex()
	userObjId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		suite.T().Errorf("Error converting user ID: %v", err)
		return
	}
	user := domain.User{
		ID:       userObjId,
		Username: "Teklu Moges",
		Password: "teklumo",
	}

	suite.mockRepo.On("DeleteUserID", userID).Return(user, nil)
	result, err := suite.mockRepo.DeleteUserID(userID)

	suite.NoError(err)
	suite.Equal(user, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestGetAllUsers() {
	users := []domain.User{
		{Username: "Teklu Moges", Password: "teklumo"},
		{Username: "Yehone sew", Password: "yehonesew"},
	}

	suite.mockRepo.On("GetAllUsers").Return(users, nil)

	result, err := suite.mockRepo.GetAllUsers()

	suite.NoError(err)
	suite.ElementsMatch(users, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestGetMyProfile() {
	username := "Teklu Moges"
	user := domain.User{
		Username: username,
		Password: "teklumo",
	}

	suite.mockRepo.On("GetMyProfile", username).Return(user, nil)

	result, err := suite.mockRepo.GetMyProfile(username)

	suite.NoError(err)
	suite.Equal(user, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestGetUserByID() {
	userID := primitive.NewObjectID().Hex()
	userObjId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		suite.T().Errorf("Error converting user ID: %v", err)
		return
	}
	user := domain.User{
		ID:       userObjId,
		Username: "Teklu Moges",
		Password: "teklumo",
	}

	suite.mockRepo.On("GetUserByID", userID).Return(user, nil)

	result, err := suite.mockRepo.GetUserByID(userID)

	suite.NoError(err)
	suite.Equal(user, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestLoginUser() {
	username := "Teklu Moges"
	password := "teklumo"
	user := domain.User{
		Username: username,
		Password: password,
	}

	suite.mockRepo.On("LoginUser", username, password).Return(user, nil)

	result, err := suite.mockRepo.LoginUser(username, password)

	suite.NoError(err)
	suite.Equal(user, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}
