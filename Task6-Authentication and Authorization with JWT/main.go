package main

import (
	"jwt_task_manegnment/config"
	"jwt_task_manegnment/db"
	"jwt_task_manegnment/routes"
)

func main() {
	db.ConnectDB()
	config.InitiEnvConfigs() 
	routes.IntializeRoutes()

	routes.Router.Run(config.EnvConfigs.LocalServerPort)


	 
} 