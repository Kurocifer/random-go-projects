package main

import (
	"fmt"
	"sync"

	"github.com/kurocifer/randomgop/concurrent_back_account/account"
)

func main() {
	acc := account.NewAccount(0)

	var wg sync.WaitGroup

	wg.Add(5)

	go func() {
		defer wg.Done()
		acc.Deposit(250)
		acc.Balance()
	}()

	go func() {
		defer wg.Done()
		err := acc.Withdraw(200)
		if err != nil {
			fmt.Println(err.Error())
		}
		acc.Balance()
	}()

	go func() {
		defer wg.Done()
		acc.Deposit(1500)
		acc.Balance()
	}()

	go func() {
		defer wg.Done()
		acc.Deposit(600)
		acc.Balance()
	}()

	go func() {
		defer wg.Done()
		err := acc.Withdraw(8000)
		if err != nil {
			fmt.Println(err)
		}
		acc.Balance()
	}()

	wg.Wait()
	fmt.Printf("Seems everyone is done doing what ever. Final account balance: ")
	acc.Balance()
}
