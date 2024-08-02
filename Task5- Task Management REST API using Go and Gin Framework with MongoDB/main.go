package main

import (
	"task_manager_DB/config"
	"task_manager_DB/db"
	"task_manager_DB/routes"
)

func main() {
	db.ConnectDB()
	config.InitiEnvConfigs() 
	routes.IntializeRoutes()

	routes.Router.Run(config.EnvConfigs.LocalServerPort)


	 
} 