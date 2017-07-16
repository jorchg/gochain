package main

import (
	"flag"
	"github.com/davecgh/go-spew/spew"
	// "github.com/gorilla/websocket"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Message struct {
	Data interface{}
}

type ArbitraryMessage interface{}

var addr = flag.String("addr", "localhost:3000", "http service address")

func main() {
	http.Handle("/getPeers", http.HandlerFunc(handleRouteGetPeers))
	http.Handle("/mineBlock", http.HandlerFunc(handleRouteMineBlock))

	log.Fatal(http.ListenAndServe(*addr, nil))
}

func handleRouteGetPeers(res http.ResponseWriter, req *http.Request) {
	log.Println("Get Peers")
	res.WriteHeader(http.StatusOK)
	io.WriteString(res, "Yeah!!")
}

func handleRouteMineBlock(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, "mineBlock route method must be POST", http.StatusBadRequest)
		log.Println("Wrong Method")
		return
	}

	// Read body
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg Message
	// var arbitrary ArbitraryMessage
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	writeChannel := make(chan *BlockChain, 1)
	newBlock, _ := newBlock(msg.Data)
	go mineBlock(newBlock, writeChannel)
	blockChain := <-writeChannel
	spew.Dump("BLOCKCHAIN: ", blockChain.Blocks)

	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	res.Header().Set("content-type", "application/json")
	res.Write(output)
}
