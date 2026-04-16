package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
	Address    string
}

func NewWallet() (*Wallet, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	pub := elliptic.Marshal(priv.Curve, priv.X, priv.Y)
	addr := hex.EncodeToString(pub)[:40]
	return &Wallet{
		PrivateKey: priv,
		PublicKey:  pub,
		Address:    addr,
	}, nil
}

func main() {
	wallet, _ := NewWallet()
	fmt.Printf("钱包地址: %s\n", wallet.Address)
}
