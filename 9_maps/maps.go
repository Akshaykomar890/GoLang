package main

import "fmt"

func main(){
	//Maps in  Go
	//key-> String value ->String
	//n:= make(map[string]string)
	// n["name"] = "akshay"
	// n["backend"] = "area"
	//fmt.Println(n["name"],n["backend"])


	//key-> String value ->Int
	n:=make(map[string]int)
	n["age"] = 10
	fmt.Println(n["age"])
	fmt.Println(len(n))
	delete(n,"age")
	fmt.Println(n)

}