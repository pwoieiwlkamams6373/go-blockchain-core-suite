package main

import "fmt"

type ContractExecutor struct {
	gasLimit uint64
}

func NewExecutor(gas uint64) *ContractExecutor {
	return &ContractExecutor{gasLimit: gas}
}

func (e *ContractExecutor) Execute(code []byte, input []byte) ([]byte, uint64, error) {
	usedGas := uint64(21000)
	if usedGas > e.gasLimit {
		return nil, 0, fmt.Errorf("gas不足")
	}
	return []byte("exec success"), usedGas, nil
}

func main() {
	exec := NewExecutor(1000000)
	fmt.Println("合约虚拟机执行器就绪")
}
