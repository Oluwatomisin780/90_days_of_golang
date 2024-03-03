package main

import "fmt"

func charge(s *int) *float64{
	*s =100
	b := 5.5 

	return &b
}
//
func changeValues(quantity int ,price float64,name string,sold bool){
  quantity=3
  price=900.
  name = "yam" 
  sold=false
}
func changeValueBypointer(quantity *int, price *float64, name *string , sold *bool){
	*quantity= 3
	*price= 300.
	*name="apple"
	*sold = true
}
 type Product struct {
	price float64
	productName string
 }

 func changeProduct (p Product){
   p.price=200.
   p.productName ="bicycle"
 }
 func changeProductValue (p *Product){
   (*p).price=200
   p.productName ="bicycle"
 }
func main(){
	name := "tomisin"

	fmt.Println(&name)
	//Declaring a  pointer 
	var x int =2
	ptr := &x

	fmt.Printf("ptr is a type of %T with a value of %v",ptr,x)


	//
	var ptr1 *float64
	_=ptr1
	 
	p:= new (int)
x =100
p =&x
//to print out address of a variable you use %p
fmt.Printf("p is of type %T with a value of %v \n",p,p)

*p= 90 // equivalent to x
//-----comparing pointers-----
a:= 5.5
p1 := &a 
pp1:=&p1
fmt.Printf("the value of p1 %v and address of p1 %v", p1,&p1)
fmt.Printf("the value of pp1 %v and address of p1 %v", pp1,&pp1)
 fmt.Printf("*p1 is %v\n",*p1)
 fmt.Printf("*pp1 is %v\n",*pp1)
 //two pointers are equal if they point to same varaible vice versa
 g:= 6

p2:= &g

p4:= &g
fmt.Println(p2==p4)
x1:= 8 
p3 := &x1
fmt.Println("value of x before calling change()", x1)
charge(p3)
fmt.Println("value of x before calling change()", x1)


// passing pointer to function 
 quantity,price, name,sold := 5,300.,"laptop",true   
 fmt.Println("Before calling changevalue func() ", quantity, price, name,sold)
 changeValues(quantity,price,name,sold)
 fmt.Println("After calling changevalue func() ", quantity, price, name,sold)
 changeValueBypointer(&quantity,&price,&name,&sold)
 fmt.Println("After calling changevalue func() ", quantity, price, name,sold)

gift :=Product{
	price: 100,productName: "watch",
}
changeProduct(gift)
fmt.Println(gift)

fmt.Println("Before calling changeProductByPointer", gift)
changeProductValue(&gift)
fmt.Println("After calling changeProductByPointer", gift)
} 