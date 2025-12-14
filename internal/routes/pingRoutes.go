// internal/routes/pingRoute.go
package routes

import (
	"account-guru/internal/controller"

	"github.com/gin-gonic/gin"
)

func PingRoute(router *gin.RouterGroup) {
	router.GET("/ping", controller.Ping)
}
