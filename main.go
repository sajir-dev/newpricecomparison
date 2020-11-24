package main

import (
	"fmt"

	"./app"
	marketplace "./marketplaces"
)

func main() {
	// app.StartApp()
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

func init() {
	mysqlconfig.Init()
	marketplace.InitializeClient()
	app.StartApp()
}
