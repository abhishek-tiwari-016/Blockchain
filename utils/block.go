package utils

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"main/utils/transactions"
	"time"
)

// func (b *utils.Block) SetHash() {
// 	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
// 	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
// 	hash := sha256.Sum256(headers)

//		b.Hash = hash[:]
//	}
func (pow *ProofOfWork) PrepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.HashTransactions(),
			intToHex(pow.Block.Timestamp),
			intToHex(int64(TargetBits)),
			intToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func NewBlock(transactions []*transactions.Transaction, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func intToHex(i int64) []byte {
	return fmt.Appendf(nil, "%x", i)
}

func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.Hash())
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}
