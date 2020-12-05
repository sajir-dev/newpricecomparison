package main

import (
	"fmt"

	"./app"
)

func main() {
	app.StartApp()
	fmt.Println("app starting...")
}

// TODO -
// #1 db error to be given as internal server error
// #2 db to initialize at the starting of the app
// #3 objects to init at the starting of the app
// #4 learn- variable scope and garbage collection in go
// #5 add req struct api to dao
// #6 marketplace init at marketplace.go
// #7 c.Keys("from_data") in controllers
// #8 learn memory allocation in go
// #9

// #1 write a go script to add 100k products into the table with keys - itemid(unique), itemname, category, price(random), weight(random), qty
// #2 api for 4 category total weight, total price, avg  price (getCategoryPriceWeight) - CategoryName(list)
// #3 use channels for better UX (streaming function in gin)
// #4 auto migrate-(gorm, sqlc)

// #5 db-indexing

// #6 apply interface
// #1 make four apis into one
// #2 add marketplace in every header
// #3 calculate the
// #4 create model for req and res
// #5 api models in a file

// type CatRequest struct {
// 	Categories []string
// 	Username   string
// }

// type CatResponse struct {
// 	CategoryInfo []CategoryData
// 	Username string
// }

// type CategoryData struct {
// 	// av wt, name, price
// }


