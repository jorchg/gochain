package main

import (
	"crypto/sha256"
	"encoding/hex"
	// "github.com/davecgh/go-spew/spew"
	// "sync"
	"time"
)

type Block struct {
	Index     uint64
	PrevHash  string
	PrevBlock *Block
	Timestamp int64
	Data      interface{}
	Hash      string
}

func newBlock(Data interface{}) (*Block, error) {
	readChannel := make(chan *Block, 1)
	go getLastBlock(readChannel)

	lastBlock := <-readChannel

	if lastBlock == nil {
		writeChannel := make(chan *BlockChain, 1)
		go createGenesisBlock(writeChannel)
		<-writeChannel
		go getLastBlock(readChannel)
		lastBlock = <-readChannel
	}

	Index := lastBlock.Index + 1
	PrevHash := lastBlock.Hash
	PrevBlock := lastBlock
	Timestamp := time.Now().Unix()
	hash := sha256.New()

	_, err := hash.Write([]byte(Data.(string)))
	if err != nil {
		return nil, err
	}
	Hash := hex.EncodeToString(hash.Sum(nil))

	return &Block{
		Index,
		PrevHash,
		PrevBlock,
		Timestamp,
		Data,
		Hash,
	}, nil
}
