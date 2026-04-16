package main

import (
	"fmt"
	"net"
	"sync"
)

type Peer struct {
	ID      string
	Address string
	conn    net.Conn
}

type PeerManager struct {
	peers map[string]*Peer
	mu    sync.RWMutex
}

func NewPeerManager() *PeerManager {
	return &PeerManager{
		peers: make(map[string]*Peer),
	}
}

func (pm *PeerManager) AddPeer(peer *Peer) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.peers[peer.ID] = peer
}

func (pm *PeerManager) RemovePeer(id string) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	delete(pm.peers, id)
}

func (pm *PeerManager) Broadcast(msg []byte) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	for _, p := range pm.peers {
		if p.conn != nil {
			p.conn.Write(msg)
		}
	}
}

func main() {
	fmt.Println("P2P节点管理器已启动")
}
