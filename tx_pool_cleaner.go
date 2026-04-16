package main

import (
	"time"
)

type TxCleaner struct {
	interval time.Duration
	mempool  *Mempool
}

func NewTxCleaner(m *Mempool) *TxCleaner {
	return &TxCleaner{
		interval: 30 * time.Second,
		mempool:  m,
	}
}

func (c *TxCleaner) Start() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		txs := c.mempool.GetAllTxs()
		for _, tx := range txs {
			if time.Now().Unix()-tx.Timestamp > 3600 {
				c.mempool.RemoveTx(tx.ID)
			}
		}
	}
}

func main() {
	pool := NewMempool()
	cleaner := NewTxCleaner(pool)
	go cleaner.Start()
}
