package controllers

import (
	"jwt_task_manegnment/middleware"
	"jwt_task_manegnment/model"
	"jwt_task_manegnment/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user model.User
	var token string
	var Role_ string
	
	c.BindJSON(&user)
	Role_,err := services.Login(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	user.Role = Role_
	token, err = middleware.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}


func Register(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	err := services.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user created"})
}