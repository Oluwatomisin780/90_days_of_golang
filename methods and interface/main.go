package main

import (
	"fmt"
	"time"
)

type names [] string
func (n names) print(){
  for i,names := range n{
    fmt.Println(names,i)
  }
}
func main(){
//Go does not have classes but you can define methods  on predfined types
  const day = 24* time.Hour

  fmt.Printf("%T\n",day)
  seconds:=  day.Seconds()
   fmt.Printf("%T\n",seconds)
   fmt.Printf("Seconds in a day: %v\n",seconds)
//difference between a  method and function is that a method belongs to a type in a functions
   friends := names{"Dan","Mary"}
   friends.print()

   //var n int64 =99696965
}