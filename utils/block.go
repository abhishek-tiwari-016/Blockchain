package utils

import (
	"bytes"
	"fmt"
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
			pow.Block.Data,
			intToHex(pow.Block.Timestamp),
			intToHex(int64(TargetBits)),
			intToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func intToHex(i int64) []byte {
	return fmt.Appendf(nil, "%x", i)
}
