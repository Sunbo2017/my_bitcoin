package main

import (
	"my_bitcoin/core"
)

func main() {
	bc := core.NewBlockChain()
	defer bc.DB.Close()

	cli := core.CLI{BC: bc}
	cli.Run()
}
