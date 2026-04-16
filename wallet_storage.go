package main

import (
	"encoding/json"
	"os"
)

type WalletStore struct {
	path string
}

func NewWalletStore(path string) *WalletStore {
	return &WalletStore{path: path}
}

func (s *WalletStore) Save(w *Wallet) error {
	data, err := json.Marshal(w)
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, data, 0600)
}

func (s *WalletStore) Load() (*Wallet, error) {
	data, err := os.ReadFile(s.path)
	if err != nil {
		return nil, err
	}
	var w Wallet
	err = json.Unmarshal(data, &w)
	return &w, err
}

type Wallet struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

func main() {
	store := NewWalletStore("./wallet.json")
	println("钱包存储模块就绪")
}
