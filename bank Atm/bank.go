package main

import (
	"fmt"
)

func main() {
	var accountBalance float64 = 4500
	fmt.Println("Welcome to you personal  Bank account  ")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {

		fmt.Println("Your account balance is", accountBalance)
	} else if choice == 2 {
		fmt.Print("how much mooney you want to deposit:")
		var Deposit float64
		fmt.Scan(&Deposit)
		if Deposit <= 0 {
			fmt.Print("Invalid amount! The amount must be greater to 0")
		} else {

			accountBalance = accountBalance + Deposit
			fmt.Printf("Your new account balance is %v", accountBalance)
		}
	} else if choice == 3 {
		fmt.Print("how much mooney you want to withdraw:")
		var withdraw float64
		fmt.Scan(&withdraw)
		if withdraw > accountBalance || withdraw < 0 {
			fmt.Println("insufficient funds")
			fmt.Println("Your account balance is :", accountBalance)
		} else {
			accountBalance = accountBalance - withdraw
			fmt.Printf("Your new account balance is %v", accountBalance)
		}
	} else if choice == 4 {

		fmt.Print("Goodbye")
	} else {
		fmt.Print("Invalid choice")
	}

}
