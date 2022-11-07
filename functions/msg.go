package functions

import "fmt"

func GreetUsers(confName string, confTicket int, remainingTickets uint) {

	fmt.Println("Welcome to", confName, "booking application")
	fmt.Println("We have total of", confTicket, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend the conference!")
}
func ConfirmationMsg(firstName string, lastName string, bookedTickets uint, email string, remainingTickets uint, bookingList []string) {
	fmt.Printf("Thank you %v %v for booking %v ticket/s. You will receive a confirmation email at %v.\n", firstName, lastName, bookedTickets, email)
	fmt.Println(remainingTickets, "tickets are left !")
	fmt.Println("Confirmed Booking List", bookingList)
}
func LastMsg() {
	fmt.Println("All Tickets are booked. Thank you for your concern. Please comeback in next season.")
}
