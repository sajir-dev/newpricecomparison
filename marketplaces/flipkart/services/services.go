package services

import (
	"../domain"
)

// GetProductDetails ...
func GetProductDetails(itemName string) (interface{}, error) {
	var fp *domain.ProductData
	productData, err := fp.GetProduct(itemName)
	return productData, err
}
