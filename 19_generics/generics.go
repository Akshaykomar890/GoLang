package main

import (
	"fmt"
)

// func printSlice(items []int){
// 	for _,item := range items{
// 		fmt.Println(item)
// 	}
// }

//Problem duplicate
// func printSliceString(items []string){
// 	for _,item := range items{
// 		fmt.Println(item)
// 	}
// }

//To solve  duplicate we use generics
//any or interface{}
// func printSlice[T int | string | bool](items [] T){
// 	for _,item:= range items{
// 		fmt.Println(item)
// 	}
// }

// //Or

// func print[T any](print T){
// 	fmt.Println("Print",print)
// }



type stack[T any] struct{
	elements []T
}


func main(){

	result:=stack[int]{
		elements: []int{1,2,3,4},
	}

	fmt.Println(result)



	// nums:=[]int{1,2,3}
	// names:=[]string{"golang","kotlin"}
	// yn:=[]bool{true,false,true,true}
	// //printSliceString(names)
	// printSlice(names)
	// printSlice(nums)
	// printSlice(yn)

	// //normal usage
	// print("Akshay")
	// print(1)
	// print(true)
}