package message

import (
	"encoding/json"
	"hashcash/internal/hashcash"
	"log"
)

// Message types
const (
	GetChallenge = iota
	SendSolution
)

type ClientRequest struct {
	Type     int                `json:"type"`
	Solution *hashcash.Hashcash `json:"solution,omitempty"`
}

type ServerResponse struct {
	Challenge *hashcash.Hashcash `json:"challenge,omitempty"`
	Quote     string             `json:"quote,omitempty"`
	Error     string             `json:"error,omitempty"`
}

func LogStruct(prefix string, s any) {
	encoded, err := json.Marshal(s)
	if err != nil {
		log.Println("can't encode struct", s)
	}
	log.Printf("%s %s", prefix, encoded)
}
