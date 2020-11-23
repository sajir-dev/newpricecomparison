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

// GetProductDetails ...
func (a *Amazon) GetItem(itemName string) (string, error) {
	productData, err := services.GetProductDetails(itemName)
	return productData, err
}
