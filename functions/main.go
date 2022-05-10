package main

import (
	"fmt" //format package for i/o: 1.print; 2.user input; 3.write into a file
	"strings"
)

// imported packages, variables must be used

//establish package-level varaibles that are available to all functions
var conferenceName = "Go Conference" // cannot user syntactic sugar for package-level variables
const conferenceTickets = 50         // constants, doesnt change
var remainingTickets uint = 50       // unssigned
var bookings = []string{}            // unknown length //dynamic lists using slice

// best practice: create variables as local as possible

func main() {
	//call function
	greetUsers()

	for {
		// to add conditions:
		// for remainingTickets > 0 && len(bookings) < 50 {
		//when we assign value immediately, go assumes the type for this variable
		//when we dont assign value immediately, we need to define the type explicitly
		firstName, lastName, email, userTickets := getUserInputs()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			// call function printFirstName
			firstNames := getFirstName()
			fmt.Printf("First names of our bookings:%v\n", firstNames)
			//slice is an abstraction of an Array, variable-length, index-based, resizable

			noTickersRemaining := remainingTickets == 0
			// equivalent to
			// var noTickersRemaining bool = remainingTickets == 0
			if noTickersRemaining {
				//end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break //break the infinite for loop
			}
		} else { //if users ask for more than we have
			// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, userTickets)
			// // break //break out of the loop
			// continue //proceed to the next iteration in the loop
			if !isValidName {
				fmt.Println("Your first name or last name is invalid, try again.")
			}

			if !isValidEmail {
				fmt.Println("Your email is invalid, try again.")
			}

			if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid, try again.")
			}

		}

	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickes and %v are still avaiable\n", conferenceTickets, remainingTickets)
	//%v is a default placeholder
	fmt.Println("Get your tickets here to attend")
}

func getFirstName() []string { //bookings is a slice of string
	//inside bracket: parameter
	//outside bracket: output/return

	//want to print first name only
	firstNames := []string{}
	// for each loop: index,element we get by iterating := range array we iterate in
	for _, booking := range bookings {
		// range iterates over elements for different data structures (not just arrays and slices)
		var names = strings.Fields(booking) //splits the space into array with white space as separator
		firstNames = append(firstNames, names[0])

		//_ is index not used, called blank identifier
		//to ignore a variable you dont want to use
		//So with Go you need to make unused varaibles explicit
	}
	// for arrays and slices, range provides the index and value for each element
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	// isValidCity := city == "Singapore" || city == "London"
	// isInvalidCity := city != "Singapore" && city != "London"
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)
	// &userName is the pointer of userName. It prompts the user to type an input,
	// and fmt will use the address of the input to find/scan the input, otherwise it wont find it.

	fmt.Println("Enter number:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// var bookings = [50]string{"Nana", "Nicole", "Peter"}

	//cant fit many types in one array
	// bookings[0] = firstName + " " + lastName //for index-based append
	bookings = append(bookings, firstName+" "+lastName)
	//on the basis of bookings, add new ele
	//dynamic lists using slice

	//syntactic sugar of var bookings = []string{}:
	// bookings := []string{}

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array length: %v\n", len(bookings))
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive an email confirmation at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
