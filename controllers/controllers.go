package controllers

import (
	"fmt"
	"net/http"

	"../services"
	"github.com/gin-gonic/gin"
)

// ItemReq ...
type ItemReq struct {
	ItemName    string `json:"item_name"`
	MarketPlace string `json:"marketplace"`
}

// GetItem ...
func GetItem(c *gin.Context) {
	var req ItemReq
	err := c.ShouldBindJSON(&req)
	fmt.Println("itemname:", req)
	mp := c.Request.Header.Get("marketplace")
	fmt.Println("marketplace", mp)
	req = ItemReq{req.ItemName, mp}
	fmt.Println(req)
	if len(req.ItemName) == 0 || len(req.MarketPlace) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	// itemData, err := services.GetItem(req.ItemName, req.MarketPlace)
	itemData, err := services.GetItem(req.ItemName, req.MarketPlace)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, itemData)
}
