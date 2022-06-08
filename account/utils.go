package account

import "crypto/ecdsa"

type Account struct {
	CreationTime uint64           `json:"creationTime"`
	Name         string           `json:"name"`
	PublicKey    ecdsa.PublicKey `json:"publicKey"`
}
