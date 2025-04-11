package main

import (
	"fmt"
	"os"
)

// Goals
/*	1) Validate user input
*	=> Show error message & exit if invalid input is provided
*	- No negative numbers
* 	- Not 0
*	2) Store calculated returns into file
 */

func main() {
	revenue, expenses, taxRate := setValues()

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)

	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	err := os.WriteFile("results.txt", []byte(results), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Results stored in results.txt")
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func setValues() (float64, float64, float64) {
	revenue := getUserInput("Revenue: ")
	for revenue < 1 {
		if revenue == 0 {
			fmt.Println("Invalid input. Revenue cannot be zero.")
		} else {
			fmt.Println("Invalid input. Revenue cannot be negative.")
		}
		fmt.Print("Please enter a valid revenue: ")
		fmt.Scan(&revenue)
	}
	expenses := getUserInput("Expenses: ")
	for expenses < 1 {
		if expenses == 0 {
			fmt.Println("Invalid input. Expenses cannot be zero.")
		} else {
			fmt.Println("Invalid input. Expenses cannot be negative.")
		}
		fmt.Print("Please enter a valid revenue: ")
		fmt.Scan(&expenses)
	}
	taxRate := getUserInput("Tax Rate: ")
	for taxRate < 1 {
		if taxRate == 0 {
			fmt.Println("Invalid input. Tax rate cannot be zero.")
		} else {
			fmt.Println("Invalid input. Tax rate cannot be negative.")
		}
		fmt.Print("Please enter a valid revenue: ")
		fmt.Scan(&taxRate)
	}

	return revenue, expenses, taxRate
}
