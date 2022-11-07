package functions

import (
	"fmt"
	"strings"
)

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

func GetConfimedBookingList(firstName string, lastName string, bookingList []string) []string {
	var bookings []string
	bookings = append(bookings, firstName+" "+lastName)

	for _, booking := range bookings {
		var name = strings.Fields(booking) //Splits the string with white space as separator and addresses every splits with index number
		bookingList = append(bookingList, name[0])
	}
	return bookingList
}
