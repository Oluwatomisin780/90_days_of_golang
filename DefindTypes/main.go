package main

import (
	"fmt"
)
 


func main(){
	type speed uint
	type kilo float64
	type mileStond float64 
	var s1 speed= 10 
	var s2 speed = 20
   //conversion of types
var x = float32(s2)
_=x 
	fmt.Println(s2-s1) 
	// var ParisToLondon kilo = 65
	// var distance mileStond 
	//distance = ParisToLondon/0.67

 // type alias
 var a uint8 = 10
 var b byte 
    b=a
	_=b

	//type alias  
	type second = uint 
	var hour second  = 3600
	fmt.Printf("mins in an hour: %d \n", hour/60)
 }