package main

import "fmt"


const age = 20
func main(){
	//const name  = "String"
	fmt.Println(age)

	//Mutliple Values

	const (
		port = 400
		host = "local host"
	)
	fmt.Println(port,host)

}