package main

import "fmt"

type ChainSyncer struct {
	localHeight  int
	remoteHeight int
}

func NewChainSyncer() *ChainSyncer {
	return &ChainSyncer{}
}

func (s *ChainSyncer) Sync(local, remote int) bool {
	if remote > local {
		s.localHeight = remote
		return true
	}
	return false
}

func (s *ChainSyncer) RequestBlocks(from, to int) []int {
	res := make([]int, 0)
	for i := from; i <= to; i++ {
		res = append(res, i)
	}
	return res
}

func main() {
	syncer := NewChainSyncer()
	fmt.Println("链数据同步服务运行中")
}
