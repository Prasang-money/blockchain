package main

import "fmt"

func main() {
	bc := NewBlockChain()

	bc.AddBlock("Send 1 BTC to Monu")
	bc.AddBlock("Send 2 BTC to Golu")
	bc.AddBlock("Send one more BTC to Monu")

	for _, block := range bc.block {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
