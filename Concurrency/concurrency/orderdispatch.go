package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DispatchNotification struct {
	Customer string
	*Product
	Quantity int
}

var Customers = []string{"Bob", "Alice", "Joe", "Doe", "Smith"}

func DispatchOrder(channel chan<- DispatchNotification) {
	rand.Seed(time.Now().UnixNano())
	orderCount := rand.Intn(3) + 2
	fmt.Println("Order count:", orderCount)
	for i := 0; i < orderCount; i++ {
		channel <- DispatchNotification{
			Customer: Customers[rand.Intn(len(Customers)-1)],
			Quantity: rand.Intn(10),
			Product:  ProductList[rand.Intn(len(ProductList)-1)],
		}
		//if i == 1{
		//	notification := <- channel
		//	fmt.Println("Read:", notification.Customer)
		//}
	}
	close(channel)
}
