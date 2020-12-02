package flipkart

import (
	"../services"
)

// Flipkart ...
type Flipkart struct {
	MarketPlace string `json:"marketplace"`
}

// InitFlipkart ...
func InitFlipkart() *Flipkart {
	return &Flipkart{MarketPlace: "flipkart"}
}

// GetProduct ...
func (f *Flipkart) GetProduct(itemName string) (interface{}, error) {
	productData, err := services.GetProductDetails(itemName)
	return productData, err
}
