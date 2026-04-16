package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

func ValidateBlockHash(block *ChainBlock) bool {
	record := strconv.Itoa(block.Index) +
		strconv.FormatInt(block.Timestamp, 10) +
		string(block.Data) +
		string(block.PrevHash) +
		strconv.Itoa(block.Nonce)

	hash := sha256.Sum256([]byte(record))
	generatedHash := hex.EncodeToString(hash[:])
	return generatedHash == string(block.Hash)
}

func ValidateBlockChain(blocks []*ChainBlock) bool {
	for i := 1; i < len(blocks); i++ {
		currentBlock := blocks[i]
		prevBlock := blocks[i-1]

		if !ValidateBlockHash(currentBlock) {
			return false
		}

		if string(currentBlock.PrevHash) != string(prevBlock.Hash) {
			return false
		}
	}
	return true
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
	fmt.Println("区块验证工具已加载")
}
