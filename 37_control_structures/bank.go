package main

import (
	"fmt"
)

func main() {
	var accountBalance float64 = 1000
	var addMoney float64
	var subMoney float64
	var choice int

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Printf("Your account balance is $%.2f\n", accountBalance)
		} else if choice == 2 {
			addMoney = depositMoney()
			if addMoney != 0 {
				accountBalance += addMoney
				fmt.Println("Balance update! New ammount: $", accountBalance)
			} else {
				fmt.Println("No money was added.")
			}
		} else if choice == 3 {
			subMoney = withdrawMoney(accountBalance)
			if subMoney != 0 {
				accountBalance -= subMoney
				fmt.Println("Balance update! New ammount: $", accountBalance)
			} else {
				fmt.Println("No money was withdrawn.")
			}
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}
	fmt.Println("Thanks for choosing our bank")
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
