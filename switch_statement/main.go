package main

import "fmt"

func main() {
	 
	language:= "golang"
	switch language {
	case "python":
		fmt.Println("this is python") 
	case "Javascript":
		fmt.Println("javascript is the best")
	case "go","golang":
		fmt.Println("Good, Go for golang ,we're using curly braces")
	default:
		fmt.Println("Any other language is good start")
	
	}
   n:= 5
   switch true{
   case n%2 ==0:
	 fmt.Println("Even number!")
   }

 
 
}