package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	// "log"
	"sync"
)

type BlockChain struct {
	Blocks []*Block
}

var instance *BlockChain
var once sync.Once

func getInstance() *BlockChain {
	once.Do(func() {
		instance = &BlockChain{}
	})
	return instance
}

func getLastBlock(readChannel chan *Block) {
	blockChain := getInstance()
	// TODO: defer
	if len(blockChain.Blocks) == 0 {
		readChannel <- nil
	}

	readChannel <- blockChain.Blocks[len(blockChain.Blocks)-1]
}

func createGenesisBlock(writeChannel chan *BlockChain) {
	blockChain := getInstance()

	if len(blockChain.Blocks) != 0 {
		writeChannel <- blockChain
		return
	}

	block := Block{Data: "TEST"}
	blockChain.Blocks = append(blockChain.Blocks, &block)
	fmt.Println("CREATING GENESIS BLOCK: ", block)
	spew.Dump("BLOCKCHAIN: ", blockChain.Blocks)
	writeChannel <- blockChain
	return
}

func mineBlock(block *Block, writeChannel chan *BlockChain) {
	blockChain := getInstance()
	// TODO: defer
	if len(blockChain.Blocks) == 0 {
		createGenesisBlock(writeChannel)
		return
	}

	blockChain.Blocks = append(blockChain.Blocks, block)
	writeChannel <- blockChain
}
