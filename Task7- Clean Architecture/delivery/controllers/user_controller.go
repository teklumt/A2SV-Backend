package controllers

import (
	"clean_architecture/domain"
	"clean_architecture/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
    UserUsecase *usecase.UserUsecase
}

func NewUserController(userUsecase *usecase.UserUsecase) *UserController {
    return &UserController{UserUsecase: userUsecase}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
    var user domain.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := uc.UserUsecase.RegisterUser(user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})
}

func (uc *UserController) LoginUser(c *gin.Context) {
    var user domain.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := uc.UserUsecase.LoginUser(user.Username, user.Password)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "user logged in successfully", "user": user})
}