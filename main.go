package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Timestamp int64
	Data      string
	PrevHash  string
	Hash      string
}

// NewBlock creates and returns a new block.
func NewBlock(data string, prevHash string) *Block {
	block := &Block{
		Timestamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevHash,
	}
	block.Hash = calculateHash(block)
	return block
}

// calculateHash generates a hash from a block's data.
func calculateHash(block *Block) string {
	record := fmt.Sprintf("%d%s%s", block.Timestamp, block.Data, block.PrevHash)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// DisplayAllBlocks prints all blocks in the blockchain to the console.
func DisplayAllBlocks(blocks []*Block) {
	for _, block := range blocks {
		fmt.Printf("Timestamp: %d, Data: %s, PrevHash: %s, Hash: %s\n",
			block.Timestamp, block.Data, block.PrevHash, block.Hash)
	}
}

// ModifyBlock changes the data of a block and updates its hash.
func ModifyBlock(block *Block, newData string) {
	block.Data = newData
	block.Hash = calculateHash(block)
}

func main() {
	genesisBlock := NewBlock("Genesis Block", "")
	blocks := []*Block{genesisBlock}

	// Adding new blocks
	blocks = append(blocks, NewBlock("First Block", genesisBlock.Hash))
	blocks = append(blocks, NewBlock("Second Block", blocks[len(blocks)-1].Hash))

	// Display all blocks before modification
	fmt.Println("Initial blockchain:")
	DisplayAllBlocks(blocks)

	// Modify a block
	fmt.Println("\nModifying the second block...")
	ModifyBlock(blocks[1], "Modified Second Block")

	// Display all blocks after modification
	fmt.Println("\nBlockchain after modification:")
	DisplayAllBlocks(blocks)
}
