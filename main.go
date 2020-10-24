package main

import (
	"fmt"
	"github.com/KINTEP/golang-blockchain/blockchain"
)


func main(){

	chain := InitBlockChain()

	chain.AddBlock("First Block After Genesis")
	chain.AddBlock("Second Block After Genesis")
	chain.AddBlock("Third Block After Genesis")

	for _, block := range chain.blocks{
		//fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
	//b1 := Block{
	//	Hash: []byte{2},
	//	Data: []byte{20},
	//	PrevHash: []byte{21},
	//}

	//fmt.Println(b1)
}

