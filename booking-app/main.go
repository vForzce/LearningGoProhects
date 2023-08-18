package main

import "fmt"


func welcomeMsg(conferenceName string, conferenceTickets int, remainingTickets int){
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and ", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here")
}

func main(){
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	welcomeMsg(conferenceName, conferenceTickets, remainingTickets)
}