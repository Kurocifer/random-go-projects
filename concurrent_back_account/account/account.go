package account

import (
	"fmt"
	"sync"
)

type Account struct {
	mu      sync.Mutex
	balance float64
}

func NewAccount(balance float64) *Account {
	return &Account{
		balance: balance,
	}
}

func (a *Account) Deposit(ammount float64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += ammount
	fmt.Printf("Deposited %.2f\n", ammount)
}

func (a *Account) Withdraw(ammount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.balance < ammount {
		return fmt.Errorf("could not withdraw %.2f, account balance not enough. consult balance", ammount)
	}

	a.balance -= ammount
	fmt.Printf("Withdrawn %.2f\n", ammount)
	return nil
}

func (a *Account) Balance() {
	a.mu.Lock()
	defer a.mu.Unlock()

	fmt.Printf("Balance: %.2f cfa\n", a.balance)
}
