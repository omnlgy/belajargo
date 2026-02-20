// package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func saveBalanceToFile(balance float64) {
	balanceString := fmt.Sprint(balance)

	os.WriteFile("balance.txt", []byte(balanceString), 0644)
}

func readBalanceFromFile() (float64, error) {
	balanceBytes, err := os.ReadFile("balance.txt")

	if err != nil {
		return 0, errors.New("balance file not found")
	}

	balanceString := string(balanceBytes)
	balanceFloat, err := strconv.ParseFloat(balanceString, 64)

	if err != nil {
		return 0, errors.New("invalid balance format")
	}

	return balanceFloat, nil
}

func main() {
	balance, err := readBalanceFromFile()
	// var choice int

	if err != nil {
		fmt.Println("ERORR! :", err)
		panic("Something wen't wrong!")
	}

	fmt.Println("Welcome to Go Bank!")
	for {
		// for choice != 4 {
		fmt.Println("\nWhat do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
				saveBalanceToFile(balance)
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
				saveBalanceToFile(balance)
				fmt.Println("Your balance is now:", balance)
			}

		} else if choice == 4 {
			fmt.Println("Goodbye!")
			break
		} else {
			fmt.Println("Invalid choice")
		}
	}

	fmt.Println("Thnaks for using Go Bank!\n")
}
