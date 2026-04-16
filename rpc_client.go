package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RPCClient struct {
	url string
}

func NewRPCClient(url string) *RPCClient {
	return &RPCClient{url: url}
}

type RPCRequest struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
	ID     int         `json:"id"`
}

func (c *RPCClient) Call(method string, params interface{}) (map[string]interface{}, error) {
	reqBody := RPCRequest{
		Method: method,
		Params: params,
		ID:     1,
	}
	data, _ := json.Marshal(reqBody)
	resp, err := http.Post(c.url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

func main() {
	client := NewRPCClient("http://localhost:8080/rpc")
	fmt.Println("RPC客户端已就绪")
}
