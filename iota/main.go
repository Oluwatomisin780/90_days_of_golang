package main

import "fmt"


func main(){
	const(
		c1=iota
		c2=iota
		c3=iota
	)
    var w = 676
	fmt.Println(c1,c2,c3)
	fmt.Print((w))
	 const(
		North =iota
		East
		West 
		South
	 )

	 fmt.Println(West)

	const(
		a =(iota*2)+3 //default 
		b 
		c 
	)

	fmt.Println(a,b,c)

	const(
		x = -(iota +2 )
		_
		y 
		z 
	)
	fmt.Println(x,y,z)
}