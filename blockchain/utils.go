package blockchain

import "os"

var validatorNodeIps = os.Args[1:]

type IsValid struct {
	Valid bool `json:"valid"`
}

type ValidationData struct {
    MessageHash []byte `json:"messageHash"`
    Signature []byte `json:"signature"`
}

type Vote struct {
	Id          uint64 `json:"id"`
	Signature   []byte `json:"signature"`
	MessageHash []byte `json:"messageHash"`
	Message     string `json:"message"`
}

type Block struct {
	Index     uint64 `json:"index"`
	Timestamp uint64 `json:"timestamp"`
	Hash      []byte `json:"hash"`
	PrevHash  []byte `json:"prevHash"`

	Votes []Vote `json:"votes"`
}

type Blockchain []Block
