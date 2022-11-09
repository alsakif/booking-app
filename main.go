package main

import (
	"booking-app/functions"
	"fmt"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var firstName string
var lastName string
var email string
var remainingTickets uint = 50
var bookedTickets uint
var userDetail = make([]userMetaData, 0)
var bookingList []string

type userMetaData struct {
	Firstname string
	Lastname  string
	Email     string
	Tickets   uint
}

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

		newBookingList := getConfimedBookingList()
		bookingList = newBookingList

		functions.ConfirmationMsg(firstName, lastName, bookedTickets, email, remainingTickets, bookingList)
	}

	functions.LastMsg()
}

func getConfimedBookingList() []string {

	var userdata = userMetaData{
		Firstname: firstName,
		Lastname:  lastName,
		Email:     email,
		Tickets:   bookedTickets,
	}

	userDetail = append(userDetail, userdata)
	fmt.Println("User Detail", userDetail)
	firstNames := []string{}
	for _, booking := range userDetail {
		firstNames = append(bookingList, booking.Firstname)
	}
	bookingList = firstNames
	return bookingList
}
