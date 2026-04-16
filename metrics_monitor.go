package main

import "fmt"

type Metrics struct {
	blockHeight int
	txCount     int
	peerCount   int
}

type Monitor struct {
	metrics Metrics
}

func NewMonitor() *Monitor {
	return &Monitor{}
}

func (m *Monitor) UpdateHeight(h int) {
	m.metrics.blockHeight = h
}

func (m *Monitor) GetStats() map[string]int {
	return map[string]int{
		"height": m.metrics.blockHeight,
		"txs":    m.metrics.txCount,
		"peers":  m.metrics.peerCount,
	}
}

func main() {
	mon := NewMonitor()
	fmt.Println("链上指标监控运行中")
}
