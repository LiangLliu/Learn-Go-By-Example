package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

/*
实现 Bitcoin 的 Stringer 方法
*/
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// Deposit * : 指针让我们 指向 某个值，然后修改它
func (wallet *Wallet) Deposit(amount Bitcoin) {
	fmt.Println("address of balance in Deposit is", &wallet.balance)
	wallet.balance += amount
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if amount > wallet.balance {
		return InsufficientFundsError
	}

	wallet.balance -= amount
	return nil
}
