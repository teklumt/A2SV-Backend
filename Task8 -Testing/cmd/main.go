package main

import (
	"clean_architecture_Testing/config"
	"clean_architecture_Testing/config/db"
	"clean_architecture_Testing/delivery/routers"
)

func main() {
    config.InitiEnvConfigs() 
    db.ConnectDB(config.EnvConfigs.MongoURI)
    //printGreen color 



    
    router := routers.SetupRouter()

    router.Run(config.EnvConfigs.LocalServerPort)
}
