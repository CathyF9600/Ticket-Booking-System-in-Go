package main

// "main" package has two go-files: main.go and helper.go
// to run main.go helper.go
import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	// isValidCity := city == "Singapore" || city == "London"
	// isInvalidCity := city != "Singapore" && city != "London"
	return isValidName, isValidEmail, isValidTicketNumber
}
