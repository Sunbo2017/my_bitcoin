package core

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type SBlock struct {
	Magic         int    //魔数
	Size          int64  //区块大小
	Header   	  *BlockHeader  //区块头
	TxCount       int64  //交易总数
	Transactions  []*Transaction  //交易数据
}

type BlockHeader struct {
	Version   string     //版本号
	Timestamp     int64
	PrevBlockHash []byte  //前一个区块的hash
	MerkleHash    []byte  //该区块所有交易的merkle根hash
	Hash          []byte  //经过pow计算后的该区块hash
	Nonce         int     //pow确定的随机数
	Height        int     //区块高度
}

// Block represents a block in the blockchain
// 该结构并非严格的比特币区块结构,标准结构参考SBlock
type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
	Height        int
}

// NewBlock creates and returns Block
func NewBlock(transactions []*Transaction, prevBlockHash []byte, height int) *Block {
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0, height}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{}, 0)
}

// HashTransactions returns a merkle hash of the transactions in the block
func (b *Block) HashTransactions() []byte {
	var transactions [][]byte

	for _, tx := range b.Transactions {
		transactions = append(transactions, tx.Serialize())
	}
	mTree := NewMerkleTree(transactions)

	return mTree.RootNode.Data
}

// Serialize serializes the block
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// DeserializeBlock deserializes a block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
