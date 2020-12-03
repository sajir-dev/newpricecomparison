package app

import (
	"../controllers"
	marketplace "../marketplaces"
	psqlconfig "../utils/psql"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hey I'm working")
	})
	router.POST("/", controllers.GetItem)
	router.GET("/catweight", controllers.GetCategoryWeight)
	router.GET("/catprice", controllers.GetCategoryPrice)
	router.GET("/catavg", controllers.GetCategoryAvg)
	router.GET("/listcat", controllers.ListCategories)

	router.Run(":8030")
}

func init() {
	psqlconfig.Init()
	marketplace.InitializeClient()
}
