package accounts

import "alura-banco/customers"

type CheckingAccount struct {
	Owner                     customers.Customer
	AgencyCode, AccountNumber int
	balance                   float64
}

func (c *CheckingAccount) Withdraw(value float64) bool {
	if value > 0 && c.balance >= value {
		c.balance -= value
		return true
	}

	return false
}

func (c *CheckingAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		c.balance += value
		return "Deposito realizado com sucesso", c.balance
	}

	return "Depósito não realizado", c.balance
}

func (c *CheckingAccount) Transfer(value float64, destinyCheckingAccount *CheckingAccount) bool {
	if value > 0 && c.balance >= value {
		c.balance -= value
		destinyCheckingAccount.Deposit(value)
		return true
	}
	return false
}

func (c *CheckingAccount) GetBalance() float64 {
	return c.balance
}
