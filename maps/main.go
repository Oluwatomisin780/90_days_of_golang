package main

import (
	"fmt"
)

//A map is a collection type  like  array or slices for storing it in keyvalue pairs

 func main(){
var employes map[string] string 
fmt.Printf("%#v\n",employes)
fmt.Printf("No of pairs:%d\n",len(employes))
fmt.Printf("The value for key  %q is %q \n","dane",employes["dane"])


 var account map [string] float64
 fmt.Printf("%#v\n",account["0x434"])

 var m1 map [[5]int] string
 _=m1

 //employes["dane"]= "Programmer"
 peoples :=map[string]float64{}

 peoples["john"]= 21.4
 peoples ["Mary"]=10
 fmt.Println(peoples)

map1:= make(map[int]int)
map1[4]=8
//initializethe map with the  key values pair when declaring it which means using a map literal 
balances :=map[string]float64{
	"usd":23.43,
	"euro":23.11,
	"Ngn": 10.00,
}

fmt.Println(balances)

balances["usd"]=34
balances["ghf"]=43
fmt.Println(balances)

fmt .Println(balances["GHS"])
balances["Ron"] =32
//distingushing between a zero value from missing value in go..... we use coma ok idioms to do that
v,ok:= balances["Ron"]
   

if ok {
	fmt.Println("the balance ron is: " , v)
}else{
	fmt.Println("the run does not exist in map")
}

//comparing map : A map can only be compare  
a2:= map[string]string{
	"A":"a",
}
b2:= map[string]string{
	"A":"a",
}

// fmt.Println(a2== b2) 

 s1:= fmt.Sprintf("%s", a2)
 s2 := fmt.Sprintf("%s",b2)

 fmt.Println(s2,s1)
  if s1==s2 {
	fmt.Println("map is equal")
  }else{
	fmt.Println("map is not equal")
  }
  //how to clone or duplicate a Map

  friends := map [string]int{
	"Dan":40,
	"maira":50,

  }

  neighbour:= friends

  friends["Dan"] = 60

  fmt.Println(neighbour)

  people := make(map[string]int)

  for k,v := range friends{
	people[k]=v
  }
people["Anne"]=46
  fmt.Printf("%#v--------- %#v\n",people,friends)

  delete(friends, "Dan")
  fmt.Println(friends)
 }