// Delivery/routers/router.go
package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    // user routes
    SetupUserRegisterRoutes(router)
    SetupUserLoginRoutes(router)
    SetupUserOperationRoutes(router)

    // task routes
    SetupTaskRoutes(router)
    return router
}

