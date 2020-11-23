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

// GetItem ...
func (f *Flipkart) GetItem(itemName string) (interface{}, error) {
	productData, err := services.GetProductDetails(itemName)
	return productData, err
}
