package helper

// "main" package has two go-files: main.go and helper.go
// to run two packages: go run .
import (
	// "main"
	// you cannot import main
	// to import a variable in main, you need to pass it to whatever function you're using
	"strings"
)

// whenever we want to use a function of one package in another package,
// we have to explicitly export the function from that package
// to be imported in the other package
// in go, to export, just capitalize the first letter

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, RemainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= RemainingTickets
	// isValidCity := city == "Singapore" || city == "London"
	// isInvalidCity := city != "Singapore" && city != "London"
	return isValidName, isValidEmail, isValidTicketNumber
}
