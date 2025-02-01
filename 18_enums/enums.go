package main

import "fmt"

//Enumerated Types

// func changeOrder(status string){
// 	fmt.Println("change to",status)
// }

// func main(){
// 	changeOrder("received")
// }

//To Define a Enums using a Int
// type Orders int

// const (
// 	Received Orders = iota
// 	Updated
// 	Confirmed
// )

// func getStatus(status Orders){
// 	fmt.Println("Order Status",status)
// }

// func main(){
// 	getStatus(Received)
// 	getStatus(Confirmed)
// }

//To Define a Enums using a String Values
//It helps in readability ,maintains a code structure


type Orders string

const (
	Received Orders  = "Received"
	Updated Orders   = "Updated"
	Confirmed Orders = "Confirmed"
)

func getStatus(status Orders){
	fmt.Println("Order Status",status)
}

func main(){
	getStatus(Received)
	getStatus(Confirmed)
}