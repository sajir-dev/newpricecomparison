package amazon

import (
	"../services"
)

// Amazon ...
type Amazon struct {
	MarketPlace string `json:"marketplace"`
}

// CreateItem ...
func CreateItem() *Amazon {
	return &Amazon{MarketPlace: "amazon"}
}

// GetItem ...
func (a *Amazon) GetItem(itemName string) (interface{}, error) {
	productData, err := services.GetProductDetails(itemName)
	return productData, err
}
