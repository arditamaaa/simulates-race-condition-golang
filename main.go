package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance float64
	mutex   sync.Mutex
}

func (b *BankAccount) Deposit(amount float64) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if amount > 0 {
		b.balance += amount
		fmt.Printf("Setor %.2f. Saldo baru: %.2f\n", amount, b.balance)
	} else {
		fmt.Println("Jumlah setoran harus positif.")
	}
}

func (b *BankAccount) Withdraw(amount float64) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if amount > 0 && b.balance >= amount {
		b.balance -= amount
		fmt.Printf("Tarik %.2f. Saldo baru: %.2f\n", amount, b.balance)
		return nil
	} else if amount <= 0 {
		return fmt.Errorf("jumlah penarikan harus positif")
	} else {
		return fmt.Errorf("saldo tidak mencukupi")
	}
}

func (b *BankAccount) GetBalance() float64 {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.balance
}

func simulatesTransactions(rekening *BankAccount, transaksi []struct {
	operation string
	amount    float64
}, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, t := range transaksi {
		if t.operation == "deposit" {
			rekening.Deposit(t.amount)
		} else if t.operation == "withdraw" {
			rekening.Withdraw(t.amount)
		} else {
			fmt.Printf("Transaksi tidak valid: %s\n", t.operation)
		}
	}
}

func main() {
	account := BankAccount{balance: 2000}

	transactions1 := []struct {
		operation string
		amount    float64
	}{
		{"deposit", 200},
		{"withdraw", 100},
		{"withdraw", 200},
	}

	transactions2 := []struct {
		operation string
		amount    float64
	}{
		{"withdraw", 500},
		{"deposit", 1500},
		{"withdraw", 200},
	}

	fmt.Printf("Saldo Awal: %.2f\n", account.balance)
	var wg sync.WaitGroup
	wg.Add(2)

	go simulatesTransactions(&account, transactions1, &wg)
	go simulatesTransactions(&account, transactions2, &wg)

	wg.Wait()
	fmt.Printf("Saldo Akhir: %.2f\n", account.GetBalance())
}
