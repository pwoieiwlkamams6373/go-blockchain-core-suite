package main

import (
	"sync"
	"time"
)

type Transaction struct {
	ID        string
	Data      []byte
	Timestamp int64
}

type Mempool struct {
	pool map[string]*Transaction
	mu   sync.RWMutex
}

func NewMempool() *Mempool {
	return &Mempool{
		pool: make(map[string]*Transaction),
	}
}

func (m *Mempool) AddTx(tx *Transaction) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.pool[tx.ID] = tx
}

func (m *Mempool) RemoveTx(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.pool, id)
}

func (m *Mempool) GetAllTxs() []*Transaction {
	m.mu.RLock()
	defer m.mu.RUnlock()
	list := make([]*Transaction, 0, len(m.pool))
	for _, tx := range m.pool {
		list = append(list, tx)
	}
	return list
}

func main() {
	pool := NewMempool()
	pool.AddTx(&Transaction{ID: "tx001", Data: []byte("transfer"), Timestamp: time.Now().Unix()})
}
