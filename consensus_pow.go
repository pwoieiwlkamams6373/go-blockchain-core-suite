package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

const Difficulty = 4

type PoWBlock struct {
	Index     int
	Timestamp int64
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int
}

func calculateHash(b PoWBlock) string {
	record := strconv.Itoa(b.Index) +
		strconv.FormatInt(b.Timestamp, 10) +
		b.Data +
		b.PrevHash +
		strconv.Itoa(b.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func mineBlock(b PoWBlock) PoWBlock {
	target := string(make([]byte, Difficulty))
	for i := 0; ; i++ {
		b.Nonce = i
		hash := calculateHash(b)
		if hash[:Difficulty] == target {
			b.Hash = hash
			break
		}
	}
	return b
}

func main() {
	genesisBlock := PoWBlock{
		Index:     0,
		Timestamp: time.Now().Unix(),
		Data:      "genesis block",
		PrevHash:  "",
	}
	genesisBlock = mineBlock(genesisBlock)
	fmt.Printf("创世区块哈希: %s\n", genesisBlock.Hash)
}
