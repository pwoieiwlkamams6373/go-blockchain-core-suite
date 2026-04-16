package main

import "fmt"

type BlockVerifier struct {
	difficulty int
}

func NewBlockVerifier(d int) *BlockVerifier {
	return &BlockVerifier{difficulty: d}
}

func (v *BlockVerifier) CheckDifficulty(hash string) bool {
	prefix := hash[:v.difficulty]
	for _, c := range prefix {
		if c != '0' {
			return false
		}
	}
	return true
}

func (v *BlockVerifier) FullVerify(b *ChainBlock) bool {
	return v.CheckDifficulty(string(b.Hash))
}

type ChainBlock struct {
	Hash []byte
}

func main() {
	verifier := NewBlockVerifier(4)
	fmt.Println("区块难度验证器已启动")
}
