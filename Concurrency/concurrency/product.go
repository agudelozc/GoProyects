package main

import "strconv"

type Product struct {
	Name, Category string
	Price          float64
}

var ProductList = []*Product{
	{"Laptop", "Electronics", 1000},
	{"Milk", "Food", 2},
	{"TV", "Electronics", 500},
	{"Bread", "Food", 1},
	{"Mouse", "Electronics", 20},
	{"Orange", "Food", 3},
	{"Keyboard", "Electronics", 50},
}

type ProductGroup []*Product

type ProductData = map[string]ProductGroup

var Products = make(ProductData)

func ToCurrency(val float64) string {
	return "$" + strconv.FormatFloat(val, 'f', 2, 64)
}
func init() {
	for _, p := range ProductList {
		if _, ok := Products[p.Category]; ok {
			Products[p.Category] = append(Products[p.Category], p)
		} else {
			Products[p.Category] = ProductGroup{p}
		}
	}
}
