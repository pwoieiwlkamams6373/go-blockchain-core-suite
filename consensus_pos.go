package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Validator struct {
	Address string
	Stake   int64
}

type PoS struct {
	Validators []Validator
}

func NewPoS() *PoS {
	return &PoS{
		Validators: make([]Validator, 0),
	}
}

func (pos *PoS) RegisterValidator(addr string, stake int64) {
	pos.Validators = append(pos.Validators, Validator{
		Address: addr,
		Stake:   stake,
	})
}

func (pos *PoS) ElectValidator() *Validator {
	if len(pos.Validators) == 0 {
		return nil
	}

	totalStake := int64(0)
	for _, v := range pos.Validators {
		totalStake += v.Stake
	}

	rand.Seed(time.Now().UnixNano())
	value := rand.Int63n(totalStake)
	current := int64(0)

	for _, v := range pos.Validators {
		current += v.Stake
		if current > value {
			return &v
		}
	}
	return &pos.Validators[0]
}

func main() {
	pos := NewPoS()
	pos.RegisterValidator("node1", 100)
	pos.RegisterValidator("node2", 300)
	elect := pos.ElectValidator()
	fmt.Printf("当选验证节点: %s\n", elect.Address)
}
