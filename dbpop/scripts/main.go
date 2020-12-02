package main

import (
	"database/sql"
	"fmt"
	"math"
	"math/rand"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var err error

func init() {
	DB, err = sql.Open("postgres", "postgres://postgres:password@localhost/pricecomparisonv2?sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("You are connected to local postgres")
}

func main() {
	createItems()

	// for v := range c {
	// 	fmt.Println(v)
	// }

}

func createItems() {
	// INSERT INTO products(ITEMNAME, PRICE, BRAND, DESCRIPTION, RATING, MARKETPLACE) values ('demo2', 12.2, 'efaef', 'demo descriptionfewfcves', 4.5, 'amazon');
	// var c chan int
	for j := 1001; j < 100000; j++ {
		rand.Seed(int64(j * 1000))
		itemname := `item100` + fmt.Sprint(rand.Intn(1000000))
		price := math.Floor(float64(rand.Float32()*10*100)) / 100
		brand := `brand100` + fmt.Sprint(rand.Intn(40))
		description := `demo descrption` + fmt.Sprint(rand.Intn(100000))
		rating := math.Floor(float64(rand.Float32()*5*10)) / 10
		var arr [2]string
		arr[0] = `flipkart`
		arr[1] = `amazon`
		mp := arr[rand.Intn(2)]
		// fmt.Println(
		// 	itemname,
		// 	price,
		// 	brand,
		// 	description,
		// 	rating,
		// 	mp,
		// )
		row, err := DB.Exec(`INSERT INTO products(itemname, price, brand, description, rating, marketplace) values ($1, $2, $3, $4, $5, $6);`, itemname, price, brand, description, rating, mp)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(row.RowsAffected())
		// c <- j
	}
	// close(c)
	return
}

// ItemID, ItemName, Price, Brand, Description, Rating, MarketPlace

// CREATE TABLE "products" (
// 	"id" bigserial PRIMARY KEY,
// 	"itemname" varchar NOT NULL,
// 	"price" decimal NOT NULL,
// 	"brand" varchar,
// 	"description" varchar,
// 	"rating" decimal,
// 	"marketplace" varchar
// );
