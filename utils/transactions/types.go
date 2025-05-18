package transactions

// TXOutput represents a transaction output
type TXOutput struct {
	Value      int
	PubKeyHash []byte
}

// TXInput represents a transaction input
type TXInput struct {
	Txid      []byte
	Vout      int
	Signature []byte
	PubKey    []byte
}
type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}
