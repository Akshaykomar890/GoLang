package main

import "fmt"

//Closure fun
func incriment() func () int{
	var counter int = 0

	return func() int{
		counter+=1
		return counter
	}
}


func main(){

	result := incriment()

	fmt.Println(result())
}