package network

import (
	"crypto/ecdsa"
	"github.com/solidassassin/worker/blockchain"
	"os"
)

var validatorIPList = os.Args[1:]

type ValidationData struct {
	PublicKey   *ecdsa.PublicKey `json:"public_key"`
	MessageHash []byte           `json:"message_hash"`
	Signature   []byte           `json:"signature"`
}

type BlockchainHandler struct {
	blch blockchain.Blockchain
}

type BlockHandler struct {
	block blockchain.Block
}

type NameData struct {
	Name string `json:"name"`
}

type PrivKeyRes struct {
	Key *ecdsa.PrivateKey `json:"keyPair"`
}

