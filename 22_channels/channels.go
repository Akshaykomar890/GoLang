package main

import (
	"fmt"
)

//Send data
// func processNum(numChan chan int){
// 	for num:= range numChan{
// 		fmt.Println("Processing",num)
// 		time.Sleep(time.Second)
// 	}
// }

//Receive
// func sum(values chan int,num1 int,num2 int){
// 	result := num1+num2
// 	//send data to value
// 	values <- result //blocking
// }

//How to use wait group in channels
func task(done chan bool){
	defer func () { done <-true}()

	fmt.Println("Procssing....")
}


func main(){

	//How to use wait group in channels
	done:=make(chan bool)
	go task(done)
	<-done //block









	//Channels Communication Done
	//Receive
	// resultChan:=make(chan int) //make channel
	// go sum(resultChan,1,2) //run on non blocking
	// ans:= <-resultChan //receive <-blocking
	// fmt.Println(ans)


	//Send data
	// numChan:= make(chan int) //channel create
	// go processNum(numChan) //non blocking //added go to run on seperate goroutine
	// for {
	// 	numChan <- rand.Intn(100) //send values
	// }




	//Start example
	// messageChan:= make(chan string) //Created a channel
	// messageChan <- "ping" //Send //Blocking
	// msg := <- messageChan //Receive
	// fmt.Println(msg)





}