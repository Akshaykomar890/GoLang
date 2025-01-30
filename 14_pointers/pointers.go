package main

import "fmt"

// func changeNum(num int){
// 	//num = 1
// 	//but here we change num to 5
// 	num = 5
// 	fmt.Println("Before Change",num)
// }

//we can not change main num to do we use
func byReference(num *int){
	//deReference
	*num = 5
	fmt.Println("Before Change",*num)
}


func main(){
	num:=1
	//changeNum(num)
	//fmt.Println("Memory address",&num)
	byReference(&num)
	fmt.Println("After Change",num)
}