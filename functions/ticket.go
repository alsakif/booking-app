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

func GetConfimedBookingList(firstName string, lastName string, bookedTickets uint, email string, remainingTickets uint) {
	var bookings []string
	bookings = append(bookings, firstName+" "+lastName)

	cnfirmedList := []string{}
	for _, booking := range bookings {
		var name = strings.Fields(booking) //Splits the string with white space as separator and addresses every splits with index number
		cnfirmedList = append(cnfirmedList, name[0])
	}

	fmt.Printf("Thank you %v %v for booking %v ticket/s. You will receive a confirmation email at %v.\n", firstName, lastName, bookedTickets, email)
	fmt.Println(remainingTickets, "tickets are left !")
	fmt.Println("Confirmed Booking List", cnfirmedList)
	fmt.Println("Confirmed Booking List with full name", bookings)
}
