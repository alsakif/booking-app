package main

import (
	"booking-app/functions"
	"fmt"

	//"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var firstName string
var lastName string
var email string
var remainingTickets uint = 50
var bookedTickets uint

// var userDetail = make([]userMetaData, 0)
var bookingList []string

//var wg = sync.WaitGroup{}

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

		newBookingList := functions.GetConfimedBookingList(firstName, lastName, email, bookedTickets, bookingList)
		bookingList = newBookingList
		//wg.Add(1)
		//initiate another thread
		functions.ConfirmationMsg(firstName, lastName, bookedTickets, email, remainingTickets, bookingList)
		go sendTicket()

	}
	//wg.Wait()
	functions.LastMsg()
}

func sendTicket() {
	time.Sleep(30 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", bookedTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	//wg.Done()
}
