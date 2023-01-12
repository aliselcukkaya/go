package wallet

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, want Bitcoin, wallet Wallet) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, want, wallet)
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, want, wallet)
	})
}
