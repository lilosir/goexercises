package main

import (
	"demochain/core"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var blockchain *core.BlockChain

func run() {
	http.HandleFunc("/blockchain/get", blockchainGetHandlder)
	http.HandleFunc("/blockchain/write", blockchainWriteHandlder)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server, err: %s\n", err.Error())
		panic(err)
	}

	println("Running code after ListenAndServe (only happens when server shuts down)")
}

func blockchainGetHandlder(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(blockchain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func blockchainWriteHandlder(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	blockchain.SendData(blockData)
	blockchainGetHandlder(w, r)
}

func main() {
	blockchain = core.NewBlockChain()
	run()
}