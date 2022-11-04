package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookedTickets uint
var bookings []string

func main() {

	start := time.Now()
	for remainingTickets > 0 {
		var firstName string
		var lastName string
		var email string

		//Function 1
		greetUsers(conferenceName, conferenceTickets, remainingTickets)

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		regexPattern := regexp.MustCompile(`^[a-zA-Z]{3,10}$`).MatchString(firstName)

		for !regexPattern {
			fmt.Println("Names must be only letters!")

			fmt.Println("Enter your first name: ")
			fmt.Scan(&firstName)
			regexPattern = regexp.MustCompile(`^[a-zA-Z]{3,10}$`).MatchString(firstName)
		}

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		lastNameRegexPattern := regexp.MustCompile(`^[a-zA-Z]{3,10}$`).MatchString(lastName)

		if !lastNameRegexPattern {
			fmt.Println("Names must be only letters!")

			fmt.Println("Enter your last name: ")
			fmt.Scan(&lastName)
			lastNameRegexPattern = regexp.MustCompile(`^[a-zA-Z]{3,10}$`).MatchString(lastName)
		}

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		emailRegexPattern := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`).MatchString(email)

		if !emailRegexPattern {
			fmt.Println("Invalid Email Address!")

			fmt.Println("Enter your email address: ")
			fmt.Scan(&email)
			emailRegexPattern = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`).MatchString(email)
		}

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&bookedTickets)

		//Function 2
		newBookedTicketsCount, newRemainingTicketsCount := bookingTicketValidation(bookedTickets, remainingTickets)

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v ticket/s. You will receive a confirmation email at %v.\n", firstName, lastName, newBookedTicketsCount, email)
		fmt.Println(newRemainingTicketsCount, "tickets are left !")

		//Function 3
		firstNames := getConfimedBookingList(bookings)
		fmt.Println("Confirmed Booking List", firstNames)
		fmt.Println("Confirmed Booking List with full name", bookings)
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

func bookingTicketValidation(bookedTickets uint, remainingTickets uint) (uint, uint) {
	var maxTickets uint = 45
	var newCount uint
	for bookedTickets == 0 {
		fmt.Println("Invalid Input!")

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&bookedTickets)
		newCount = bookedTickets

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
	remainingTickets = remainingTickets - newCount

	return newCount, remainingTickets
}

func getConfimedBookingList(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var name = strings.Fields(booking) //Splits the string with white space as separator and addresses every splits with index number
		firstNames = append(firstNames, name[0])
	}
	return firstNames
}
