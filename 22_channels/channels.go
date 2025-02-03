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

//How to use wait group in channels unbufferd
// func task(done chan bool){
// 	defer func () { done <-true}()

// 	fmt.Println("Procssing....")
// }

//Single Chan
// func task(emailChan chan string,done chan bool){
// 	defer func ()  { done <- true}()
// 	for email:= range emailChan{
// 		fmt.Println("Data received",email)
// 		time.Sleep(time.Second)
// 	}
// }

//What if we want only chan to send or receive
func task(emailChan <-chan string,done chan<- bool){
	defer func ()  { done <- true}()
	//emailChan <- "ddf" //we can send or receive by default 
	//if we need to send only we use header "chan<-"
	//if we need to receive only we use header "<-chan"
	//emailChan<-"ffg" // gives error because it only receive
	//<-done it gives error because i used only to send not receive
	for email:= range emailChan{
		fmt.Println("Data received",email)
	}
}

func main(){

	//What if we have multipleChan
	chan1:=make(chan int)
	chan2:=make(chan string)
	
	go func ()  {
		chan1 <- 12
	}()

	go func ()  {
		chan2 <- "Send"
	}()

	fmt.Println(<-chan1)
	fmt.Println(<-chan2)
	





	//Single Chan //Buffered
	// emailChan:= make(chan string,100)
	// done := make(chan bool)

	// go task(emailChan,done)

	// for i := 0; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("%d data sent",i)
	// }
	// fmt.Println("Done Sending")
	// //this is important
	// close(emailChan)
	// <-done //receive 
	// fmt.Println("Done")







	//How to use wait group in channels unbufferd
	// done:=make(chan bool)
	// go task(done)
	// <-done //block  //receive


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