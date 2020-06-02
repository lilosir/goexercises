package core

import (
	"fmt"
)

// BlockChain defination
type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	genesisBlock := CreateGenesisBlock()
	bc := BlockChain{}
	bc.AppendBlock(&genesisBlock)
	return &bc
}

func (bc *BlockChain) SendData(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := CreateNewBlock(*prevBlock, data)
	bc.AppendBlock(&newBlock)
}

func (bc *BlockChain) AppendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		fmt.Println("invalid block")
	}
}

func isValid(newBlock, oldBlock Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		return false
	}
	if newBlock.PrevBlockHash != oldBlock.CurrBlockHash {
		return false
	}
	if GetHash(newBlock) != newBlock.CurrBlockHash {
		return false
	}
	return true
}

func (bc *BlockChain) Print() {
	for _, b := range bc.Blocks {
		fmt.Printf("Index: %d\n", b.Index)
		fmt.Printf("Previous Hash: %s\n", b.PrevBlockHash)
		fmt.Printf("Current Hash: %s\n", b.CurrBlockHash)
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Timestamp: %d\n", b.Timestamp)
		fmt.Println()
	}
}
