package domain

import (
	"errors"
	"fmt"

	config "../utils"
)

type ProductWQ struct {
	ProductId   int     `json:"id"`
	ProductName string  `json:"itemname"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Weight      float64 `json:"weight"`
	Quantity    int     `json:"quantity"`
}

type CategoryData struct {
	CategoryName string  `json:"category"`
	TotalQty     float64 `json:"quantity"`
	TotalWeight  float64 `json:"total_weight"`
	TotalPrice   float64 `json:"total_price"`
	AvgPrice     float64 `json:"avg_price"`
}

func GetTotalWeightOfTheCategory(category string) (float64, error) {
	c := make(chan float64)
	var sum float64
	go func() {
		var weight float64
		q, err := config.DB.Query(`select weight from products2 where category = '` + category + `';`)
		if err != nil {
			return
		}
		for q.Next() {
			q.Scan(&weight)
			c <- weight
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
		var price float64
		q, err := config.DB.Query(`select price from products2 where category = '` + category + `';`)
		if err != nil {
			return
		}
		for q.Next() {
			q.Scan(&price)
			c <- price
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
		var qty float64
		q, err := config.DB.Query(`select quantity from products2 where category = '` + category + `';`)
		if err != nil {
			return
		}
		for q.Next() {
			q.Scan(&qty)
			c <- qty
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
		q, _ := config.DB.Query(`select distinct itemname from products2;`)
		fmt.Println(q)
		for q.Next() {
			var c string
			q.Scan(&c)
			cs <- c
			// time.Sleep(time.Millisecond)
			// fmt.Println(<-cs)
		}
		close(cs)
	}()
	// go test(cs)
	return cs
}

// func test(cs chan string){
// 	q, _ := config.DB.Query(`select distinct category from products2;`)
// 	fmt.Println(q)
// 	for q.Next() {
// 		var c string
// 		q.Scan(&c)
// 		cs <- c
// 		// fmt.Println(<-cs)
// 	}
// 	close(cs)
// }

func GetCategoryInfo(category chan string) chan CategoryData {
	var c chan CategoryData
	for v := range category {
		c <- GetCategoryData(v)
	}
	close(c)
	return c
}

func GetCategoryData(category string) CategoryData {
	var categoryData CategoryData
	weight, _ := GetTotalWeightOfTheCategory(category)
	qty, _ := GetTotalQtyOfTheCategory(category)
	avg, _ := AvgPriceOfTheCategory(category)
	price, _ := GetTotalPriceOfTheCategory(category)

	// if err != nil {
	// 	return nil, err
	// }

	categoryData = CategoryData{category, weight, qty, price, avg}
	return categoryData
}
