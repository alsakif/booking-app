package main

import (
	"booking-app/functions"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var firstName string
var lastName string
var email string
var remainingTickets uint = 50
var bookedTickets uint
var bookingList []string

func main() {

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

		newBookingList := functions.GetConfimedBookingList(firstName, lastName, bookingList)
		bookingList = newBookingList

		functions.ConfirmationMsg(firstName, lastName, bookedTickets, email, remainingTickets, bookingList)
	}

	functions.LastMsg()
}
