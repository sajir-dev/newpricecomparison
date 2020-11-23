package flipkart

import (
	"../services"
)

// Flipkart ...
type Flipkart struct {
	MarketPlace string `json:"marketplace"`
}

// CreateItem ...
func CreateItem() *Flipkart {
	return &Flipkart{MarketPlace: "flipkart"}
}

// GetProductDetails ...
func (f *Flipkart) GetItem(itemName string) (string, error) {
	productData, err := services.GetProductDetails(itemName)
	return productData, err
}
