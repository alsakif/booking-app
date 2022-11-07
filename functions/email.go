package functions

import (
	"fmt"
	"regexp"
)

func GetEmail(email string) string {
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
