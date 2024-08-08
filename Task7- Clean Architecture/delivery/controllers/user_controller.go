package controllers

import (
	"clean_architecture/domain"
	"clean_architecture/infrastracture"
	"clean_architecture/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
    fmt.Println(user,"**************")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, err := infrastracture.GenerateToken(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "error generating token"})
        return  
    }


    c.JSON(http.StatusOK, gin.H{"message": "user logged in successfully", "Token": token})
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
    users, err := uc.UserUsecase.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"users": users})
}

func (uc *UserController) DeleteUserID(c *gin.Context) {
    id := c.Param("id")
    Role_user := c.GetString("Role")

    if Role_user != "admin" {
        user, err := uc.UserUsecase.GetUserByID(id)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        newID, err := primitive.ObjectIDFromHex(id)
        if user.ID != newID {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }
    }




    user, err := uc.UserUsecase.DeleteUserID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully", "user": user})
}

func (uc *UserController) GetUserByID(c *gin.Context) {
    id := c.Param("id") 
    
    user, err := uc.UserUsecase.GetUserByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"user": user})
}

func (uc *UserController) GetMyProfile(c *gin.Context) {
    // Retrieve the user claims from the context
    userName := c.GetString("username")
    user, err := uc.UserUsecase.GetMyProfile(userName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"user": user})
}
