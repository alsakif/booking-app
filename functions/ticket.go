package functions

import (
	"fmt"
	"strconv"
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

func GetConfimedBookingList(firstName string, lastName string, email string, bookedTickets uint, bookingList []string, userDetail []map[string]string) []string {

	var userdata = make(map[string]string)

	userdata["Firstname"] = firstName
	userdata["Lastname"] = lastName
	userdata["Email"] = email
	userdata["Tickets"] = strconv.FormatUint(uint64(bookedTickets), 10)

	userDetail = append(userDetail, userdata)
	fmt.Println("User Detail", userDetail)

	for _, booking := range userDetail {
		bookingList = append(bookingList, booking["Firstname"])
	}
	return bookingList
}
