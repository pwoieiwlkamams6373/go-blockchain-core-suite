package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type TxSigner struct {
	privateKey string
}

func NewTxSigner(priv string) *TxSigner {
	return &TxSigner{privateKey: priv}
}

func (s *TxSigner) SignTx(tx *Transaction) string {
	data := tx.From + tx.To + fmt.Sprintf("%d", tx.Amount) + fmt.Sprintf("%d", tx.Timestamp)
	hash := sha256.Sum256([]byte(data + s.privateKey))
	return hex.EncodeToString(hash[:])
}

func (s *TxSigner) VerifySig(tx *Transaction, sig string) bool {
	data := tx.From + tx.To + fmt.Sprintf("%d", tx.Amount) + fmt.Sprintf("%d", tx.Timestamp)
	hash := sha256.Sum256([]byte(data + s.privateKey))
	return hex.EncodeToString(hash[:]) == sig
}

type Transaction struct {
	From      string
	To        string
	Amount    uint64
	Timestamp int64
}

func main() {
	signer := NewTxSigner("priv-key-demo")
	fmt.Println("交易签名器就绪")
}
