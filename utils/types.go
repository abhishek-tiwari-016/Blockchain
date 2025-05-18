package utils

import (
	"main/utils/transactions"
	"math/big"

	"github.com/boltdb/bolt"
)

const TargetBits = 14

// Block represents a block in the blockchain
type Block struct {
	Timestamp     int64
	Transactions  []*transactions.Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	TIP []byte
	DB  *bolt.DB
}

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}
