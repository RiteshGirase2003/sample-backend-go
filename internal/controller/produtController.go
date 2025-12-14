// controller/product_controller.go
package controller

import (
	"net/http"
	"strconv"

	"account-guru/entity"
	"account-guru/internal/service"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product entity.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := service.CreateProductService(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := service.GetProductService(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func GetAllProducts(c *gin.Context) {
	products, err := service.GetAllProductsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	updated, err := service.UpdateProductService(uint(id), &product)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := service.DeleteProductService(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}
