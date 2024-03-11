package main

import (
	"fmt"
	"time"
)


func main(){
	c1:=make( chan int) //unbuffered channel
	c2 := make(chan int ,3) //buffered channel 
    _= c2
	go func (c chan int)  {
		fmt.Print("before")
		c<- 10
		fmt.Println("After")
	}(c1)
	time.Sleep(time.Second*2)

	fmt.Println("main go routine starts rceiving data") 
	d := <- c1 

	fmt.Println(d)
	time.Sleep(time.Second)
}