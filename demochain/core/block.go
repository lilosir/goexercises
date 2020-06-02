package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Block defination
type Block struct {
	Index         int64  //block index
	Timestamp     int64  //block timestamp
	PrevBlockHash string // previous block hash value
	CurrBlockHash string // current block hash value
	Data          string // block data
}

func GetHash(b Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Data
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

func CreateNewBlock(prev Block, data string) Block {
	newB := Block{
		Data: data,
	}
	newB.Index = prev.Index + 1
	newB.Timestamp = time.Now().Unix()
	newB.PrevBlockHash = prev.CurrBlockHash
	newB.CurrBlockHash = GetHash(newB)
	return newB
}

func CreateGenesisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.CurrBlockHash = ""
	return CreateNewBlock(preBlock, "genesis block")
}
