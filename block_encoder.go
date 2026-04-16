package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

type ChainBlock struct {
	Index     int
	Timestamp int64
	Data      []byte
	PrevHash  []byte
	Hash      []byte
	Nonce     int
}

func EncodeBlock(block *ChainBlock) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(block)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func DecodeBlock(data []byte) (*ChainBlock, error) {
	var block ChainBlock
	decoder := gob.NewDecoder(bytes.NewBuffer(data))
	err := decoder.Decode(&block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

func main() {
	testBlock := &ChainBlock{
		Index:     1,
		Timestamp: time.Now().Unix(),
		Data:      []byte("genesis block data"),
		PrevHash:  []byte{},
		Hash:      []byte("genesis-hash"),
		Nonce:     0,
	}

	encoded, _ := EncodeBlock(testBlock)
	decodedBlock, _ := DecodeBlock(encoded)
	fmt.Printf("解码后区块索引: %d\n", decodedBlock.Index)
}
