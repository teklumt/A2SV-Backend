package main

import (
	"task_manager/routes"
)

func main(){
	routes.IntializeRoutes()
	routes.Router.Run(":8080")
}