package services

import (
	"fmt"

	amazon "../marketplaces/amazon/controllers"
	flipkart "../marketplaces/flipkart/controllers"
)

// TODO: change the name to MarketPlaceClient
type ItemInterface interface {
	GetItem(string) (string, error)
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

func GetItem(itemname string, marketplace string) (string, error) {
	// var i ItemInterface
	// switch marketplace {
	// case "amazon":
	// 	i = amazon.CreateItem()
	// case "flipkart":
	// 	i = flipkart.CreateItem()
	// }
	// a1 := amazon.CreateItem()
	// f1 := flipkart.CreateItem()

	marketPlaceObj := map[string]ItemInterface{"amazon": amazon.CreateItem(), "flipkart": flipkart.CreateItem()}

	dataString, err := marketPlaceObj[marketplace].GetItem(itemname)
	if err != nil {
		return "", err
	}

	// var data *ItemInterface

	fmt.Println(dataString)
	// err = json.Unmarshal([]byte(dataString), &i)
	// if err != nil {
	// 	return nil, err
	// }

	return dataString, nil
}

// POST: GetProductDetails(itemid), HEADER: marketplace define struct for this id
// Struct to hold product response data the one i have now
// marketplace struct needs just a string field marketplace
