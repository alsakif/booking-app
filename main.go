package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Package level variables
const conferenceTickets = 50

var conferenceName = "Go Conference"
var firstName string
var lastName string
var email string
var remainingTickets uint = 50
var bookedTickets uint

func main() {

	for remainingTickets > 0 {

		greetUsers()

		//Dosen't need to pass the variable value while calling function, as the variables are package level
		updatedFirstName := getFirstName()
		updatedLastName := getLastName()
		updatedEmail := getEmail()
		newBookedTicketsCount, newRemainingTicketsCount := getTicket()

		firstName = updatedFirstName
		lastName = updatedLastName
		email = updatedEmail
		bookedTickets = newBookedTicketsCount
		remainingTickets = newRemainingTicketsCount

		getConfimedBookingList()

	}
	lastMsg()
}

func greetUsers() {

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend the conference!")
}

func getFirstName() string {
	var changedFirstName string
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	regexPattern := regexp.MustCompile(`^[a-zA-Z]{3,10}$`).MatchString(firstName)

	for !regexPattern {
		fmt.Println("Names must be only letters!")

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		regexPattern = regexp.MustCompile(`^[a-zA-Z]{3,10}$`).MatchString(firstName)
	}
	changedFirstName = firstName
	return changedFirstName
}

func getLastName() string {
	var changedLastName string
	fmt.Println("Enter your Last name: ")
	fmt.Scan(&lastName)

	lastNameRegexPattern := regexp.MustCompile(`^[a-zA-Z]{3,10}$`).MatchString(lastName)

	for !lastNameRegexPattern {
		fmt.Println("Names must be only letters!")

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		lastNameRegexPattern = regexp.MustCompile(`^[a-zA-Z]{3,10}$`).MatchString(lastName)
	}
	changedLastName = lastName
	return changedLastName
}

func getEmail() string {
	var changedEmail string
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	emailRegexPattern := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`).MatchString(email)

	for !emailRegexPattern {
		fmt.Println("Invalid Email Address!")

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
		emailRegexPattern = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`).MatchString(email)
	}
	changedEmail = email
	return changedEmail
}

func getTicket() (uint, uint) {
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

func getConfimedBookingList() {
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

func lastMsg() {
	fmt.Println("All Tickets are booked. Thank you for your concern. Please comeback in next season.")
}
