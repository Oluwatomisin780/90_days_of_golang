package main

import "fmt"

func  main()  {
	//Int 
	var i1 int8 = 100 //minimum value  of ann int8 is -128
	fmt.Printf("%T\n",i1)
	 var i26 uint16 = 65535
	 fmt.Printf("%T\n",i26)

	 //Floats
	var f1,f2,f4 float64 = 1.1, -.2, 5.
	fmt.Printf( "%T,%T,%T\n",f1,f2,f4,)

	//Rune alias fot int 32 bit
	var r rune='f'
    fmt.Println(r)


}