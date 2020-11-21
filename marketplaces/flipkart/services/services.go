package flipkart

import (
	"encoding/json"

	"../../../config"
)

// FlipkartItem ...
type FlipkartItem struct {
	ItemID      string  `json:"itemid"`
	ItemName    string  `json:"item"`
	Price       float32 `json:"price"`
	Brand       string  `json:"brand"`
	Description string  `json:"description"`
	Rating      string  `json:"rating"`
	MarketPlace string  `json:"platform"`
}

// GetItem returns a specific item if available
func (a *FlipkartItem) GetItem(itemname string) (string, error) {
	row := config.DB.QueryRow(`select * from items join marketplace using(itemid) where (marketplace = 'flipkart' and name = '` + itemname + `');`)

	var item FlipkartItem
	err := row.Scan(&item.ItemID, &item.ItemName, &item.Price, &item.Brand, &item.Description, &item.Rating, &item.MarketPlace)
	if err != nil {
		return "", err
	}

	itemJSON, _ := json.Marshal(item)
	return string(itemJSON), nil
}

// CreateItem ...
func CreateItem() *FlipkartItem {
	return &FlipkartItem{MarketPlace: "flipkart"}
}
