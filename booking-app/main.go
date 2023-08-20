package main

import (
	"fmt"
	"strings"
)

var conferenceName = "Go Conference"
const conferenceTickets = 50	
var remainingTickets uint = 50

func welcomeMsg(conferenceName string, conferenceTickets int, remainingTickets uint){
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")
	fmt.Println()
}

func getUserInput(conferenceName string, remainingTickets uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	for {
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)
		
		fmt.Println("Enter your email name")
		fmt.Scan(&email)
		
		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets

		confirmationMessage := fmt.Sprintf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		ticketCheck := fmt.Sprintf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		
		fmt.Print(confirmationMessage)
		fmt.Print(ticketCheck)
		fmt.Print(bookingList(firstName, lastName))
	}
}

func bookingList(firstName string, lastName string) string {
	bookings := [] string {}
	firstNames := []string {}
	
	bookings = append(bookings, firstName + " " + lastName)
	for _, booking := range bookings{
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	
	bookingMessage := fmt.Sprintf("These are all our bookings: %v\n", firstNames)
	return bookingMessage

}

func main(){
	welcomeMsg(conferenceName, conferenceTickets, remainingTickets)
	getUserInput(conferenceName, remainingTickets)
}