package main

import (
	"fmt"
	f "fmt"
	"strings"
)

func main(){
var array[4 ] int 
 f.Printf("%v\n",  array )
//declaring an  array
  var a1 =[4]float64{}
 f.Printf("%#v\n",a1)
 // declaring an array with a given value 
  var a2 = [4]string{"tomisin","oluwatomisin","tosin","ola"}
   f.Printf("%#v\n",a2)

   //intializing some element of an array
 var a3=[4]string{"tms","michael"}
 f.Printf("%#v\n",a3)

 //ellipis operator
 var a5=[...]int{4,5,6,7,8,4}
  f.Printf("%#v\n",a5) 
  f.Printf("the length of a5 %d\n",len(a5))

  // array operations 
    numbers:=[3]int{3,4,5}
    f.Printf("%#v\n", numbers) 

    numbers[0]= 7
     f.Printf("%#v\n", numbers) 
    
     // iterating through an array
     for i,v :=range numbers{
      fmt.Println("index:",i,"value:",v)
     }
fmt.Println(strings.Repeat("-",20))
     //second method
       for i:=0; i<len(numbers); i++{
         f.Println("position:" ,i,"values:",numbers[i])
       }
  // multi dimensional array
  balances:= [2][3]int{
    {2,3,4},
    {5,6,7},
  }
   f.Printf("%#v\n", balances) 

   //------------------------------------------------------------------
   //Array With Keyed Elements

   grades:= [3]int{
    1:10,
    0:20,
    2:30,
   }
   f.Println(grades)

  accounts:= [3]int{2:4}
  fmt.Println(accounts)
}