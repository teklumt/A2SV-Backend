package routes

import (
	"jwt_task_manegnment/controllers"
	"jwt_task_manegnment/middleware"

	"github.com/gin-gonic/gin"
)

var Router  = gin.Default()

func IntializeRoutes() {

	publicRoutes := Router.Group("/auth") 
	{
		publicRoutes.POST("login",controllers.Login)
		publicRoutes.POST("register", controllers.Register)

	}

	protected := Router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{		
		protected.GET("tasks", controllers.GetAllTasksForUser)
		protected.GET("AllTasks",middleware.RoleMiddleware()  , controllers.GetAllTasks)
		protected.GET("tasks/:id",   controllers.GetSpecificTask)
		protected.PUT("tasks/:id" , controllers.UpdateSpecificTask)
		protected.DELETE("tasks/:id", controllers.DeleteSpecificTask)
		protected.POST("tasks", controllers.AddSpecificTask)

	}



	Router.NoRoute(func(c *gin.Context) {
		c.IndentedJSON(404, gin.H{"message": "Route Not Found"})
	})
}

