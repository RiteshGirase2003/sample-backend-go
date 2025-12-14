// service/product_service.go
package service

import (
	"errors"

	"account-guru/entity"
	"account-guru/internal/repository"
)

func CreateProductService(product *entity.Product) error {
	if product.Name == "" || product.Price <= 0 {
		return errors.New("invalid product data")
	}
	return repository.CreateProduct(product)
}

func GetProductService(id uint) (*entity.Product, error) {
	return repository.GetProductByID(id)
}

func GetAllProductsService() ([]entity.Product, error) {
	return repository.GetAllProducts()
}

func UpdateProductService(id uint, updated *entity.Product) (*entity.Product, error) {
	product, err := repository.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	product.Name = updated.Name
	product.Description = updated.Description
	product.Price = updated.Price
	product.Qty = updated.Qty

	return product, repository.UpdateProduct(product)
}

func DeleteProductService(id uint) error {
	return repository.DeleteProduct(id)
}
