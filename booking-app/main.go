package main

import "fmt"

func main() {
    conferenceName := "Go Conference"
	const conferenceTicket = 50
	var remainingTicket uint = 50
	// formatting in golang
	fmt.Printf("Welcome to %v Booking application\n", conferenceName )
	fmt.Println("we have total of ", conferenceTicket, "tickets and", remainingTicket, "are still available")
	fmt.Println("Get your tickets here to attend")
	var userTicket uint
	var firstName string
	var lastName string
	var email string
	//asking user for their name
	fmt.Println("Enter your FullName:") 
	//& to send data in  memeory
	fmt.Scan(&firstName)

	fmt.Println("Enter your LastName")
	 fmt.Scan(&lastName)
    fmt.Println("Enter your email") 
	 fmt.Scan(&email)
	userTicket=45
    remainingTicket=remainingTicket-userTicket
//fmt.Printf("User %v Booked this %v tickets\n",userName,userTicket)
	fmt.Printf("Thank you for %v %v for booking %v ticket the confirmation will be send to your %v email\n",firstName,lastName,userTicket,email)
	fmt.Printf("%v tickets remaining for %v\n",remainingTicket,conferenceName)

	//arrays and slices in go
	 booking:=[]string{}
	
	booking=append(booking,firstName+lastName)
	fmt.Printf("The whole array %v\n",booking)
	fmt.Printf("The first value array %v\n",booking[0])
	fmt.Printf("The whole array %v\n",len(booking))
	fmt.Printf("The whole array %T\n",booking) 
	fmt.Printf("These are all the booking%v\n", booking)
	//loops

}