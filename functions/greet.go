package functions

import "fmt"

func GreetUsers(confName string, confTicket int, remainingTickets uint) {

	fmt.Println("Welcome to", confName, "booking application")
	fmt.Println("We have total of", confTicket, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend the conference!")
}

func LastMsg() {
	fmt.Println("All Tickets are booked. Thank you for your concern. Please comeback in next season.")
}
