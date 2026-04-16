package main

import "fmt"

type ChainRouter struct {
	chains map[string]ChainGateway
}

type ChainGateway interface {
	SendMessage(data []byte) error
}

type EthereumGateway struct{}
type BSCGateway struct{}

func (e *EthereumGateway) SendMessage(data []byte) error {
	fmt.Printf("以太坊跨链消息: %s\n", string(data))
	return nil
}

func (b *BSCGateway) SendMessage(data []byte) error {
	fmt.Printf("BSC跨链消息: %s\n", string(data))
	return nil
}

func NewChainRouter() *ChainRouter {
	return &ChainRouter{
		chains: make(map[string]ChainGateway),
	}
}

func (r *ChainRouter) RegisterChain(chain string, gateway ChainGateway) {
	r.chains[chain] = gateway
}

func (r *ChainRouter) Route(chain string, data []byte) error {
	gw, ok := r.chains[chain]
	if !ok {
		return fmt.Errorf("链不存在")
	}
	return gw.SendMessage(data)
}

func main() {
	router := NewChainRouter()
	router.RegisterChain("eth", &EthereumGateway{})
	router.RegisterChain("bsc", &BSCGateway{})
	router.Route("eth", []byte("cross chain data"))
}
