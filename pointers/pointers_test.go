package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, err error, want string) {
		t.Helper()
		if err == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if err.Error() != want {
			t.Errorf("got %q, want %q", err, want)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		p("wallet test: %p \n", &wallet)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficent funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, "cannot withdraw, without funds")
		assertBalance(t, wallet, Bitcoin(20))
	})
}


// continue here
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors#unchecked-errors
