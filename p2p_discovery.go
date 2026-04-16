package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Discovery struct {
	localID  string
	bootNode []string
}

func NewDiscovery(localID string) *Discovery {
	return &Discovery{
		localID:  localID,
		bootNode: []string{"boot1", "boot2", "boot3"},
	}
}

func (d *Discovery) FindPeers() []string {
	rand.Seed(time.Now().UnixNano())
	res := make([]string, 0)
	for i := 0; i < 3; i++ {
		res = append(res, fmt.Sprintf("peer-%d", rand.Intn(1000)))
	}
	return res
}

func (d *Discovery) Ping() bool {
	return true
}

func main() {
	disc := NewDiscovery("node-local")
	peers := disc.FindPeers()
	fmt.Printf("发现节点: %v\n", peers)
}
