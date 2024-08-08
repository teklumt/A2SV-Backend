// Delivery/routers/router.go
package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    SetupUserRegisterRoutes(router)
    // SetupTaskRoutes(router)


    return router
}

