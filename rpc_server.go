package main

import (
	"encoding/json"
	"net/http"
)

type RPCServer struct {
	port string
}

func NewRPCServer(port string) *RPCServer {
	return &RPCServer{port: port}
}

func (s *RPCServer) handler(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	json.NewDecoder(r.Body).Decode(&req)
	resp := map[string]interface{}{
		"id":     req["id"],
		"result": "rpc response",
	}
	json.NewEncoder(w).Encode(resp)
}

func (s *RPCServer) Start() error {
	http.HandleFunc("/rpc", s.handler)
	return http.ListenAndServe(":"+s.port, nil)
}

func main() {
	server := NewRPCServer("8080")
	server.Start()
}
