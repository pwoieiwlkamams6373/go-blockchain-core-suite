package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

func GenerateECDSAKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

func SignData(priv *ecdsa.PrivateKey, data []byte) (r, s *big.Int, err error) {
	hash := sha256.Sum256(data)
	return ecdsa.Sign(rand.Reader, priv, hash[:])
}

func VerifyData(pub *ecdsa.PublicKey, data []byte, r, s *big.Int) bool {
	hash := sha256.Sum256(data)
	return ecdsa.Verify(pub, hash[:], r, s)
}

func main() {
	priv, _ := GenerateECDSAKey()
	data := []byte("block transaction data")
	r, s, _ := SignData(priv, data)
	valid := VerifyData(&priv.PublicKey, data, r, s)
	fmt.Printf("签名验证结果: %t\n", valid)
}
