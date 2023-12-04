package main

import (
	"fmt"
	"strings"
)

func main() {
    conferenceName := "Go Conference"
	const conferenceTicket = 50
	var remainingTicket uint = 50
	// formatting in golang
	fmt.Printf("Welcome to %v Booking application\n", conferenceName )
	fmt.Println("we have total of ", conferenceTicket, "tickets and", remainingTicket, "are still available")
	fmt.Println("Get your tickets here to attend")
	//loop
	 greetUser(conferenceName)
	for{
		var userTicket uint
		var firstName string
		var lastName string
		var email string
// 		 var city string
		//asking user for their name
		fmt.Println("Enter your FirstName:") 
		//& to send data in  memeory
		fmt.Scan(&firstName)

		fmt.Println("Enter your LastName")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email") 
		fmt.Scan(&email)
		fmt.Println("Enter number of ticket")
		fmt.Scan(&userTicket)
		   isValidName :=  len(firstName)>= 2 && len(lastName)>=2
		  isEmailValid:= strings.Contains(email,"@")
		 isValidtTicketNumber :=userTicket> 0 && userTicket<= remainingTicket
		// isValidCity = city == "singapore" || city== "newyork"
		 // check for invalid  city 
	     //invalidCity := city != "singapore" || city != "newyork"
		if isValidName && isEmailValid && isValidtTicketNumber{
			fmt.Printf("We Only have %v ticket available,so you can't book %v ticket \n", remainingTicket,userTicket)
			continue
		}
		remainingTicket=remainingTicket-userTicket
		//fmt.Printf("User %v Booked this %v tickets\n",userName,userTicket)
		fmt.Printf("Thank you for %v %v for bookings %v ticket the confirmation will be send to your %v email\n",firstName,lastName,userTicket,email)
		fmt.Printf("%v tickets remaining for %v\n",remainingTicket,conferenceName)

		//arrays and slices in go
		bookings:=[]string{}
	 
		bookings =append(bookings,firstName+" "+lastName)
		fmt.Printf("The whole array %v\n",bookings)
		fmt.Printf("The first value array %v\n",bookings[0])
		fmt.Printf("The whole array %v\n",len(bookings))
		fmt.Printf("The whole array %T\n",bookings) 
		fmt.Printf("These are all the bookings%v\n", bookings)

		firstNames:=[]string{}
		for _,booking := range bookings{
				var names= strings.Fields(booking)
				firstNames= append(firstNames,names[0])
			}
			fmt.Printf("The firstname of the bookings are %v \n", firstNames)
			var noTicketRemaining bool= remainingTicket ==0
			//alternative syntax
			 //var notickerRemaing := remainingticket....
			if  noTicketRemaining {
				//end program
				fmt.Println("The conference is booked out,Thanks Come back later")
				break
				//elseif in golang
			}else{
				if !isValidName{
					fmt.Println("The firstname or lastname entered is not correct")
				}  
				 if !isEmailValid{
					fmt.Println("the email you entered is invalid")
				}
				if !isValidtTicketNumber{
					fmt.Println("invalid ticket number ")
				}
				fmt.Println("your Input data is invalid")
			}

		}
		
		//loops
		//for loop(infinte loop,)
		//forEach
		//conditioinal loops in go
		// user Input validation
	//switch cases 
	// city :="singapore"
	// switch{
	// case "Newyork":
	// 		 fmt.Printf("")//do something here
	// case  "Singapore":
	// 		fmt.Printf("")//do something here
	// case  "Toronto":
	// 		//do Something here
	// case  "Lagos":
	// 		//do something here
	// default:
	// 	fmt.Println("no valid city")
	// }
	
}

//functions in go

func greetUser(confName string ){
	fmt.Printf("Welcome to our %v Booking System",confName)
}