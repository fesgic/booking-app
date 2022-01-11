package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = uint(conferenceTickets)
	fmt.Printf("Welcome to our %v\n"+
		"We have a total of %v tickets and %v are still available \n"+
		"Get your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)

	var bookings []string
	for {
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

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you cant"+
				" book %v tickets\n", remainingTickets, userTickets)
			continue
		}
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets.\n"+
			"You will receive a confimation email at %v\n"+
			"Remaining tickets are %v\n", firstName, lastName, userTickets, email, remainingTickets)

		firstnames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstnames = append(firstnames, names[0])
		}
		fmt.Printf("This first names of bookings: %v\n", firstnames)

		if remainingTickets == 0 {
			//end the program
			fmt.Printf("Our conference is booked out, Come again nxt yr")
			break
		}
	}
}
