package marketplace

import (
	amazon "./amazon/controllers"
	flipkart "./flipkart/controllers"
)

// TODO: change the name to MarketPlaceClient
type MarketPlaceClient interface {
	GetProduct(string) (interface{}, error)
}

var MarketPlaceObj map[string]MarketPlaceClient

func InitializeClient() {
	MarketPlaceObj = make(map[string]MarketPlaceClient)
	MarketPlaceObj["flipkart"] = flipkart.InitFlipkart()
	MarketPlaceObj["amazon"] = amazon.InitAmazon()
}

func GetMarketPlace(marketplace string) MarketPlaceClient {
	return MarketPlaceObj[marketplace]
}
