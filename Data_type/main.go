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
	fmt.Printf("%T\n",r)
    fmt.Println(r)

	//Bool Type
	 var b bool = true
	 fmt.Printf("%T",b)

	// Strin (" %T \n",s)

	//arrays
	var numbers = [4] int{2,4,5,6}
	fmt.Printf("%T\n",numbers)
	//slice Type
	var  cities = []string {"london", "tokyo","New York"}
	fmt.Printf("%T\n",cities)

//Map TYPES
m := map[string]float64{
	"USD":3.33,
	"EUR":70.67,

}
fmt.Printf("%T\n",m)

//STRUCT TYPE
type Person struct{
 name string
 age int
}
var you Person
fmt.Printf("%T\n",you)

//POINTER TYPE
var x int = 2 
  ptr:= &x
  fmt.Printf("%T\n",ptr)
  fmt.Printf("%T\n",f)

  //operator are symbol 
  e,f:= 4,2 

 d:= (e+f)/(e-f)*2
 d = 9% e
 //increment assgnment
 e+=f 
println(e)
 //decrement assgnments
f-=4

 fmt.Println(d)
 //multiple assignment
 r *= 45
 // comparison  and logical operator
 // == equal to 
// != not equal to
// < less
// >greater 
// <= less or equal
// >= greater or equal
// 
  s,j:= 5,10
 fmt.Println(s==j)
 fmt.Println( s !=j)
 fmt.Println(s>j )
 //conditional operators
 // && -->  conditional and
 // || --> conditional or 
 // 
 fmt.Println(s>1&&j== 20)
 fmt.Println(j== 10 || s >100)
}

//FUNCTION TYPE 
func f(){;

}