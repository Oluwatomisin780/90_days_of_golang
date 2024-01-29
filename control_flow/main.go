package main

import (
	"fmt"
	"os"
	"strconv"
)


func main(){
	price,inStock:= 100,true
	
	//control flow
	if price>40 {
     fmt.Println("It's too expensive")
	}
	_ =inStock

	if price <= 100 && inStock ==true {
		fmt.Println("Buy It !!")
	}
	if price<100{
		fmt.Printf("Only %d left in stock\n",price)
	}else if price ==100{
		fmt.Println("on the edge")
	}else{
		fmt.Println("It's too expensive")
	}

	age:=50

	if age< 18{
	  fmt.Println("Sorry yah eligible to vote")
	}else if age ==18 {
		fmt.Println("This your first vote")
	}else if age>18 && age< 100{
		 fmt.Println("you can now vote ")
	}else{
		fmt.Println("Invalid age")
	}
//command line Input in Go
  fmt.Println("os.Args",os.Args)
   //simple statement
   i,err := strconv.Atoi("45")

   if err != nil{
	fmt.Println(err)
   }else{
	fmt.Println(i)
   }
 //simple  statement
if  i,err:= strconv.Atoi("20"); err ==nil {
	fmt.Println("No error, i is:", i)
}else{
	fmt.Println(err)
}
	//

 if args := os.Args; len(args) !=2{
	fmt.Println("one argument is required")
 }else  if km,err:= strconv.Atoi(args[1]); err != nil{
   fmt.Println("The argument must be an integer",err)
 }else{
	fmt.Printf("%d km in miles %v \n", km,float64(km)*1.609)
	//  fmt.Printf("%T\n",)
    
 }
 // loops
  for i:= 0; i<=10; i++{
	 fmt.Println(i)
  }

  j:=10 
 for j>=0{
	 fmt.Println(j)
	 j--
 }

 sum  :=0 
 for{ 
	sum++
 }
// println(sum)
//----------------------------------------------------------------------
// for and break
//   count :=0
//   for i:=0; true; i++{
// 	if i%3==0 {
// 		fmt.Printf("%d is divisible by 13 \n", i)
// 		count++
// 	}
// 	if  count==10{
// 		break;
// 	}  
//   }
//   fmt.Println("just a message after the for loop")

  //-----------------------------------------------------------------------------------
//   dkdd
 
  
}