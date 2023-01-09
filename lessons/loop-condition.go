package lessons

import (
	"fmt"
	"strings"
)

func main() {

	tempNum := 0
	for {
		fmt.Println("Current Temp Number ", tempNum, "Enter a new number")
		fmt.Scan(&tempNum)
	}

	var tempCity string
	cityList := []string{}
	for {
		fmt.Println("Enter a city")
		fmt.Scan(&tempCity)
		cityList = append(cityList, tempCity)

		fmt.Println("Current City List:")
		for index, city := range cityList {
			fmt.Printf("%v. %v\n", index, city)
		}
	}

	nameList := []string{
		"Arif Faysal",
		"Hiba Al ashraf",
	}
	firstNameList := []string{}

	for _, name := range nameList {
		namedSlices := strings.Fields(name)
		firstNameList = append(firstNameList, namedSlices[0])
	}

	for index, name := range firstNameList {
		fmt.Println(index, name)
	}

	fmt.Println("First Names :", firstNameList)

}
