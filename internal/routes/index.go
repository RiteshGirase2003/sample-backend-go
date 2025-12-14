package routes

import (
	"account-guru/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// routes sets up the routes for the application
func InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,                  // allow all origins
		AllowMethods:     []string{"*"},         // allow all HTTP methods
		AllowHeaders:     []string{"*"},         // allow all headers
		ExposeHeaders:    []string{"*"},         // expose all headers
		AllowCredentials: true,                  // allow cookies/credentials
		MaxAge:           12 * time.Hour,       // preflight cache duration
	}))

	routerGroup := router.Group("/" + config.Config.String("context-path"))


	PingRoute(routerGroup)
	ProductRoutes(routerGroup)
	

	return router
}
