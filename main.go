package main

import (
	"main/utils/cli"
)

func main() {
	// bc := utils.NewBlockchain()

	// bc.AddBlock("Send 1 BTC to Ivan")
	// bc.AddBlock("Send 2 more BTC to Ivan")

	// for _, block := range bc.Blocks {
	// 	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	// 	fmt.Printf("Data: %s\n", block.Data)
	// 	fmt.Printf("Hash: %x\n", block.Hash)
	// 	fmt.Println()
	// 	pow := utils.NewProofOfWork(block)
	// 	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	// 	fmt.Println()
	// }

	// bc := utils.NewBlockchain()
	// defer bc.DB.Close()

	cli := cli.CLI{}
	cli.Run()
}
