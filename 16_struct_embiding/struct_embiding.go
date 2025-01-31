package main

import "fmt"

//struct in a struct
func main(){
	//1st one
	// newOrder := orders{
	// 	1,
	// 	122,
	// 	"received",
	// 	customer{
	// 		"Akshay",
	// 		"454545",
	// 	},
	// }

	//or we can use 
	customer := customer{
		name: "akshay",
		phone: "12345",
	}
	newOrder := orders{
		1,
		122,
		"received",
		customer,
	}

	newOrder.customer.name = "boby"

	fmt.Println(newOrder.customer.name)


}

type orders struct{
	id int
	ammount int
	status string
	customer
}

type customer struct{
	name string
	phone string
}

