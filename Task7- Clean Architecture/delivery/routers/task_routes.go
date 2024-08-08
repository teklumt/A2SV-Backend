package routers

import (
	"clean_architecture/Delivery/controllers"
	"clean_architecture/bootstrap/db"
	"clean_architecture/infrastracture"
	"clean_architecture/repository"
	"clean_architecture/usecase"

	"github.com/gin-gonic/gin"
)

func SetupTaskRoutes(router *gin.Engine) {
	taskRepo := repository.NewTaskRepositoryImpl(db.TaskCollection)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	taskController := controllers.NewTaskController(taskUsecase)
	taskRoutes := router.Group("/tasks")
	taskRoutes.Use(infrastracture.AuthMiddleware())
	{
		taskRoutes.POST("/", taskController.CreateTask)
		taskRoutes.GET("/",infrastracture.RoleMiddleware("admin") , taskController.GetTasks)
		taskRoutes.GET("/:id", taskController.GetTaskByID)
		taskRoutes.GET("/me", taskController.GetMyTasks)
		taskRoutes.DELETE("/:id", taskController.DeleteTask)
		taskRoutes.PUT("/:id", taskController.UpdateTask)
	}
}