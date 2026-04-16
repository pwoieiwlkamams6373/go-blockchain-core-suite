package main

import (
	"encoding/json"
	"net/http"
)

type NodeAPI struct {
	port string
}

func NewNodeAPI(port string) *NodeAPI {
	return &NodeAPI{port: port}
}

func (api *NodeAPI) getBlockHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"height": 100,
		"status": "active",
	})
}

func (api *NodeAPI) Start() error {
	http.HandleFunc("/api/block", api.getBlockHandler)
	return http.ListenAndServe(":"+api.port, nil)
}

func main() {
	api := NewNodeAPI("8080")
	api.Start()
}
