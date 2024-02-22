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
  m:=[]int{}
fmt.Println(m==nil)
a,b:= []int{1,2,3}, []int{1,2,3}
// fmt.Println(v==b)

//var eq bool=true
 var eq bool =true
 a= nil 
for i,valueA := range a {
	if valueA !=b[i]{
	eq =false
	break
	}
}
if len(a) !=len(b){
	eq= false
}
if eq{
	fmt.Println("a and b slices are equal")


}else{
	fmt.Println("a and b slices are not equal") 
}
//Appending to a slice and 
numbers1:=  [] int {1,2}
numbers1 = append(numbers1, 10) 
numbers1 = append(numbers1, 20,30,40,50)
fmt.Println(numbers1)

n1:= [] int {100,200}
numbers1 = append(numbers1,n1...)

  fmt.Println(numbers1)

//copying a slice
 src:= [] int {12,23,43}
  
 dist:= make([]int,len(src))

 nn := copy(dist, src)

  
 fmt.Println(src,dist,nn )
 //slice expression 
 a1:= [5] int {1,2,3,4,5}
 b1 := a1[1:3]
 fmt.Printf("%v,  %T \n",b1,b1)

 s1 := [] int {1,2,3,4,5,6}
s2 := s1[0:4] 
   fmt.Println(s2)
//slice bcking //underlying array

//s3 := []
}