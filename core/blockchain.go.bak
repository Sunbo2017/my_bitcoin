package core

type BlockChain struct {
	Blocks []*Block
}

func (bc *BlockChain) AddBlock(data string) {
	//获取当前最后一个区块
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	//关联hash值
	block := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, block)
}

func NewBlockChain() *BlockChain {
	return &BlockChain{Blocks: []*Block{NewGenesisBlock()}}
}
