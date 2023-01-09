package lessons

import (
	"fmt"
)

func main() {

	firstName := "Arif"
	lastName := "Faysal"

	// var bookings []string
	bookings := []string{}
	bookings = append(bookings, firstName+" "+lastName, "Random Name")
	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first element: %v\n", bookings[0])
	fmt.Printf("The second element: %v\n", bookings[1])
	fmt.Printf("Type of slice: %T\n", bookings)
	fmt.Printf("Length of slice: %v\n", len(bookings))

	mappedData := map[int]string{
		1: firstName + " " + lastName,
		2: "Random Name",
	}
	mappedData[3] = "Final Name"
	mappedData[2] = "name random" // updating

	data, ok := mappedData[1]
	fmt.Println("\nKey present or not:", ok)
	fmt.Println("Value:", data)

	_, ok1 := mappedData[100]
	fmt.Println("\nKey present or not:", ok1)

	delete(mappedData, 3) // delete
	fmt.Println(mappedData)

}
