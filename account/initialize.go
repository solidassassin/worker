package account

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"
	"time"
)

var accounts []Account
var PrivateKey *ecdsa.PrivateKey

// The probability of this matching is basically impossible
// Better safe than sorry tho
func accountExists(account Account) bool {
	for _, i := range accounts {
		if i.PublicKey == account.PublicKey {
			return true
		}
	}
	return false
}

func NewAccount(username string) (Account, *ecdsa.PrivateKey) {

	for {
		privKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
		if err != nil {
			log.Fatal("Failed to generate key pair")
		}

		acc := Account{
			CreationTime: uint64(time.Now().Unix()),
			Name:         username,
			// Server will not store the private key and only sends it to the user
			// If the user loses the key, their account is as good as gone
			PublicKey: privKey.PublicKey,
		}
		if !accountExists(acc) {
			accounts = append(accounts, acc)
			return acc, privKey
		}
	}
}

func NewWorkerId() *ecdsa.PublicKey {
	privKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		log.Fatal("Failed to generate key pair")
	}
    PrivateKey = privKey
    return &privKey.PublicKey
}

func DeleteAccount(privKey *ecdsa.PrivateKey) bool {

	for i, x := range accounts {
		if x.PublicKey == privKey.PublicKey {
			accounts = append(accounts[:i], accounts[i+1:]...)
			return true
		}
	}
	return false

}

// NOTE: DEAD CODE BELLOW
// This is just an example of how signing would look on the server
type Account2 struct {
	CreationTime uint64            `json:"creationTime"`
	Name         string            `json:"name"`
	Keys         *ecdsa.PrivateKey `json:"privateKey"`
}

func accountExists2(accounts2 []Account2, account Account2) bool {
	for _, i := range accounts2 {
		if i.Keys.PublicKey == account.Keys.PublicKey {
			return true
		}
	}
	return false
}
func NewAccount2(username string, accounts []Account2) Account2 {

	for {
		privKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
		if err != nil {
			log.Fatal("Failed to generate key pair")
		}

		acc := Account2{
			CreationTime: uint64(time.Now().Unix()),
			Name:         username,
			// Server will not store the private key and only sends it to the user
			// If the user loses the key, their account is as good as gone
			Keys: privKey,
		}
		if !accountExists2(accounts, acc) {
			return acc
		}
	}
}
func hash(message string) []byte {
	hashed := sha256.Sum256([]byte(message))
	return hashed[:]
}

// Signing should be handled by the end user due to security concerns, this is an example
func (account *Account2) SignVote(message string) []byte {
	hashedMsg := hash(message)
	signature, _ := ecdsa.SignASN1(rand.Reader, account.Keys, hashedMsg)
	return signature
}
