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

//Contructor in Go
// func newOrder(id int,status string) *orders{
// 	myOrders:=orders{
// 		id : id,
// 		status: status,
// 		createdAt: time.Now(),
// 	}
// 	return &myOrders
// }

//If we want to add methods to struct
// func(o *orders) changeStatus(status string){
// 	//here we have not used * because derefrence done automatically
// 	o.status = status
// }

// func (o *orders) changeId(id int){
// 	o.id = id
// }


func main(){

	language := struct{
		name string
		isGood bool
	}{
		name: "Akshay",
		isGood: true,
	}

	fmt.Println(language)








	// myOrder:=newOrder(
	// 	13,
	// 	"NotReceived",
	// )

	// myOrder.changeId(11)
	// fmt.Println(myOrder)

	


	//If u dont set any filed it sets to 0
	//  myOrders:=orders{
	//  	id : 1,
	//  	amount: 100,
	//  	status: "received",
	//  	createdAt: time.Now(),
	//  }
	//  fmt.Println(myOrders)

	//  myOrders.changeStatus("NotReceived")
	//  fmt.Println(myOrders.status)
	//  myOrders.changeId(100)
	//  fmt.Println(myOrders.id)
	// myOrders.id = 3
	// fmt.Println(myOrders.id)
	// fmt.Println(myOrders.amount)
	// fmt.Println(myOrders)



}