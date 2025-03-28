package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	var balanceBytes []byte
	var err error
	var balanceText string
	var balance float64

	balanceBytes, err = os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}
	balanceText = string(balanceBytes)
	balance, err = strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance float64
	var addMoney float64
	var subMoney float64
	var choice int
	var err error

	accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR:")
		fmt.Println(err)
		fmt.Print("---------------------")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		for choice < 1 || choice > 4 {
			fmt.Print("Invalid choice. Please enter a valid choice: ")
			fmt.Scan(&choice)
		}

		switch choice {
		case 1:
			fmt.Printf("Your account balance is $%.2f\n", accountBalance)
		case 2:
			addMoney = depositMoney()
			if addMoney != 0 {
				accountBalance += addMoney
				fmt.Println("Balance update! New ammount: $", accountBalance)
				writeBalanceToFile(accountBalance)
			} else {
				fmt.Println("No money was added.")
			}
		case 3:
			subMoney = withdrawMoney(accountBalance)
			if subMoney != 0 {
				accountBalance -= subMoney
				fmt.Println("Balance update! New ammount: $", accountBalance)
				writeBalanceToFile(accountBalance)
			} else {
				fmt.Println("No money was withdrawn.")
			}
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}
}

func depositMoney() float64 {
	var depositAmount float64
	fmt.Print("Enter the amount you want to deposit: ")
	fmt.Scan(&depositAmount)
	if depositAmount <= 0 {
		depositAmount = 0
		fmt.Println("Invalid amount. Must be greater than 0.")
	} else {
		fmt.Println("Depositing $", depositAmount)
	}
	return depositAmount
}

func withdrawMoney(accountBalance float64) float64 {
	var withdrawAmount float64
	fmt.Print("Enter the amount you want to withdraw: ")
	fmt.Scan(&withdrawAmount)
	if withdrawAmount > accountBalance {
		withdrawAmount = 0
		fmt.Println("Insufficient funds")
	} else if withdrawAmount <= 0 {
		withdrawAmount = 0
		fmt.Println("Invalid amount. Must be greater than 0.")
	} else {
		fmt.Println("Withdrawing $", withdrawAmount)
	}
	return withdrawAmount
}
