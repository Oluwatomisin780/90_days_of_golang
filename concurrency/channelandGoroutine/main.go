package main

import "fmt"


func factorial(n int , c chan int ){
	f :=1 
	
	for i :=2;i <=n; i++{
		f*=i
	}
c <- f
}

func main(){


	ch:= make(chan int )

	defer close(ch)


	go factorial(5,ch)

	f:= <- ch

	fmt.Println(f)

	for i:=5; i<= 15;i++{
go func(n int, c chan int){
			f :=1 
	
	for i :=2;i <=n; i++{
		f*=i
	}
	c <- f
		}(i,ch)
		fmt.Printf("Factorial of %d\n",i,<- ch)
	}
} 