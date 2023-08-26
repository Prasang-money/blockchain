package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	TimeStamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{}
	block.TimeStamp = time.Now().Unix()
	block.Data = []byte(data)
	block.PrevBlockHash = prevBlockHash
	block.SetHash()
	return block
}
func (b *Block) SetHash() {
	timeStamp := []byte(strconv.FormatInt(b.TimeStamp, 10))
	headers := bytes.Join([][]byte{timeStamp, b.PrevBlockHash, b.Data}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
