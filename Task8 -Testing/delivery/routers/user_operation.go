package routers

import (
	"clean_architecture_Testing/Delivery/controllers"
	"clean_architecture_Testing/config/db"
	"clean_architecture_Testing/infrastracture"
	"clean_architecture_Testing/repository"
	"clean_architecture_Testing/usecase"

	"github.com/gin-gonic/gin"
)

func SetupUserOperationRoutes(router *gin.Engine) {
	userRepo := repository.NewUserRepositoryImpl(db.UserCollection)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userController := controllers.NewUserController(userUsecase)
	userRoutes := router.Group("/")
	userRoutes.Use(infrastracture.AuthMiddleware())
	{
		userRoutes.GET("/users",infrastracture.RoleMiddleware("admin") , userController.GetAllUsers)
		userRoutes.DELETE("/users/:id" , userController.DeleteUserID)
		userRoutes.GET("/users/:id", userController.GetUserByID)
		userRoutes.GET("/users/me", userController.GetMyProfile)

	}

}