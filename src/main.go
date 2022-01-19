package main

import (
	"fmt"
	"strings"
)

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to our %v\n"+
		"We have a total of %v tickets and %v are still available \n"+
		"Get your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Println("Welcome to our Conference")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName, lastName, email string, userTickets, remainingTickets uint) (bool, bool, bool) {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets
	//var isValidCity = city == "singapore" || city == "london"
	//var isInvalidCity = city != "singapore" || city != "london"

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//ask user for their name
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your first email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets, userTickets uint, firstName, lastName, email string, bookings []string) (uint, []string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets.\n"+
		"You will receive a confimation email at %v\n"+
		"Remaining tickets are %v\n", firstName, lastName, userTickets, email, remainingTickets)
	return remainingTickets, bookings

}
func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = uint(conferenceTickets)

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	var bookings []string
	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber == true {
			remainingTickets, bookings := bookTicket(remainingTickets, userTickets, firstName, lastName, email, bookings)
			//getFirstNames(bookings)
			fmt.Printf("This first names of bookings: %v\n", getFirstNames(bookings))

			if remainingTickets == 0 {
				//end the program
				fmt.Printf("Our conference is booked out, Come again nxt yr")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("The email you entered does not contain \"@\" sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}

	//switch statement example
	/*city := "London"
	switch city {
	case "New York":
		//execute code for booking New York Conferences
	case "Singapore", "Hong Kong":
		//execute code for booking Singapore and Hong Kong Conferences
	case "London", "Berlin":
		//some code here
	case "Mexico":
		//some code here
	default:
		fmt.Println("NO valid city selected")
	}*/
}
