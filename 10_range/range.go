package main

import (
	"fmt"
)

//Used for itirating over data structure
func main(){
	//nums:= []int{6,7,8}
	// for i:=0; i<len(nums); i++{
	// 	fmt.Println(nums[i])
	// }
	// sum:=0
	// for _,num := range nums{
	// 	sum = num+sum
	// }
	// fmt.Println(sum)

	//i is index
	// for i,num := range nums{
	// 	fmt.Println("index:",i ,"value:",num)
	// }


	//If for map
	// nums:= map[string]string{"name":"akshay","age":"22"}
	//i is key , nums is value
	// for i,num := range nums{
	// 	fmt.Println("key",i,"value",num)
	// }


	//We can use it for a string
	for i ,v:= range "golang"{
		//fmt.Println(i,v)
		fmt.Println(i,string(v))
	}
	//unicode code
	// 0 103
	// 1 111
	// 2 108
	// 3 97
	// 4 110
	// 5 103
	
}