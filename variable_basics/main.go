package main

import (
	"fmt"
)

func main(){
	var age int =8
    var name ="oluwatomisin"
fmt.Println("Age:",age)
	// fmt.Println("Name:", name)
	_=name

	s := "learning golang"
	fmt.Println(s)
	name= "andrei"
	 name1 := "james"
	 _=name1
	 //declaring multiple variable in golang
	 car,cost:= "audi",5000
	 fmt.Println(car,cost)

	 var(
		salary float32 
		firstName string
		gender bool
	 )
	 fmt.Println(salary,firstName,gender)

	 var a,b,c int
	  fmt.Println(a,b,c)
	var i,j int 
	 i,j = 5,8
	 // swaping variablres
	 j,i= i,j
	 fmt.Println(i,j)
	//To allow unused variable
	 //_,_= i,j
}