package main

import (
	"fmt"
	"log"
	"os"
)

type LogCollector struct {
	file *os.File
}

func NewLogCollector(path string) *LogCollector {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return &LogCollector{file: f}
}

func (l *LogCollector) Info(msg string) {
	log.SetOutput(l.file)
	log.Println("[INFO]", msg)
}

func (l *LogCollector) Error(msg string) {
	log.SetOutput(l.file)
	log.Println("[ERROR]", msg)
}

func main() {
	logger := NewLogCollector("./chain.log")
	logger.Info("节点启动成功")
	fmt.Println("日志收集服务已启动")
}
