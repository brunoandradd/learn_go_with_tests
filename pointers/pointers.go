package pointers

import (
	"errors"
	"fmt"
)

var p = fmt.Printf

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(value Bitcoin) {
	p("wallet in deposit %p \n", w)
	w.balance += value
}

var ErrInsufficientFunds = errors.New("cannot withdraw, without funds")

func (w *Wallet) Withdraw(value Bitcoin) error {
	if value > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= value
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	p("wallet in balance %p \n", w)
	return w.balance
}
