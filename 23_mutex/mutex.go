package main

import (
	"fmt"
	"sync"
)


type post struct{
	views int
	mu sync.Mutex //Add mutex create
}

func (p *post) increment(wg *sync.WaitGroup){
	defer func ()  {
		defer p.mu.Unlock() //unlock 
		wg.Done()
	}()
	//lock only do that line not whole function or body
	p.mu.Lock() //lock to complete task serial way so that value can not change
	p.views += 1
}


func main(){
	var wg sync.WaitGroup
	myPost:=post{
		views: 0,
	}
	for i:=0; i<100; i++{
		wg.Add(1)
		go myPost.increment(&wg)
	}
	wg.Wait()
	fmt.Println(myPost.views)
}