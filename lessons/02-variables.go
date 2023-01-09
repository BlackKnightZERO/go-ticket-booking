package lessons

import (
	"fmt"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	// var conferenceName = "Go Conference"
	// fmt.Println(remainingTickets, &remainingTickets)
	// remainingTickets = -1

	fmt.Printf("Type of conferenceName: %T\nType of conferenceTickets: %T\nType of remainingTickets: %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets to attend now")

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email:")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets you want to get:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	// var fullNames = [50]string{"Arif", "Faysal", "Ayon"}
	var fullNames [50]string
	fullNames[0] = "Random Name"
	fullNames[1] = firstName + " " + lastName

	fmt.Printf("The whole array: %v\n", fullNames)
	fmt.Printf("The first element: %v\n", fullNames[0])
	fmt.Printf("The second element: %v\n", fullNames[1])
	fmt.Printf("Type of array: %T\n", fullNames)
	fmt.Printf("Length of array: %v\n", len(fullNames))

}
