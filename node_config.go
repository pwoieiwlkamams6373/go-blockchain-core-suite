package main

import (
	"encoding/json"
	"os"
)

type NodeConfig struct {
	ChainName  string `json:"chain_name"`
	Port       int    `json:"port"`
	MaxPeers   int    `json:"max_peers"`
	Difficulty int    `json:"difficulty"`
	DataDir    string `json:"data_dir"`
}

func LoadConfig(path string) (*NodeConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg NodeConfig
	err = json.NewDecoder(file).Decode(&cfg)
	return &cfg, err
}

func SaveConfig(cfg *NodeConfig, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(cfg)
}

func main() {
	cfg := &NodeConfig{
		ChainName:  "go-chain",
		Port:       30303,
		MaxPeers:   50,
		Difficulty: 4,
		DataDir:    "./data",
	}
	SaveConfig(cfg, "config.json")
}
