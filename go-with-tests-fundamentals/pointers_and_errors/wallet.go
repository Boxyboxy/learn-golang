package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) { // this change makes the method receive a pointer to the wallet struct instead of a copy
	w.balance += amount
}

func (w Wallet) Balance() Bitcoin {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance) //Since this is a copy of the struct, the address will be different from using a pointer to the struct

	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds") // var keyword defines values global to the package

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds // refactoring of errors
	}
	w.balance -= amount
	return nil //errors can be nil because error is an interface.
	// Any function that takes arguments or returns values that are interfaces, they can be nillable
}
