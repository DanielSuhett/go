package account

import (
	client "bank/client"
	"math/rand"
)

type account struct {
	owner         client.Owner
	bankNumber    int
	accountNumber int
	balance       float64
}

func Create(owner client.Owner) account {
	balance := 0.
	bankNumber := rand.Int()
	accountNumber := rand.Int()

	return account{owner, bankNumber, accountNumber, balance}
}

func (c *account) Deposit(value float64) (string, float64) {
	canDeposit := value > 0

	if canDeposit {
		c.balance += value
		return "Successfully deposited", c.balance
	} else {
		return "Not deposited", c.balance
	}
}

func (c *account) Withdraw(value float64) (string, float64) {
	canWithdraw := value <= c.balance

	if canWithdraw {
		c.balance -= value
		return "Successfully withdrawn", c.balance
	} else {
		return "Not withdrawn", c.balance
	}
}

func (c *account) Transfer(value float64, payee *account) bool {
	canTransfer := c.balance > value && value > 0
	if canTransfer {
		c.balance -= value
		payee.Deposit(value)
		return true
	} else {
		return false
	}
}

func (c *account) GetBalance() float64 {
	return c.balance
}
