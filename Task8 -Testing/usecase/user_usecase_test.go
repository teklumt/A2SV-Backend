package usecase

import (
	"clean_architecture_Testing/domain"
	"clean_architecture_Testing/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecaseSuite struct {
	suite.Suite
	mockRepo *mocks.UserRepository
	uc       *UserUsecase
}

func (suite *UserUsecaseSuite) SetupTest() {
	suite.mockRepo = new(mocks.UserRepository)
	suite.uc = NewUserUsecase(suite.mockRepo)
}

func (suite *UserUsecaseSuite) TestRegisterUser() {
	user := domain.User{Username: "testuser", Password: "password123"}

	suite.mockRepo.On("CreateUser", user).Return(user, nil)

	err := suite.uc.RegisterUser(user)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestRegisterUserWithMissingFields() {
	user := domain.User{Username: "", Password: ""}

	err := suite.uc.RegisterUser(user)

	suite.EqualError(err, "missing required fields")
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestLoginUser() {
	user := domain.User{Username: "testuser", Password: "password123"}

	suite.mockRepo.On("LoginUser", "testuser", "password123").Return(user, nil)

	result, err := suite.uc.LoginUser("testuser", "password123")

	suite.NoError(err)
	suite.Equal(user, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestLoginUserWithMissingFields() {
	result, err := suite.uc.LoginUser("", "")

	suite.EqualError(err, "missing required fields")
	suite.Empty(result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestGetAllUsers() {
	objectID1, _ := primitive.ObjectIDFromHex("66b4da6f8d141a21dcc920bd")
	objectID2, _ := primitive.ObjectIDFromHex("76b4da6f8d141a21dcc920bd")

	users := []domain.User{
		{ID: objectID1, Username: "user1", Password: "password1"},
		{ID: objectID2, Username: "user2", Password: "password2"},
	}

	suite.mockRepo.On("GetAllUsers").Return(users, nil)

	result, err := suite.uc.GetAllUsers()

	suite.NoError(err)
	suite.ElementsMatch(users, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestGetUserByID() {
	objectID, _ := primitive.ObjectIDFromHex("66b4da6f8d141a21dcc920bd")
	user := domain.User{ID: objectID, Username: "user1", Password: "password1"}

	suite.mockRepo.On("GetUserByID", objectID.Hex()).Return(user, nil)

	result, err := suite.uc.GetUserByID(objectID.Hex())

	suite.NoError(err)
	suite.Equal(user, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestGetMyProfile() {
	user := domain.User{Username: "testuser", Password: "password123"}

	suite.mockRepo.On("GetMyProfile", "testuser").Return(user, nil)

	result, err := suite.uc.GetMyProfile("testuser")

	suite.NoError(err)
	suite.Equal(user, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestDeleteUserID() {
	objectID, _ := primitive.ObjectIDFromHex("66b4da6f8d141a21dcc920bd")

	suite.mockRepo.On("DeleteUserID", objectID.Hex()).Return(domain.User{}, nil)

	_, err := suite.uc.DeleteUserID(objectID.Hex())

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseSuite))
}
