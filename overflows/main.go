package main

import (
	"fmt"
	"math"
)
func main(){
	var i uint8 = 255
	i++ // overflow
	fmt.Println(i)

//	a:=int8 (225+1) 
   var b int8= 127
    
   fmt.Printf("%d\n", b+1)
   b=-128
   b--
   fmt.Println(b)
   f:= float32(math.MaxFloat32)

   fmt.Println(f)
   f=f*1.2 //overflow to infinite
   fmt.Println(f)

   //converting types
   var x = 3
   var y = 3.2 //float type
//     x = x* y 
 x = x* int(y)
y = y* float64(x)
 fmt.Println(x)
 fmt.Println(y)
}