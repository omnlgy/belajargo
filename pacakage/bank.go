package main

import (
	"fmt"

	"example.com/bank/utils"
)

const fileName = "balance.txt"

func main() {
	balance, err := utils.ReadFloadFromFile(fileName)
	// var choice int

	if err != nil {
		fmt.Println("ERORR! :", err)
		panic("Something wen't wrong!")
	}

	fmt.Println("Welcome to Go Bank!")
	for {
		// for choice != 4 {

		listMenu()

		fmt.Print("Your choice: ")
		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is:", balance)
		} else if choice == 2 {
			var amount float64

			fmt.Print("Enter the amount to deposit: ")
			fmt.Scan(&amount)

			if amount <= 0 {
				fmt.Println("Invalid amount, deposit must be greater than 0")
			} else {
				balance += amount
				utils.WriteFloadToFile(balance, fileName)
				fmt.Println("Your balance is now:", balance)
			}

		} else if choice == 3 {
			var amount float64

			fmt.Print("Enter the amount to withdraw: ")
			fmt.Scan(&amount)

			if amount <= 0 {
				fmt.Println("Invalid amount, withdrawal must be greater than 0")
			} else if amount > balance {
				fmt.Println("Insufficient balance")
			} else {
				balance -= amount
				utils.WriteFloadToFile(balance, fileName)
				fmt.Println("Your balance is now:", balance)
			}

		} else if choice == 4 {
			fmt.Println("Goodbye!")
			break
		} else {
			fmt.Println("Invalid choice")
		}
	}

	fmt.Println("Thnaks for using Go Bank!")
}
