package main

import (
	"fmt"
	"sync"
)

func task(id int,w *sync.WaitGroup){
	defer w.Done() //defer executes when fun work is done
	fmt.Println("Done",id)
}


func main(){
	var wg sync.WaitGroup
	for i:=0; i<=10; i++{
		wg.Add(1) //1 goroutine started we use ADD
		go task(i,&wg)
		// wg.Add(1)
		// go task(i,&wg)

		// or 
		// wg.Add(2) if we have more than 2 go we update Add()
		// go task(i,&wg)
		// go task(i,&wg)
	}
	wg.Wait()
}