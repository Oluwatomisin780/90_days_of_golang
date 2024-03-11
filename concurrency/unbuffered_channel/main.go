package main

import (
	"fmt"
	"time"
)


func main(){

	 c:= make(chan int,3)
	
	 go func (c chan int)  {
		 for i :=0 ;i <= 5; i++{

	 
		fmt.Print("before")
		c<- 10
		fmt.Println("After")
		}
		close(c)
	 
	}(c) 
	fmt.Println("main goroutin sleep for 2 second")
	time.Sleep(time.Second*2)
 for v := range c {
	fmt.Println("main goroutine recieved value from a channel",v)
 }
 fmt.Println(<-c)
// c<-10 // panic
}