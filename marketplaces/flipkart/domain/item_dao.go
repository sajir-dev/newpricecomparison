package domain

import (
	config "../../../utils"
)

// ProductData ...
type ProductData struct {
	ItemID      string  `json:"itemid"`
	ItemName    string  `json:"item"`
	Price       float32 `json:"price"`
	Brand       string  `json:"brand"`
	Description string  `json:"description"`
	Rating      string  `json:"rating"`
	MarketPlace string  `json:"platform"`
}

// GetItem returns a specific item if available
func (p *ProductData) GetProduct(itemname string) (*ProductData, error) {
	// row := config.DB.QueryRow(`select * from items join marketplace using(itemid) where (marketplace = 'flipkart' and name = '` + itemname + `');`)
	row := config.DB.QueryRow(`select * from products where (marketplace = 'flipkart' and itemname = '` + itemname + `');`)
	var item *ProductData = new(ProductData)
	err := row.Scan(&item.ItemID, &item.ItemName, &item.Price, &item.Brand, &item.Description, &item.Rating, &item.MarketPlace)
	if err != nil {
		// return "", err
		return nil, err
	}

	// itemJSON, _ := json.Marshal(item)
	// return string(itemJSON), nil
	return item, nil
}

// func (p *ProductData) CreateProduct()
