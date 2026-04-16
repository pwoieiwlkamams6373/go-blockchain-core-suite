package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func ComputeMerkleRoot(transactions [][]byte) string {
	if len(transactions) == 0 {
		return ""
	}

	var hashes [][]byte
	for _, t := range transactions {
		hash := sha256.Sum256(t)
		hashes = append(hashes, hash[:])
	}

	for len(hashes) > 1 {
		var newLevel [][]byte
		for i := 0; i < len(hashes); i += 2 {
			if i+1 >= len(hashes) {
				newLevel = append(newLevel, hashes[i])
				continue
			}
			combined := append(hashes[i], hashes[i+1]...)
			hash := sha256.Sum256(combined)
			newLevel = append(newLevel, hash[:])
		}
		hashes = newLevel
	}
	return hex.EncodeToString(hashes[0])
}

func main() {
	txs := [][]byte{[]byte("tx1"), []byte("tx2"), []byte("tx3")}
	root := ComputeMerkleRoot(txs)
	fmt.Printf("默克尔根: %s\n", root)
}
