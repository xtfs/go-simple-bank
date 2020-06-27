package main

import (
	"alura-banco/accounts"
	"alura-banco/customers"
	"fmt"
)

type customerBankAccount interface {
	Withdraw(value float64) bool
}

func doPaymentSlip(account customerBankAccount, value float64) {
	account.Withdraw(value)
}

func main() {
	owner1 := customers.Customer{"Tiago Farias", "123.123.123.12", "Desenvolvedor"}
	owner2 := customers.Customer{"Litane Farias", "321.321.321.32", "Desenvolvedora"}

	checkingAccount := accounts.CheckingAccount{Owner: owner1, AgencyCode: 100, AccountNumber: 100700}
	checkingAccount.Deposit(1000)
	checkingAccount2 := accounts.CheckingAccount{Owner: owner2, AgencyCode: 200, AccountNumber: 100800}
	checkingAccount2.Deposit(500)

	fmt.Println(checkingAccount)

	withdraw := checkingAccount.Withdraw(100)

	if withdraw {
		fmt.Println("Saque realizado com sucesso", checkingAccount.GetBalance())
	} else {
		fmt.Println("Saque não realizado", checkingAccount.GetBalance())
	}

	message, deposit := checkingAccount.Deposit(300)

	fmt.Println(message, deposit)

	checkingAccount.Transfer(500.30, &checkingAccount2)

	doPaymentSlip(&checkingAccount2, 150)

	fmt.Println(checkingAccount, checkingAccount2)
}
