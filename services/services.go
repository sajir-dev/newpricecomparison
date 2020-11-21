package services

import (
	"encoding/json"

	amazon "../marketplaces/amazon/services"
	flipkart "../marketplaces/flipkart/services"
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
func GetItem(itemname string, marketplace string) (ItemInterface, error) {
	var i ItemInterface
	switch marketplace {
	case "amazon":
		i = amazon.CreateItem()
	case "flipkart":
		i = flipkart.CreateItem()
	}
	// a1 := amazon.CreateItem()
	// f1 := flipkart.CreateItem()
	dataString, err := i.GetItem(itemname)
	if err != nil {
		return nil, err
	}

	// var data *ItemInterface

	err = json.Unmarshal([]byte(dataString), &i)
	if err != nil {
		return nil, err
	}

	return i, nil
}
