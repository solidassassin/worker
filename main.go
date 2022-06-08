package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
    "github.com/solidassassin/worker/account"
	"github.com/solidassassin/worker/network"
)

var StartTime time.Time

// Info handler defined here to avoid circular imports
func infoHandler(w http.ResponseWriter, req *http.Request) {
	info := map[string]string{
		"uptime":           time.Since(StartTime).String(),
		"key_type":         "ecdsa",  // value defined here for future use of ed25519, for now a static value
	}
	res, _ := json.Marshal(info)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(res))
}

func main() {
	StartTime = time.Now()
    account.NewWorkerId()
	http.HandleFunc("/status", network.StatusHandler)
	http.HandleFunc("/info", infoHandler)
	http.HandleFunc("/create_account", network.AccountCreator)
	http.HandleFunc("/delete_account", network.AccountDeletor)
	http.HandleFunc("/vote", network.VotingHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
