package main

import "fmt" //format package for i/o: 1.print; 2.user input; 3.write into a file
// imported packages, variables must be used

func main1() {
	fmt.Print("Hello World")
	// to run: go run main.go -> looks for a function called main
	fmt.Println("Hello World")
	// ln gives new line
	fmt.Print("Hello World")
}

func main() {
	var conferenceName = "Go Conference"
	//var conferenceName string = "Go Conference"
	const conferenceTickets = 50   // constants, doesnt change
	var remainingTickets uint = 50 //unssigned
	//var remainingTickets uint = 50 //unit: unassigned int then u should assign it -ve later
	// fmt.printf prints formatted data

	// syntactic sugar: a feature that makes u do smth more easily
	// var conferenceName = "Go Conference" is equiovalent to:
	//conferenceName := "Go Conference"
	// this doesnt work for const or explicit type definition

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	//%T is the placeholder for type of whatever variable we're refering in the end

	// fmt.Print("Welcome to", conferenceName, "booking application")
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickes and %v are still avaiable", conferenceTickets, remainingTickets)
	//%v is a default placeholder
	fmt.Println("Get your tickets here to attend")

	//when we assign value immediately, go assumes the type for this variable
	//when we dont assign value immediately, we need to define the type explicitly
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

	remainingTickets = remainingTickets - userTickets

	// var bookings = [50]string{"Nana", "Nicole", "Peter"}
	var bookings = []string{} //unknown length //dynamic lists using slice
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
	fmt.Printf("These are all our bookings:%v\n", bookings)
	//slice is an abstraction of an Array, variable-length, index-based, resizable
}
