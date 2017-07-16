package main

type Block struct {
	// Index     uint64
	// PrevHash  string
	// PrevBlock *Block
	// Timestamp int64
	Data string
	// Hash      string
}

func newBlock(Index uint64, PrevHash string, PrevBlock *Block, Timestamp int64, Data, Hash string) *Block {
	return &Block{
		// Index,
		// PrevHash,
		// PrevBlock,
		// Timestamp,
		Data,
		// Hash,
	}
}
