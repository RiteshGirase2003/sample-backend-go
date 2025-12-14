// repository/product_repository.go
package repository

import (
	"account-guru/database"
	"account-guru/entity"
)

func CreateProduct(product *entity.Product) error {
	return database.DB.Create(product).Error
}

func GetProductByID(id uint) (*entity.Product, error) {
	var product entity.Product
	err := database.DB.First(&product, id).Error
	return &product, err
}

func GetAllProducts() ([]entity.Product, error) {
	var products []entity.Product
	err := database.DB.Find(&products).Error
	return products, err
}

func UpdateProduct(product *entity.Product) error {
	return database.DB.Save(product).Error
}

func DeleteProduct(id uint) error {
	return database.DB.Delete(&entity.Product{}, id).Error
}
