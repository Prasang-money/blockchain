package main

type BlockChain struct {
	block []*Block
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.block[len(bc.block)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.block = append(bc.block, newBlock)
}

// Function create the BlockChain instance with a genric block
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
