package controllers

import (
	"net/http"

	"../services"
	"github.com/gin-gonic/gin"
)

// ItemReq ...
type ItemReq struct {
	ItemName    string `json:"item_name"`
	MarketPlace string `json:"marketplace"`
}

// formRequest := c.Keys["form_data"].(*dao.modelname) used to get a

// GetItem ...
func GetItem(c *gin.Context) {
	// fmt.Println("itemname:", req)
	mp := c.Request.Header.Get("marketplace")
	// fmt.Println("marketplace", mp)

	// req = ItemReq{req.ItemName, mp}
	formRequest := c.Keys["form_data"].(*ItemReq)
	// check more about c.Keys
	// fmt.Println(req)
	// if len(req.ItemName) == 0 || len(req.MarketPlace) == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Bad request",
	// 	})
	// 	return
	// }

	// itemData, err := services.GetItem(req.ItemName, req.MarketPlace)
	itemData, err := services.GetItem(formRequest, mp)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, itemData)
}
