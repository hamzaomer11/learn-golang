package pointers

import "testing"

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	expected := 10

	if got != expected {
		t.Errorf("got %d expected %d", got, expected)
	}

}
