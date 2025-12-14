// internal/routes/pingRoute.go
package routes

import (
	"account-guru/internal/controller"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.RouterGroup) {
	products := router.Group("/products")
	{
		products.POST("", controller.CreateProduct)
		products.GET("", controller.GetAllProducts)
		products.GET("/:id", controller.GetProduct)
		products.PUT("/:id", controller.UpdateProduct)
		products.DELETE("/:id", controller.DeleteProduct)
	}
}
