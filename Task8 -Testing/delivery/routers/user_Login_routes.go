package routers

import (
	"clean_architecture_Testing/Delivery/controllers"
	"clean_architecture_Testing/config/db"
	"clean_architecture_Testing/repository"
	"clean_architecture_Testing/usecase"

	"github.com/gin-gonic/gin"
)

func SetupUserLoginRoutes(router *gin.Engine) {
	userRepo := repository.NewUserRepositoryImpl(db.UserCollection)
    userUsecase := usecase.NewUserUsecase(userRepo)
    userController := controllers.NewUserController(userUsecase)
	userRoutes := router.Group("/auth")
	{
		userRoutes.POST("/login", userController.LoginUser)
	}

}