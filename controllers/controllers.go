package controllers

import (
	"fmt"
	"io"
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

// GetCategoryWeight ...
func GetCategoryWeight(c *gin.Context) {
	category, isNot := c.GetPostForm("category")
	if !isNot {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Bad Request"})
		return
	}
	wt, err := services.GetCategoryWeight(category)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err})
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"total weight": wt})
	return
}

// GetCategoryPrice ...
func GetCategoryPrice(c *gin.Context) {
	category, isNot := c.GetPostForm("category")
	if !isNot {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Bad Request"})
		return
	}
	price, err := services.GetCategoryPrice(category)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err})
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"total price": price})
	return
}

// GetCategoryAvg ...
func GetCategoryAvg(c *gin.Context) {
	category, isNot := c.GetPostForm("category")
	if !isNot {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Bad Request"})
		return
	}
	price, err := services.GetCategoryAvg(category)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err})
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"category avg price": price})
	return
}

// ListCategories ...
func ListCategories(c *gin.Context) {
	cs := services.ListCategories()

	c.Stream(func(w io.Writer) bool {
		if channel, ok := <-cs; ok {
			c.JSON(http.StatusOK, channel)
			return true
		}
		return false
	})
}
