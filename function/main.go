package main

import (
	"fmt"
	"math"
	"strings"
)

// function parameter,argumnt and return keyword
func f1(){
	fmt.Println("this is f1 function")
}
func f2(a int,b int){
	fmt.Println("sum " ,a+b)
}
// shorthand parameter notation 
func f3 (a,b,c int,d ,e float64,g bool){
fmt.Println(a,b,c,d,e,g)
}
func f4 (a float64)float64{
return math.Pow(a,a)
}
// named reuturn  values
func sum(a,b int)(s int){
fmt.Println(s)
s=a+b
return
}
//variodic function
func f6(a...int){
	fmt.Printf("%T\n",a)
	fmt.Printf("%#v\n",a)

}
func  f7(a...int ){
	a[0]= 50
}
func fr(a...int){
	a[0]= 32
}
func sumAndProduct(a...float64)(float64,float64){
sum:= 0.
product := 1.
for _,v :=range a{
	sum += v
	product *=v
}
return sum,product
}
// mixing variodic parameter and normal parameter
func personInformation(age int, names ...string) string{
fullName := strings.Join(names," ")
returnString:=fmt.Sprintf("Age: %d, Full Name : %s",age,fullName)
return returnString
}
// defer statement or postpone the execution of the function until the surrounding function returns even normally or antropogenic
func foo(){
 fmt.Println("this is foo()")
}
func boo(){
	fmt.Println("this is boo()")
}

func foobar(){
	fmt.Println("foo boo")
}
func  increment(x int) func() int{
	return func() int {
		x++
		return x
	}
}
func main(){
 f1()
 f2(4,5)
 f3(2,3,4,4.5,65.0,false)
 p:= f4(4.4)
 fmt.Println(p)
 mySumm:=sum(4,8)
 fmt.Println(mySumm)
 f6(1,3,2,3,4,)
nums := []int {1,2,3} 
nums = append(nums, 4,5,6)
f7(nums...)
fr(nums...)
fmt.Println("nums:",nums)
s,r:= sumAndProduct(2.0,5.5,4.0)

fmt.Println(s,r)

info:= personInformation(30,"tomisin","fade","toyosi")
fmt.Println(info)
defer foo()
boo()
fmt.Println("After defer foo and calling boo")

// anonymous function----->It  can be use as closure
 func(msg string){
	fmt.Println(msg)
 }("Anonymous Function")

 a:= increment(10)
 fmt.Printf("%T\n",a)
 a()

 fmt.Println(a())
}  