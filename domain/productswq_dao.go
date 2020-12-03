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
		q, _ := config.DB.Query(`select distinct category from products2;`)
		fmt.Println(q)
		for q.Next() {
			var c string
			q.Scan(&c)
			cs <- c
			// fmt.Println(<-cs)
		}
		close(cs)
	}()
	return cs
}
