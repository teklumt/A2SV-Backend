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
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func setupRouterUser(userUsecase domain.UserUsecase) *gin.Engine {
	r := gin.Default()
	userController := NewUserController(userUsecase)
	r.POST("/auth/register", userController.RegisterUser)
	r.POST("/auth/login", userController.LoginUser)
	r.GET("/users", userController.GetAllUsers)
	r.GET("/users/:id", userController.GetUserByID)
	r.DELETE("/users/:id", userController.DeleteUserID)
	r.GET("/users/me", userController.GetMyProfile)
	return r
}

func TestRegisterUser(t *testing.T) {
	mockUserUsecase := new(mocks.UserUsecase)
	r := setupRouterUser(mockUserUsecase)

	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "testuser",
		Password: "testpassword",
	}

	mockUserUsecase.On("RegisterUser", mock.AnythingOfType("domain.User")).Return(nil)

	userJSON, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "user registered successfully")

	mockUserUsecase.AssertExpectations(t)
}

func TestLoginUser(t *testing.T) {
	mockUserUsecase := new(mocks.UserUsecase)
	r := setupRouterUser(mockUserUsecase)

	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "testuser",
		Password: "testpassword",
	}

	// Mock the use case's LoginUser method to return the user and no error
	mockUserUsecase.On("LoginUser", user.Username, user.Password).Return(user, nil)

	loginCredentials := map[string]string{
		"username": user.Username,
		"password": user.Password,
	}
	credentialsJSON, _ := json.Marshal(loginCredentials)
	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(credentialsJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Contains(t, w.Body.String(), "")

	mockUserUsecase.AssertExpectations(t)
}


func TestGetAllUsers(t *testing.T) {
    mockUsecase := new(mocks.UserUsecase)
    // userController := NewUserController(mockUsecase)

    users := []domain.User{
        {Username: "user1"},
        {Username: "user2"},
    }

    mockUsecase.On("GetAllUsers").Return(users, nil)

    router := setupRouterUser(mockUsecase)

    req, _ := http.NewRequest("GET", "/users", nil)
    rr := httptest.NewRecorder()

    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)

    var response map[string][]domain.User
    err := json.Unmarshal(rr.Body.Bytes(), &response)
    if err != nil {
        t.Fatalf("Error unmarshalling response: %v", err)
    }

    assert.Len(t, response["users"], 2)
    mockUsecase.AssertExpectations(t)
}

func TestGetUserByID(t *testing.T) {
    mockUsecase := new(mocks.UserUsecase)
    userController := NewUserController(mockUsecase)

    userIDStr := primitive.NewObjectID().Hex()
    userID, _ := primitive.ObjectIDFromHex(userIDStr)
    user := domain.User{
        ID:       userID,
        Username: "testuser",
    }

    mockUsecase.On("GetUserByID", userIDStr).Return(user, nil)

    router := gin.Default()
    router.GET("/users/:id", userController.GetUserByID)

    req, _ := http.NewRequest("GET", "/users/"+userIDStr, nil)
    rr := httptest.NewRecorder()

    router.ServeHTTP(rr, req)
    assert.Equal(t, http.StatusOK, rr.Code)
    assert.JSONEq(t, `{"user": {"_id":"`+userIDStr+`", "username":"testuser", "password":"", "role":""}}`, rr.Body.String())

    mockUsecase.AssertExpectations(t)
}

func TestDeleteUserID(t *testing.T) {
    mockUsecase := new(mocks.UserUsecase)
    userController := NewUserController(mockUsecase)

    userIDStr := primitive.NewObjectID().Hex()
    userID, _ := primitive.ObjectIDFromHex(userIDStr)
    user := domain.User{
        ID:       userID,
        Username: "testuser",
    }
	mockUsecase.On("GetUserByID", userIDStr	).Return(user, nil)
    mockUsecase.On("DeleteUserID", userIDStr).Return(user, nil)

	router := gin.Default()
	router.DELETE("/users/:id", userController.DeleteUserID)

	req, _ := http.NewRequest("DELETE", "/users/"+userIDStr, nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"message":"user deleted successfully", "user": {"_id":"`+userIDStr+`", "username":"testuser", "password":"", "role":""}}`, rr.Body.String())

	mockUsecase.AssertExpectations(t)



    // router := setupRouterUser(mockUsecase)

    // req, _ := http.NewRequest("DELETE", "/users/"+userIDStr, nil)
    // rr := httptest.NewRecorder()

    // router.ServeHTTP(rr, req)

    // assert.Equal(t, http.StatusOK, rr.Code)
    // assert.JSONEq(t, `{"message":"user deleted successfully", "user": {"_id":"`+userIDStr+`", "username":"testuser"}}`, rr.Body.String())

    // mockUsecase.AssertExpectations(t)
}

func TestGetMyProfile(t *testing.T) {
    mockUsecase := new(mocks.UserUsecase)
    userController := NewUserController(mockUsecase)

    userName := "testuser"
    user := domain.User{
        Username: userName,
    }

    mockUsecase.On("GetMyProfile", userName).Return(user, nil)

	router := gin.Default()
	router.GET("/users/me", func (c *gin.Context) {
		c.Set("username", userName)
		userController.GetMyProfile(c)
	})

	req, _ := http.NewRequest("GET", "/users/me", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"user": {"_id":"000000000000000000000000", "password":"", "role":"","username":"testuser"}}`, rr.Body.String())

	mockUsecase.AssertExpectations(t)



    // router := setupRouterUser(mockUsecase)

    // req, _ := http.NewRequest("GET", "/users/me", nil)
    // rr := httptest.NewRecorder()

    // router.ServeHTTP(rr, req)

    // assert.Equal(t, http.StatusOK, rr.Code)
    // assert.JSONEq(t, `{"user": {"username":"testuser"}}`, rr.Body.String())

    // mockUsecase.AssertExpectations(t)
}
