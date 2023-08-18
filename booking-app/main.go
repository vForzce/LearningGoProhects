package main

import "fmt"


func welcomeMsg(conferenceName string, conferenceTickets int, remainingTickets int){
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")
}

func userInput()string{
	var userName string
	var userTickets int
	fmt.Print("Enter user name\n")
	fmt.Scan(&userName)
	fmt.Print("Enter number of tickets\n")
	fmt.Scan(&userTickets)

	message := fmt.Sprintf("User %v booked %v tickets", userName, userTickets)
	return message
}

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	welcomeMsg(conferenceName, conferenceTickets, remainingTickets)
	result:=userInput()
	fmt.Println(result)	
}