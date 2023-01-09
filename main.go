// go mod init booking-app
// go run main.go

package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings = []string{}

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	for {
		if remainingTickets == 0 {
			fmt.Println("All tickets are sold out")
			break
		}
		fmt.Printf("Welcome to %v booking application\n", conferenceName)
		fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
		fmt.Println("Get your tickets to attend now")

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email:")
		fmt.Scan(&email)
		fmt.Println("Enter the number of tickets you want to get:")
		fmt.Scan(&userTickets)

		bookings = append(bookings, firstName+" "+lastName)
		remainingTickets = remainingTickets - userTickets

		firstNames := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		fmt.Printf("List of booked users: %v\n", firstNames)

	}

}
