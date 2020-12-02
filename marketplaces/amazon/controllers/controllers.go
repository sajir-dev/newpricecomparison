package amazon

import (
	"../services"
)

// Amazon ...
type Amazon struct {
	MarketPlace string `json:"marketplace"`
}

// CreateItem ...
func InitAmazon() *Amazon {
	return &Amazon{MarketPlace: "amazon"}
}

// GetProduct ...
func (a *Amazon) GetProduct(itemName string) (interface{}, error) {
	productData, err := services.GetProductDetails(itemName)
	return productData, err
}
