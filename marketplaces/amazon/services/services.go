package services

import (
	"../domain"
)

// GetProductDetails ...
func GetProductDetails(itemName string) (string, error) {
	var fp *domain.ProductData
	productData, err := fp.GetItem(itemName)
	return productData, err
}
