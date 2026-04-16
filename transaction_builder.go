package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

type TxBuilder struct{}

func NewTxBuilder() *TxBuilder {
	return &TxBuilder{}
}

func (t *TxBuilder) NewTransferTx(from, to string, amount uint64) *Transaction {
	id := make([]byte, 16)
	rand.Read(id)
	return &Transaction{
		ID:        hex.EncodeToString(id),
		From:      from,
		To:        to,
		Amount:    amount,
		Timestamp: time.Now().Unix(),
		Status:    0,
	}
}

type Transaction struct {
	ID        string
	From      string
	To        string
	Amount    uint64
	Timestamp int64
	Status    int
}

func main() {
	builder := NewTxBuilder()
	tx := builder.NewTransferTx("addr1", "addr2", 100)
	fmt.Printf("交易ID: %s\n", tx.ID)
}
