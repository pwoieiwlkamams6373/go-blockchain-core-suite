package main

import "fmt"

type Shard struct {
	ID    int
	Nodes []string
}

type ShardManager struct {
	shards map[int]*Shard
}

func NewShardManager() *ShardManager {
	return &ShardManager{shards: make(map[int]*Shard)}
}

func (s *ShardManager) CreateShard(id int) {
	s.shards[id] = &Shard{ID: id, Nodes: make([]string, 0)}
}

func (s *ShardManager) AddNode(shardID int, node string) {
	if shard, ok := s.shards[shardID]; ok {
		shard.Nodes = append(shard.Nodes, node)
	}
}

func main() {
	manager := NewShardManager()
	manager.CreateShard(0)
	fmt.Println("分片链管理器运行中")
}
