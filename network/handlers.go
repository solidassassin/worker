package network

import (
	"encoding/json"
	"fmt"
	"github.com/solidassassin/worker/account"
    "github.com/solidassassin/worker/blockchain"
	"net/http"
)

func AccountCreator(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "Only POST requets are allowed", 405)
		return
	}
	var nameData NameData
	err := json.NewDecoder(req.Body).Decode(&nameData)

	if err != nil {
		http.Error(w, "Failed to decode body data", 500)
	}

	_, privKey := account.NewAccount(nameData.Name)

	res, _ := json.Marshal(PrivKeyRes{Key: privKey})

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(res))
}

func AccountDeletor(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "Only POST requets are allowed", 405)
		return
	}
	var privKey PrivKeyRes
	err := json.NewDecoder(req.Body).Decode(&privKey)

	if err != nil {
		http.Error(w, "Failed to decode body data", 500)
	}

	if account.DeleteAccount(privKey.Key) {
		fmt.Fprint(w, "ok")
	} else {
		fmt.Fprint(w, "failed")
	}
}

func VotingHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "Only POST requets are allowed", 405)
		return
	}
    if len(blockchain.CurrentBlock.Votes) > 99 {
        blockchain.CreateBlock()
    }

    blockchain.CurrentBlock.AddVote()
}

func StatusHandler(w http.ResponseWriter, req *http.Request) {
	// Just a healthcheck endpoint
	fmt.Fprint(w, "ok")
}


// DEAD CODE
func InfoHandler(w http.ResponseWriter, req *http.Request) {
	info := map[string]string{
		"key_type":         "a",
		"validation_count": "a",
	}
	res, _ := json.Marshal(info)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(res))
}
