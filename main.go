package main

import (
	"booking-app/functions"
)

func main() {

	var conferenceName = "Go Conference"
	var firstName string
	var lastName string
	var email string
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookedTickets uint

	for remainingTickets > 0 {

		functions.GreetUsers(conferenceName, conferenceTickets, remainingTickets)
		updatedFirstName := functions.GetFirstName(firstName)
		updatedLastName := functions.GetLastName(lastName)
		updatedEmail := functions.GetEmail(email)
		newBookedTicketsCount, newRemainingTicketsCount := functions.GetTicket(bookedTickets, remainingTickets)

		firstName = updatedFirstName
		lastName = updatedLastName
		email = updatedEmail
		bookedTickets = newBookedTicketsCount
		remainingTickets = newRemainingTicketsCount

		functions.GetConfimedBookingList(firstName, lastName, bookedTickets, email, remainingTickets)
	}

	functions.LastMsg()
}
