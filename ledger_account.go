package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Account struct {
	Address string
	Balance uint64
	Nonce   uint64
}

type AccountLedger struct {
	accounts map[string]*Account
	mutex    sync.RWMutex
}

func NewLedger() *AccountLedger {
	return &AccountLedger{
		accounts: make(map[string]*Account),
	}
}

func (l *AccountLedger) CreateAccount(addr string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.accounts[addr] = &Account{
		Address: addr,
		Balance: 0,
		Nonce:   0,
	}
}

func (l *AccountLedger) AddBalance(addr string, amount uint64) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	acc, ok := l.accounts[addr]
	if !ok {
		return fmt.Errorf("账户不存在")
	}
	acc.Balance += amount
	return nil
}

func (l *AccountLedger) ToJSON() ([]byte, error) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	return json.Marshal(l.accounts)
}

func main() {
	ledger := NewLedger()
	ledger.CreateAccount("wallet1")
	ledger.AddBalance("wallet1", 1000)
	fmt.Println("账户账本已初始化")
}
