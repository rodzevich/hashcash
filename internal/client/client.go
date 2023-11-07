package client

import (
	"context"
	"encoding/json"
	"hashcash/internal/message"
	"log"
	"net"
	"time"
)

func Run(ctx context.Context, log *log.Logger, address string) error {
	log.Printf("connecting to %s", address)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	for {
		time.Sleep(10 * time.Second)

		msg := message.ClientRequest{
			Type: message.GetChallenge,
		}
		log.Printf("requesting challenge")
		resp, err := sendRequest(conn, &msg)
		if err != nil {
			log.Printf("error %v", err)
			continue
		}
		message.LogStruct("response", resp)
		task, err := resp.Challenge.Solve()
		if err != nil {
			log.Printf("can't solve %v", err)
			continue
		}
		msg = message.ClientRequest{
			Type:     message.SendSolution,
			Solution: &task,
		}
		message.LogStruct("sending solution", msg)
		resp, err = sendRequest(conn, &msg)
		if err != nil {
			log.Printf("error %v", err)
			continue
		}
		message.LogStruct("response", resp)
	}
}

func sendRequest(conn net.Conn, msg *message.ClientRequest) (*message.ServerResponse, error) {
	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)
	err := encoder.Encode(msg)
	if err != nil {
		return nil, err
	}
	var resp message.ServerResponse
	err = decoder.Decode(&resp)
	return &resp, err
}
