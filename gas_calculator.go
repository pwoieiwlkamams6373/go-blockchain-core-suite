package main

import "fmt"

type GasCalculator struct {
	baseGas uint64
}

func NewGasCalculator() *GasCalculator {
	return &GasCalculator{baseGas: 21000}
}

func (g *GasCalculator) CalcTxGas(dataSize int) uint64 {
	return g.baseGas + uint64(dataSize)*100
}

func (g *GasCalculator) CalcContractGas(opCount int) uint64 {
	return uint64(opCount) * 500
}

func main() {
	calc := NewGasCalculator()
	fmt.Println("Gas费用计算器就绪")
}
