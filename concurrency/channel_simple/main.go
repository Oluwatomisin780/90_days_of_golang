package main

import "fmt"

func channel(n int, ch chan int){
	ch  <- n
}
func main (){
	c:= make( chan int)
	//only for Recieveers
	c1 := make(<- chan string)

	//only for sending
c2:= make (chan<- string)

fmt.Printf("%T %T %T \n",c,c1,c2)


go channel(10 ,c)

n := <- c
fmt.Println(n)
fmt.Println("Exit")
}