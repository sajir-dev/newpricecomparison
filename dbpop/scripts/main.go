package main

import (
	"database/sql"
	"errors"
	"fmt"
	"math"
	"math/rand"

	_ "github.com/lib/pq"
)

type CategoryData struct {
	CategoryName string  `json:"category"`
	TotalQty     float64 `json:"quantity"`
	TotalWeight  float64 `json:"total_weight"`
	TotalPrice   float64 `json:"total_price"`
	AvgPrice     float64 `json:"avg_price"`
}

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
	// createItems()
	// fmt.Println(GetAvgCategoryPrice("cat10034"))
	// createItems2()

	// for v := range c {
	// 	fmt.Println(v)
	// }

	// c := ListCategories()
	// for v := range c {
	// 	fmt.Println(v)
	// 	fmt.Println(GetCategoryData(v))
	// }

	cd := GetCategoryInfo()
	for v := range cd {
		fmt.Println(v)
	}

	// fmt.Println(GetCategoryInfo(c))

}

func createItems() {
	for j := 10; j < 100000; j++ {
		rand.Seed(int64(j * 1000))
		itemname := `item100` + fmt.Sprint(rand.Intn(1000000))
		price := math.Floor(float64(rand.Float32()*100*100)) / 100
		category := `cat100` + fmt.Sprint(rand.Intn(100))
		weight := math.Floor(float64(rand.Float32()*15*100) / 100)
		quantity := rand.Intn(200)

		row, err := DB.Exec(`INSERT INTO products2(itemname, category, price, weight, quantity) values ($1, $2, $3, $4, $5);`, itemname, category, price, weight, quantity)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(row.RowsAffected())
	}
	return
}

func createItems2() {
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

func GetTotalWeightOfTheCategory(category string) (float64, error) {
	c := make(chan float64)
	var sum float64
	go func() {
		var weight float64
		q, err := DB.Query(`select weight from products2 where category = '` + category + `';`)
		if err != nil {
			return
		}
		for q.Next() {
			q.Scan(&weight)
			c <- weight
			// fmt.Println(<-c)
		}
		close(c)
	}()
	for v := range c {
		sum += v
	}
	if sum == 0 {
		return sum, errors.New("database error")
	}
	return sum, nil
}

func GetTotalPriceOfTheCategory(category string) (float64, error) {
	c := make(chan float64)
	var sum float64
	go func() {
		var weight float64
		q, err := DB.Query(`select price from products2 where category = '` + category + `';`)
		if err != nil {
			return
		}
		for q.Next() {
			q.Scan(&weight)
			c <- weight
			// fmt.Println(<-c)
		}
		close(c)
	}()
	for v := range c {
		sum += v
	}
	if sum == 0 {
		return sum, errors.New("database error")
	}
	return sum, nil
}

func GetTotalQtyOfTheCategory(category string) (float64, error) {
	c := make(chan float64)
	var sum float64
	go func() {
		var weight float64
		q, err := DB.Query(`select quantity from products2 where category = '` + category + `';`)
		if err != nil {
			return
		}
		for q.Next() {
			q.Scan(&weight)
			c <- weight
			// fmt.Println(<-c)
		}
		close(c)
	}()
	for v := range c {
		sum += v
	}
	if sum == 0 {
		return sum, errors.New("database error")
	}
	return sum, nil
}

func GetAvgCategoryPrice(category string) (float64, error) {
	qty, err := GetTotalQtyOfTheCategory(category)
	if err != nil {
		return 0, err
	}
	total, err := GetTotalPriceOfTheCategory(category)
	if err != nil {
		return 0, err
	}
	return (total / qty), nil
}

func AvgPriceOfTheCategory(category string) (float64, error) {
	total, err := GetTotalPriceOfTheCategory(category)
	if err != nil {
		return 0, err
	}
	qty, err := GetTotalQtyOfTheCategory(category)
	if err != nil {
		return 0, err
	}
	return (total / qty), nil
}

func ListCategories() chan string {
	cs := make(chan string)
	go func() {
		q, _ := DB.Query(`select distinct category from products2;`)
		fmt.Println(q)
		for q.Next() {
			var c string
			q.Scan(&c)
			cs <- c
			// time.Sleep(time.Second * 10)
			// fmt.Println(<-cs)
		}
		close(cs)
	}()
	return cs
}

func GetCategoryInfo() chan CategoryData {
	cd := make(chan CategoryData)
	c := ListCategories()
	// fmt.Println("1")
	go func() {
		for v := range c {
			// fmt.Println(v)
			cd <- GetCategoryData(v)
			// fmt.Println("3")
		}
		close(cd)
	}()
	return cd
}

func GetCategoryData(category string) CategoryData {
	var categoryData CategoryData
	weight, _ := GetTotalWeightOfTheCategory(category)
	qty, _ := GetTotalQtyOfTheCategory(category)
	avg, _ := AvgPriceOfTheCategory(category)
	price, _ := GetTotalPriceOfTheCategory(category)
	categoryData = CategoryData{category, weight, qty, price, avg}
	return categoryData
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

// CREATE TABLE "products2" (
// 	"id" bigserial PRIMARY KEY,
// 	"itemname" varchar NOT NULL,
// 	"category" varchar NOT NULL,
// 	"price" decimal,
// 	"weight" decimal,
// 	"quantity" decimal
// );
