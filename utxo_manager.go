package main

import "fmt"

type UTXO struct {
	TxID  string
	Index int
	Addr  string
	Value uint64
	Spent bool
}

type UTXOSet struct {
	utxos map[string]*UTXO
}

func NewUTXOSet() *UTXOSet {
	return &UTXOSet{utxos: make(map[string]*UTXO)}
}

func (u *UTXOSet) AddUTXO(utxo *UTXO) {
	key := utxo.TxID + fmt.Sprintf("%d", utxo.Index)
	u.utxos[key] = utxo
}

func (u *UTXOSet) SpendUTXO(txID string, index int) {
	key := txID + fmt.Sprintf("%d", index)
	if utxo, ok := u.utxos[key]; ok {
		utxo.Spent = true
	}
}

func (u *UTXOSet) GetBalance(addr string) uint64 {
	bal := uint64(0)
	for _, utxo := range u.utxos {
		if utxo.Addr == addr && !utxo.Spent {
			bal += utxo.Value
		}
	}
	return bal
}

func main() {
	set := NewUTXOSet()
	fmt.Println("UTXO管理模块就绪")
}
