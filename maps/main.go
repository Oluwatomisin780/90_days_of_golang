package main

import "fmt"

//A map is a collection type  like  array or slices for storing it in keyvalue pairs

 func main(){
var employes map[string] string 
fmt.Printf("%#v\n",employes)
fmt.Printf("No of pairs:%d\n",len(employes))
fmt.Printf("The value for key  %q is %q \n","dane",employes["dane"])
 }