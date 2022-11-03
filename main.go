package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookedTickets uint
	var bookings []string

	start := time.Now()
	for remainingTickets > 0 {
		var firstName string
		var lastName string
		var email string

		//Function with pre defined input parameter
		greetUsers(conferenceName, conferenceTickets, remainingTickets)

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		regexPattern := regexp.MustCompile(`^[a-zA-Z]{3,10}$`).MatchString(firstName)

		if !regexPattern {
			fmt.Println("Names must be only letters!")
			continue
		}

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		lastNameRegexPattern := regexp.MustCompile(`^[a-zA-Z]{3,10}$`).MatchString(lastName)

		if !lastNameRegexPattern {
			fmt.Println("Names must be only letters!")
			continue
		}

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		emailRegexPattern := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`).MatchString(email)

		if !emailRegexPattern {
			fmt.Println("Invalid Email Address!")
			continue
		}

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&bookedTickets)

		bookingTicketValidation(bookedTickets, remainingTickets)

		remainingTickets = remainingTickets - bookedTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v ticket/s. You will receive a confirmation email at %v.\n", firstName, lastName, bookedTickets, email)
		fmt.Println(remainingTickets, "tickets are left !")

		firstNames := getConfimedBookingList(bookings)
		fmt.Println("Confirmed Booking List", firstNames)
	}
	fmt.Println("All Tickets are booked. Thank you for your concern. Please comeback in next season.")
	duration := time.Since(start)
	fmt.Println(duration.Milliseconds())
}

func greetUsers(confName string, confTicket int, remainingTickets uint) {

	fmt.Println("Welcome to", confName, "booking application")
	fmt.Println("We have total of", confTicket, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend the conference!")
}

func bookingTicketValidation(bookedTickets uint, remainingTickets uint) {
	var maxTickets uint = 45
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
}

func getConfimedBookingList(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var name = strings.Fields(booking) //Splits the string with white space as separator and addresses every splits with index number
		firstNames = append(firstNames, name[0])
	}
	return firstNames
}
