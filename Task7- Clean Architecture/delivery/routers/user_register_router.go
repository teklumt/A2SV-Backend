package routers

import (
	"clean_architecture/Delivery/controllers"
	"clean_architecture/bootstrap/db"
	"clean_architecture/repository"
	"clean_architecture/usecase"

	"github.com/gin-gonic/gin"
)

func SetupUserRegisterRoutes(router *gin.Engine) {
    userRepo := repository.NewUserRepositoryImpl(db.UserCollection)
    userUsecase := usecase.NewUserUsecase(userRepo)
    userController := controllers.NewUserController(userUsecase)

    userRoutes := router.Group("/users")
    {
        userRoutes.POST("/register", userController.RegisterUser)
        // Add more user-related endpoints here
    }
}
