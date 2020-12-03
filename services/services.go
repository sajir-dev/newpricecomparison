package services

import (
	"fmt"

	"../domain"
	marketplace "../marketplaces"
)

// TODO: change the name to MarketPlaceClient
type ItemInterface interface {
	GetProduct(string) (interface{}, error)
}

// TODO: no definitions here
// TODO: call initialise fn amazon, flipkart etc. object specific struct here. (no definition) and store the corresponding objects in a hashMap
// key: market place, value: initialised struct
//var (
//	Market map[string]ItemInterface
//)
//func init{
//Market = make(map[string]ItemInterface)
// Market["amazon"] = amzon.initia()
// Market["flipkart"] = flip.initia()

//}

// func init() {
// 	MarketPlaceObjects := Map["string"]{}
// }

// TODO: accept the type from controller

func GetItem(itemname string, mp string) (interface{}, error) {
	var i interface{}
	// switch marketplace {
	// case "amazon":
	// 	i = amazon.CreateItem()
	// case "flipkart":
	// 	i = flipkart.CreateItem()
	// }
	// a1 := amazon.CreateItem()
	// f1 := flipkart.CreateItem()

	mpObj := marketplace.GetMarketPlace(mp)

	// marketPlaceObj := map[string]ItemInterface{"amazon": amazon.CreateItem(), "flipkart": flipkart.CreateItem()}

	i, err := mpObj.GetProduct(itemname)
	if err != nil {
		return "", err
	}

	// var data *ItemInterface

	fmt.Println(i)
	// err = json.Unmarshal([]byte(dataString), &i)
	// if err != nil {
	// 	return nil, err
	// }

	return i, nil
}

// POST: GetProductDetails(itemid), HEADER: marketplace define struct for this id
// Struct to hold product response data the one i have now
// marketplace struct needs just a string field marketplace

// GetCategoryWeight ...
func GetCategoryWeight(category string) (float64, error) {
	wt, err := domain.GetTotalWeightOfTheCategory(category)
	return wt, err
}

// GetCategoryPrice ...
func GetCategoryPrice(category string) (float64, error) {
	total, err := domain.GetTotalPriceOfTheCategory(category)
	return total, err
}

// GetCategoryAvg ...
func GetCategoryAvg(category string) (float64, error) {
	total, err := domain.AvgPriceOfTheCategory(category)
	return total, err
}

// ListCategories ...
func ListCategories() chan string {
	c := domain.ListCategories()
	return c
}
