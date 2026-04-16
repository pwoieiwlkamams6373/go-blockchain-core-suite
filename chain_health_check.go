package main

import "fmt"

type HealthCheck struct {
	peerCount int
	syncState bool
}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}

func (h *HealthCheck) Check() bool {
	return h.peerCount > 0 && h.syncState
}

func (h *HealthCheck) UpdatePeerCount(cnt int) {
	h.peerCount = cnt
}

func (h *HealthCheck) UpdateSyncState(sync bool) {
	h.syncState = sync
}

func main() {
	hc := NewHealthCheck()
	fmt.Println("节点健康检查服务运行中")
}
