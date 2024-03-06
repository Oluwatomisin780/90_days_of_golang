package main

import (
	"fmt"
	"math"
)
func print (s shape){
	fmt.Printf("%#v\n",s)
	fmt.Printf("Area: %v\n",s.area())
	fmt.Printf("Perimeter: %v\n",s.perimeter())
}

type rectangle struct{
		width,height  float64
	}

type circle struct {
	 radius float64
}

func( c circle) area() float64{
	return math.Pi * math.Pow(c.radius,2)
}
func (c circle) perimeter()float64{
	return 2*math.Pi*c.radius
}

func ( r rectangle)area() float64{
	return r.height *r.width
}
func (r rectangle )perimeter() float64{
	return 2*(r.width+r.height)
}
// func printCircle(c circle){
// 	fmt.Println("shape",c)
// 	fmt.Println("area",c.area() )
// 	fmt.Println("shape",c.perimeter())
// }

// func  printRectangle(r rectangle){
// 	fmt.Println("shape",r)
// 	fmt.Println("area",r.area())
// 	fmt.Println("perimeter",r.perimeter())

// }
type shape interface{
	area() float64
	perimeter() float64
}
func  (c circle) volume() float64{
	return 4/3* math.Pi*math.Pow(c.radius,3)
}

func main(){
	// An  Interface is the collection of signature
//  var s shape 

//  fmt.Printf("%T\n", s )
//  ball:= circle{ radius: 4.5}
//  s=ball
//  print(s)
//  fmt.Printf("Type of %T\n", s )
//  room := rectangle{width: 3,height: 5}
//  s= room
//  fmt.Printf("Type of %T\n", s )

 var s shape = circle{radius: 2.3}
 fmt.Printf("Type of %T\n", s )
 //s.volume()
 // type assertion 
ball,ok := s.(circle)
if ok== true{
	fmt.Printf("ball volume : %T\n",ball.volume())
}

// type switch
switch value := s.(type){
case circle:
	fmt.Printf("%#v has circle type\n",value)
case rectangle:
	fmt.Printf("%#v has rectangle type\n",value)
}


// embedded Interface


//Empty Interface 
var empty  interface{


}
empty= 5

fmt.Println(empty)

empty="go"
fmt.Println(empty)

empty= []int{4,5,6,5}
fmt.Println(empty)
}