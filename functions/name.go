package functions

import (
	"fmt"
	"regexp"
)

func GetFirstName(firstName string) string {
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

func GetLastName(lastName string) string {
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
