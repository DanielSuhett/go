package main

import (
	accounts "bank/account"
	client "bank/client"
	"fmt"
)

type isAccount interface {
	Withdraw(float64) (string, float64)
}

func Billing(acc isAccount, value float64) (string, float64) {
	return acc.Withdraw(value)
}

func AccCreateAndBilling() {
	owner := client.Create("Daniel", "010.101.101-11", "Software Developer")

	acc := accounts.Create(owner)

	acc.Deposit(100)

	_, balance := Billing(&acc, 60)

	fmt.Printf("CHECKING ACCOUNT: Successfully billing, balance is: %g\n", balance)
}

func SavingsCreateAndBilling() {
	owner := client.Create("Daniel", "010.101.101-11", "Software Developer")

	acc := accounts.CreateSavings(owner)

	acc.Deposit(300)

	_, balance := Billing(&acc, 120)

	fmt.Printf("SAVINGS ACCOUNT: Successfully billing, balance is: %g\n", balance)
}

func main() {
	AccCreateAndBilling()
	SavingsCreateAndBilling()
}
