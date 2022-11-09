package main

import (
	"booking-app/functions"
	"fmt"
	"sync"
	"time"
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

var wg = sync.WaitGroup{}

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
		wg.Add(1)
		go sendTicket() //initiate another thread
		functions.ConfirmationMsg(firstName, lastName, bookedTickets, email, remainingTickets, bookingList)

	}
	wg.Wait()
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

func sendTicket() {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", bookedTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
