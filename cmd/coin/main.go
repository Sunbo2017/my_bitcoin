package main

import (
	"fmt"
	"my_bitcoin/core"
)

func main() {
	bc := core.NewBlockChain()
	bc.AddBlock("send 1 btc to sun")
	bc.AddBlock("send 2 btc to sun")

	for _, b := range bc.Blocks {
		fmt.Printf("Prev.hash:%x\n", b.PrevBlockHash)
		fmt.Printf("Data:%s\n", b.Data)
		fmt.Printf("Hash:%x\n", b.Hash)
		fmt.Println("---------------")
	}
}
