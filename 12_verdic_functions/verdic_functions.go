package main

import "fmt"

func sum(nums...int) int{
	total:=0

	for _,num:= range nums{
		total = total + num
	}
	return total
}




func main(){
	nums:= []int {1,2,3,4,5}
	sum(nums...)
	result:=sum(1,2,3,4,5)
	fmt.Println(result)
}