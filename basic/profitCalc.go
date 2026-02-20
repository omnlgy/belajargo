package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)

	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)

	// fmt.Print("Tax Rate: ")
	// fmt.Scan(&taxRate)

	revenue, err := userInput("Revenue: ")
	if err != nil {
		fmt.Println("ERORR! :", err)
		panic("Something wen't wrong!")
	}

	expenses, err := userInput("Expenses: ")
	if err != nil {
		fmt.Println("ERORR! :", err)
		panic("Something wen't wrong!")
	}

	taxRate, err := userInput("Tax Rate: ")

	if err != nil {
		fmt.Println("ERORR! :", err)
		panic("Something wen't wrong!")
	}

	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	ebt, profit, ratio := calculateEarning(revenue, expenses, taxRate)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

	os.WriteFile("profitCalc.txt", []byte(fmt.Sprintf("%.2f %.2f %.2f", ebt, profit, ratio)), 0644)
}

func userInput(text string) (float64, error) {
	var input float64
	fmt.Print(text)
	fmt.Scan(&input)

	if input <= 0 {
		errorMessage := fmt.Sprintf("%smust be greater than 0", text)
		return 0, errors.New(errorMessage)
	}

	return input, nil
}

func calculateEarning(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
