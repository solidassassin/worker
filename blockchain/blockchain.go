// NOTE: The boolean and some other returs are redundant and are not used
// They're just defined for possible future expansion

package blockchain

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

var blockCounter *uint64 = new(uint64)
var CurrentBlock Block = CreateBlock()
var blockchain Blockchain = CreateBlockchain()

func Hash(votes []Vote) []byte {

	var finalByteSlice []byte
	for _, i := range votes {
		finalByteSlice = append(finalByteSlice, i.MessageHash...)
	}
	hashed := sha256.Sum256(finalByteSlice)
	return hashed[:]
}

func CreateBlock() Block {

	block := Block{
		Index:     atomic.AddUint64(blockCounter, 1),
		Timestamp: uint64(time.Now().Unix()),
		Votes:     make([]Vote, 0, 100),
	}

	return block
}

func (block *Block) GenerateHash() []byte {
	votes := block.Votes
	// The limit was chosen mostly at random
	if len(votes) < 100 {
		log.Print("The block is not full")
		return nil
	}

	totalHash := Hash(block.Votes)
	block.Hash = totalHash

	// Not needed since it's assigned to block, but might be useful to return
	return totalHash
}

func CreateBlockchain() Blockchain {
	// Initialize empty slice
	return make([]Block, 0, 420)
}

func (vote *Vote) SignatureValid(pubKey *ecdsa.PublicKey) bool {
	isValid := ecdsa.VerifyASN1(
		pubKey,
		vote.MessageHash,
		vote.Signature,
	)

	return isValid
}

func (block *Block) AddVote(vote *Vote, pubKey *ecdsa.PublicKey) bool {
	votes := block.Votes
	if len(votes) > 99 {
	}
	if vote.SignatureValid(pubKey) {
		CurrentBlock.Votes = append(CurrentBlock.Votes, *vote)
		return true
	}
	return false
}

func (block *Block) BlockValid() bool {
	validCounter := 0
	totalCounter := 0



	validationData, err := json.Marshal(ValidationData{
        MessageHash: block.GenerateHash(),
        Signature: []byte("aaa"),
    })

	if err != nil {
		log.Fatal("Failed to encode body data")
	}

	for _, ip := range validatorNodeIps {
		resp, err := http.Post(fmt.Sprintf("%s:8000/validate", ip), "application/json", bytes.NewBuffer(validationData))

        if err != nil {
          log.Fatal("Post request failed")  
        }
		var valid IsValid
		json.NewDecoder(resp.Body).Decode(&valid)

		if err != nil {
			log.Fatal("Failed to decode body data")
		}
		if valid.Valid {
			validCounter++
		}
		totalCounter++
	}
	if float64(validCounter/totalCounter) >= 0.8 {
		CurrentBlock = CreateBlock()
		return true
	}

	return false

}

func (blch *Blockchain) AddBlock(block *Block) bool {
	if CurrentBlock.BlockValid() {
		block.PrevHash = (*blch)[len(*blch)-1].Hash
		*blch = append(*blch, *block)
		return true
	}
	return false
}
