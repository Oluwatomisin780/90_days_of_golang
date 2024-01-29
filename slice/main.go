package main

import (
	"fmt"
	f "fmt"
)

func main() {
var cities[]string
f.Println("cities is equal to nills: ",cities==nil)
f.Printf("cities  %#v\n",cities)
f.Println(len(cities))

//cities[0] ="london"
numbers:=[] int {1,3,3}
 fmt.Println(numbers)

 nums := make([]int, 4)
 f.Printf("%#v\n",nums)

 type name[]string

 friends := name{"fade","joy","toyosi"}
 f.Println(friends)

 MyFriend := friends[0] 
 f.Println("my best friend is ",MyFriend)

friends[0]= "fiona"
 f.Println("my best friend is",friends[0])

 //iterating in slice
 for index,value := range numbers{
	fmt.Printf("index: %v ,value: %v\n", index,value)
 }
 var n[] int 
 n =numbers
 fmt.Println(n)
 //COMPARING SLICES
 //-----------------------------------------------------------
 var g []int
  fmt.Println("g equal to nil:",g==nil )
  

}