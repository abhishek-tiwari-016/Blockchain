package utils

import "math/big"

const TargetBits = 14

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

type Blockchain struct {
	Blocks []*Block
}

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}
