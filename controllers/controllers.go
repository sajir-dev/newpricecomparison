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

// formRequest := c.Keys["form_data"].(*dao.modelname) used to get a

// GetItem ...
func GetItem(c *gin.Context) {
	// fmt.Println("itemname:", req)
	mp := c.Request.Header.Get("marketplace")
	// fmt.Println("marketplace", mp)

	// req = ItemReq{req.ItemName, mp}
	formRequest, _ := c.GetPostForm("item_name")
	fmt.Println(formRequest)
	// check more about c.Keys
	fmt.Println(formRequest)
	if len(mp) == 0 || len(formRequest) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

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
