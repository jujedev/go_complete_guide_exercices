package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getUserInput("Enter revenue: ")
	expenses = getUserInput("Enter expenses: ")
	taxRate = getUserInput("Enter tax rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Printf("%.2f\n", ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scanln(&userInput)

	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
