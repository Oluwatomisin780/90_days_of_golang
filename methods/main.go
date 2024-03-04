package main

import (
	"fmt"
	"time"
)

type car struct{
   brand string
  price   int
}
func changeCar(c car, newBrand string,newPrice  int){
  c.price= newPrice
  c.brand = newBrand
}

func (c car) changeCar1(newBrand string,newPrice int){
  c.brand= newBrand
  c.price = newPrice
}
type names [] string
func (n names) print(){
  for i,names := range n{
    fmt.Println(names,i)
  }
}

func (c *car) changeCar2(newbrand string,newPrice int){
 (*c).brand=newbrand
 (*c).price= newPrice
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
   myCar:= car{ brand:"Toyota",price: 40000}
   changeCar(myCar,"realnut",20000)
   fmt.Println(myCar)
   myCar.changeCar1("Audi",210000)
   fmt.Println(myCar)

  (&myCar).changeCar2("seat",34000)
   fmt.Println(myCar)
    

  var yourCar *car
  yourCar=&myCar
 fmt.Println(*yourCar) 


 //valid  ways to a methods on a pointer
 yourCar.changeCar2("Bmw",700000)
 fmt.Println(*yourCar) 
}