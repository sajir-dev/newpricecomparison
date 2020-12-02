package domain

import (
	psqlconfig "../../../utils/psql"
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
	// row := psqlconfig.DB.QueryRow(`select * from products join marketplace using(itemid) where (marketplace = 'amazon' and name = '` + itemname + `');`)
	row := psqlconfig.DB.QueryRow(`select * from products where (marketplace = 'amazon' and itemname = '` + itemname + `');`)

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
