package gochain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	// "github.com/davecgh/go-spew/spew"
	// "sync"
	"time"
)

type Block struct {
	Index    uint64
	PrevHash string
	// PrevBlock *Block
	Timestamp int64
	Data      interface{}
	Hash      string
}

type DataToHash struct {
	Index     uint64
	PrevHash  string
	Timestamp int64
	Data      []byte
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
	// PrevBlock := lastBlock
	Timestamp := time.Now().Unix()
	hash := sha256.New()

	var dataBuf bytes.Buffer
	enc := gob.NewEncoder(&dataBuf)
	err := enc.Encode(Data)
	if err != nil {
		return nil, err
	}

	dataToHash := DataToHash{
		Index,
		PrevHash,
		Timestamp,
		dataBuf.Bytes(),
	}
	enc = gob.NewEncoder(&dataBuf)
	err = enc.Encode(dataToHash)
	if err != nil {
		return nil, err
	}
	_, err = hash.Write([]byte(dataBuf.Bytes()))
	if err != nil {
		return nil, err
	}
	Hash := hex.EncodeToString(hash.Sum(nil))

	return &Block{
		Index,
		PrevHash,
		// PrevBlock,
		Timestamp,
		Data,
		Hash,
	}, nil
}
