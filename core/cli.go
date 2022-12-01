package core

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type CLI struct {
	BC *BlockChain
}

func (c *CLI) printUsage()  {
	fmt.Println("Usage:")
	fmt.Println(" addblock -data BLOCK_DATA - add new block")
	fmt.Println(" printchain - print all the blocks of the blockChain")
}

func (c *CLI) validateArgs() {
	if len(os.Args) < 2 {
		c.printUsage()
		os.Exit(1)
	}
}

func (c *CLI) addBlock(data string) {
	c.BC.AddBlock(data)
	fmt.Println("Success!")
}

func (c *CLI) printChain() {
	bci := c.BC.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

func (c *CLI) Run() {
	c.validateArgs()
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlock := addBlockCmd.String("data", "", "block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		c.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlock == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		c.addBlock(*addBlock)
	}

	if printChainCmd.Parsed() {
		c.printChain()
	}
}

