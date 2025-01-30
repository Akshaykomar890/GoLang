package main

import (
	"fmt"
	"time"
)

type orders struct{
	id int
	amount int
	status string
	createdAt time.Time
}


func main(){
	myOrders:=orders{
		id : 1,
		amount: 100,
		status: "received",
		createdAt: time.Now(),
	}
	fmt.Println(myOrders.id)
	fmt.Println(myOrders.amount)
	fmt.Println(myOrders)
}