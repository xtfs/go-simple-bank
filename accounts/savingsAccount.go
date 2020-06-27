package accounts

import "alura-banco/customers"

type SavingsAccount struct {
	Owner                                customers.Customer
	AgencyCode, AccountNumber, Operation int
	balance                              float64
}

func (s *SavingsAccount) Withdraw(value float64) bool {
	if value > 0 && s.balance >= value {
		s.balance -= value
		return true
	}

	return false
}

func (s *SavingsAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		s.balance += value
		return "Deposito realizado com sucesso", s.balance
	}

	return "Depósito não realizado", s.balance
}

func (s *SavingsAccount) GetBalance() float64 {
	return s.balance
}
