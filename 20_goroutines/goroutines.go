package main

import (
	"fmt"
	"time"
)

func task(id int){
	fmt.Println("doing task",id)
}


func main(){

	for i:=0; i<=10; i++{
		//go task(i)

		//Inline function
		//	 Pass value
		 go func (i int)  {
			fmt.Println(i)
		}(i) //Receive value
	}

	time.Sleep(time.Second*2)

}