package main

import "fmt"

func main(){
  friends:= [2] string {"Mark" , "Mary"}
  people:= [5]string{"helen","Mark", "Brenda","Antonio","Micheal"}

 outer:
  	for index ,name := range people{
		for _,friend :=range friends {
			if name== friend{
				fmt.Printf("Found a friend %q at index %d \n  ", friend,index)
				break outer
			}
		}
	
	}
	fmt.Println("Next Instruction after break")

  i:=0 

 loop:
	if i<5 {
		fmt.Println(i) 
		i++ 
		goto loop
	}
	//goto todo
// 	x:= 5 
// todo:  
// 	fmt .Println("print something here")
}