package main

import (
	"fmt"

	"example.com/structcustom/person"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	person1, err := person.New(firstName, lastName, birthdate)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	person1.PrintData()
	person1.UpdateBirthdate("02/02/2002")
	person1.PrintData()

	admin1, err := person.NewAdmin("firstName", "lastName", birthdate)
	fmt.Println("admin1 : ", admin1.Level)
	admin1.PrintData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
