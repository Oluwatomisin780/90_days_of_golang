package main

import "fmt"



func main(){
	// struct is like a class  in object oriented programming  lang 
	title1,author1,year1:=  "good",  "tomisin", 2033 
	title2,author2,year2 := "bad","oluwatomisin", 3434
	fmt.Println("book1 ",title1,author1,year1)
	fmt.Println("book1 ",title2,author2,year2)
  type book struct{
title string
author string
year int
  }

  myBook := book{
	title: "rich also die",
	author: "tomisin",
year: 30443,
  }
  fmt.Println(myBook)
  lastBook:= book{
	title: "fade",
  }
fmt.Println(lastBook.title)

lastBook.author ="tomisin"
lastBook.year =2343
fmt.Printf("%+v\n",lastBook)

// two struxt are equal if there coresponding fields  are equal 
eBook := book{
	title: "fade",author: "tomisin", year: 2343,
}

 fmt.Println(  eBook==lastBook)

 //anonymous struct are struct with no explicity defined struct type alias
diana := struct{// anonymouse struxt
	 firstName,lastName  string;
	 age int
	
}{
	firstName:  "Diana",
	lastName:  "muller",
	age:10,
}

fmt.Printf("%#v\n",diana)
fmt.Println("diana  age is :", diana.age)
//anonymous field type in struct 
type Iwe struct{
	string
	float64
	bool
}
b1:= Iwe {"the lengendary book",45.6,false}
fmt.Printf("%#v\n",b1 )


//embedded struxt
// embedded struct is a sttruct that act a field inside a struct
type Contact struct{
	email ,address string 
	phone int ;
	
}
type Employee  struct{
	name string 
	salary int
	contactInfo Contact
}

john:= Employee{
	 name:"John Doe",
	 salary: 2324,
	contactInfo:Contact {
		email:"johndoe@gmail.com", 
		address: "18 ireoluwa street",
		phone: 222323,
	},
}
fmt.Println(john)
}