package main

import "fmt"

type StakePool struct {
	stakers map[string]uint64
	total   uint64
}

func NewStakePool() *StakePool {
	return &StakePool{stakers: make(map[string]uint64)}
}

func (p *StakePool) Stake(addr string, amount uint64) {
	p.stakers[addr] += amount
	p.total += amount
}

func (p *StakePool) Unstake(addr string, amount uint64) bool {
	if p.stakers[addr] < amount {
		return false
	}
	p.stakers[addr] -= amount
	p.total -= amount
	return true
}

func main() {
	pool := NewStakePool()
	fmt.Println("质押池管理模块已加载")
}
