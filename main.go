package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

//A function to create the hash of a block using its content
//It populates the Hash feild of a block.
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

//CreateBlock returns a pointer to a block
func CreateBlock(Data string, PrevHash []byte) *Block {
	var block Block
	block.Data = []byte(Data)
	block.PrevHash = PrevHash
	//populates the hash field
	block.DeriveHash()

	return &block
}

//Adds a created block to the chain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

//Creates the genesis block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

//initializes the block chain with a genesis block
func InitializeChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
func main() {
	chain := InitializeChain()

	chain.AddBlock("First block after genesis")
	chain.AddBlock("Secon block after genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("data :%s\n",block.Data)
		fmt.Printf("Hash :%x\n ",block.Hash)
	}
}
