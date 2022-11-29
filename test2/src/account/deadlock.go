package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Account struct {
	balance int
	mutex   *sync.Mutex
}

func (a *Account) Widthdraw(val int) {
	a.balance -= val
}

func (a *Account) Deposit(val int) {
	a.balance += val
}

func (a *Account) Balance() int {
	balance := a.balance
	return balance
}

var accounts []Account

func Transfer(sender, receiver, money int) {
	accounts[sender].mutex.Lock()
	fmt.Println("Lock: ", sender)
	time.Sleep(1000 * time.Millisecond)
	accounts[receiver].mutex.Lock()
	fmt.Println("Lock: ", receiver)

	accounts[sender].Widthdraw(money)
	accounts[receiver].Deposit(money)

	accounts[receiver].mutex.Unlock()
	accounts[sender].mutex.Unlock()

	fmt.Println(sender, receiver, money)
}

func GetTotalbalance() int {
	total := 0
	for i := 0; i < len(accounts); i++ {
		total += accounts[i].Balance()
	}
	return total
}

func RandomTrasfer() {
	var sender, balance int
	for {
		sender = rand.Intn(len(accounts))
		balance = accounts[sender].Balance()
		if balance > 0 {
			break
		}
	}

	var receiver int
	for {
		receiver = rand.Intn(len(accounts))
		if sender != receiver {
			break
		}
	}

	money := rand.Intn(balance)
	Transfer(sender, receiver, money)
}

func GoTransfer() {
	for {
		RandomTrasfer()
	}
}

func PrintToTalBalance() {
	fmt.Printf("Total: %d\n", GetTotalbalance())
}

func main() {
	for i := 0; i < 20; i++ {
		accounts = append(accounts, Account{balance: 1000, mutex: &sync.Mutex{}})
	}

	go func() {
		for {
			Transfer(0, 1, 100)
		}
	}()

	go func() {
		for {
			Transfer(1, 0, 100)
		}
	}()

	for {
		PrintToTalBalance()
		time.Sleep(100 * time.Millisecond)
	}
}
