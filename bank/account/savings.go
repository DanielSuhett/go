package account

import (
	client "bank/client"
	"math/rand"
)

type savings struct {
	owner                                client.Owner
	bankNumber, accountNumber, operation int
	balance                              float64
}

func CreateSavings(owner client.Owner) savings {
	balance := 0.
	bankNumber := rand.Int()
	accountNumber := rand.Int()
	operation := rand.Int()

	return savings{owner, bankNumber, accountNumber, operation, balance}
}

func (c *savings) Deposit(value float64) (string, float64) {
	canDeposit := value > 0

	if canDeposit {
		c.balance += value
		return "Successfully deposited", c.balance
	} else {
		return "Not deposited", c.balance
	}
}

func (c *savings) Withdraw(value float64) (string, float64) {
	canWithdraw := value <= c.balance

	if canWithdraw {
		c.balance -= value
		return "Successfully withdrawn", c.balance
	} else {
		return "Not withdrawn", c.balance
	}
}

func (c *savings) Transfer(value float64, payee *account) bool {
	canTransfer := c.balance > value && value > 0
	if canTransfer {
		c.balance -= value
		payee.Deposit(value)
		return true
	} else {
		return false
	}
}

func (c *savings) GetBalance() float64 {
	return c.balance
}
