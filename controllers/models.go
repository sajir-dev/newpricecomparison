package controllers

type CatRequest struct {
	Categories []string `json:"categories"`
	Username   string   `json:"username"`
}

type CatResponse struct {
	CategoryInfo []CategoryData `json:"category_data"`
	Username     string         `json:"username"`
}

type CategoryData struct {
	CategoryName string  `json:"category_name"`
	AvgPrice     float64 `json:"avg_price"`
	TotalWeight  float64 `json:"total_weight"`
	TotalPrice   float64 `json:"total_price"`
}
