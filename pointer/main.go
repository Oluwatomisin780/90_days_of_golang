package main

import "fmt"



func main(){
	name := "tomisin"

	fmt.Println(&name)
	//Declaring a  pointer 
	var x int =2
	ptr := &x

	fmt.Printf("ptr is a type of %T with a value of %v",ptr,x)


	//
	var ptr1 *float64
	_=ptr1
	 
	p:= new (int)
x =100
p =&x
fmt.Printf("p is of type %T with a value of %v \n",p,p)

}