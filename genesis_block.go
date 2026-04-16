package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Genesis struct {
	Block ChainBlock
}

func CreateGenesisBlock() *ChainBlock {
	timestamp := time.Now().Unix()
	data := []byte("chain init genesis block")
	prevHash := []byte{}
	nonce := 0

	record := strconv.Itoa(0) +
		strconv.FormatInt(timestamp, 10) +
		string(data) +
		string(prevHash) +
		strconv.Itoa(nonce)
	hash := sha256.Sum256([]byte(record))

	return &ChainBlock{
		Index:     0,
		Timestamp: timestamp,
		Data:      data,
		PrevHash:  prevHash,
		Hash:      []byte(hex.EncodeToString(hash[:])),
		Nonce:     nonce,
	}
}

type ChainBlock struct {
	Index     int
	Timestamp int64
	Data      []byte
	PrevHash  []byte
	Hash      []byte
	Nonce     int
}

func main() {
	genesis := CreateGenesisBlock()
	fmt.Printf("创世区块高度: %d\n", genesis.Index)
}
