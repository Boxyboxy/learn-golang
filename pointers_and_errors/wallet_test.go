package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		fmt.Printf("address of balance in test is %p \n", &wallet.balance) //&myVal stores the address of the variables

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, "cannot withdraw, insufficient funds")
		assertBalance(t, wallet, startingBalance)

	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	// t.Fatal will stop the test if it is called
	// We don't want to make any assertions on the error returned if there isn't one around
	// Without this the test would carry on to the next step and panic because of a nil pointer

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
