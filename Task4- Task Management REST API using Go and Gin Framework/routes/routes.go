package routes

import (
	"github.com/gin-gonic/gin"

	"task_manager/controllers"
)

var Router  = gin.Default()

func IntializeRoutes() {
	publicRoutes := Router.Group("/")
	publicRoutes.GET("tasks", controllers.GetAllTasks )
	publicRoutes.GET("tasks/:id", controllers.GetSpecificTask)
	publicRoutes.PUT("tasks/:id", controllers.UpdateSpecificTask)
	publicRoutes.DELETE("tasks/:id", controllers.DeleteSpecificTask)
	publicRoutes.POST("tasks", controllers.AddSpecificTask)


	Router.NoRoute(func(c *gin.Context) {
		c.IndentedJSON(404, gin.H{"message": "Route Not Found"})
	})
}

