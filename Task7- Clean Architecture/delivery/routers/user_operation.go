package routers

import (
	"clean_architecture/Delivery/controllers"
	"clean_architecture/bootstrap/db"
	"clean_architecture/infrastracture"
	"clean_architecture/repository"
	"clean_architecture/usecase"

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