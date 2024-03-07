package main

import "fmt"



func main(){
	// channels are use in communicating between two araigninig goroutines

	var ch chan int

	fmt.Println(ch)
	ch= make(chan  int)
	fmt.Println(ch)
	c := make(chan int)
	// <- channel operator ---->
	//send
	c<- 10

	//Recieve
	num := <- c 

	_=num
	fmt.Println(<- c)

	close(c)
}