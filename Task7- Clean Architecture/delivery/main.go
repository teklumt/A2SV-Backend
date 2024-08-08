package main

import (
	"clean_architecture/bootstrap/db"
	"clean_architecture/delivery/routers"
)

func main() {
    db.ConnectDB()
    router := routers.SetupRouter()
    router.Run(":8080")
}
