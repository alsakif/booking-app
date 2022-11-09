package functions

import (
	"fmt"
)

type userMetaData struct {
	Firstname string
	Lastname  string
	Email     string
	Tickets   uint
}

var userDetail = make([]userMetaData, 0)

func GetTicket(bookedTickets uint, remainingTickets uint) (uint, uint) {
	var maxTickets uint = 45
	var newCount uint

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&bookedTickets)

	for bookedTickets == 0 {
		fmt.Println("Invalid Input!")

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&bookedTickets)

	}
	for bookedTickets > remainingTickets {
		fmt.Println(remainingTickets, "tickets are available. You have asked too much!")

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&bookedTickets)

	}

	for bookedTickets > maxTickets {
		fmt.Println("You can not buy more than ", maxTickets, "tickets")

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&bookedTickets)
	}

	newCount = bookedTickets
	remainingTickets = remainingTickets - newCount

	return newCount, remainingTickets
}

func GetConfimedBookingList(firstName string, lastName string, email string, bookedTickets uint, bookingList []string) []string {

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
