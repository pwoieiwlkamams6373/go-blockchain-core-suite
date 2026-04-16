package main

import "fmt"

type BlockReward struct {
	baseReward uint64
}

func NewBlockReward() *BlockReward {
	return &BlockReward{baseReward: 10}
}

func (r *BlockReward) Calculate(height int) uint64 {
	if height < 100000 {
		return r.baseReward
	}
	return r.baseReward / 2
}

func (r *BlockReward) Distribute(validator string, amount uint64) {
	fmt.Printf("向 %s 发放奖励: %d\n", validator, amount)
}

func main() {
	reward := NewBlockReward()
	fmt.Println("区块奖励分发模块就绪")
}
