package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "Hello welcome to string_bytes"

	fmt.Println(s1)
	fmt.Println("He says:\"hello\"")

	fmt.Println(`He says :"Hello"`)

	s2 := `I like Go!!`
	fmt.Println(s2)// raw strings

	fmt.Println("price:100 \n Brand: Nike")
	fmt.Println(`c:\users\andrei`)
	fmt.Println("c\\users\\Andrei")
  var s3 = "I love" +" "+"G0" +" "+ "Programming"
 fmt.Println(s3+ "!")

 fmt.Println("Elements at Index 0:",s3[0])

 //coding runes and docode runes
 var1, var2 := 'a' ,'b'
 fmt.Printf("Type: %T ,Value:%d \n",var1,var2)

 str := "tara"
 fmt.Println(len(str))

 fmt.Println("Byte (not rune) at position1:", str[1])

 for i :=0; i<len(str); i++{
	fmt.Printf("%c",str[i])
 }

 for i :=0; i<len(str);{ 
r,size :=utf8.DecodeLastRuneInString(str[i:])

fmt.Printf("%c\n",r)
i += size
 }
}