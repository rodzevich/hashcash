package server

import (
	"context"
	"encoding/json"
	memcache "github.com/patrickmn/go-cache"
	"hashcash/internal/hashcash"
	"hashcash/internal/message"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

const (
	cacheTime = 24 * time.Hour
)

var quotes = []string{
	"When the going gets rough - turn to wonder.",
	"If you have knowledge, let others light their candles in it.",
	"A bird doesn't sing because it has an answer, it sings because it has a song.",
	"We are not what we know but what we are willing to learn.",
}

func Run(ctx context.Context, log *log.Logger, address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer listener.Close()

	cache := memcache.New(cacheTime, memcache.NoExpiration)
	log.Printf("serving %s", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("client connection error: %v", err)
			break
		}
		go handleConnection(ctx, log, conn, cache)
	}
	return nil
}

func handleConnection(ctx context.Context, log *log.Logger, conn net.Conn, cache *memcache.Cache) {
	//TODO: check context for graceful shutdown
	defer conn.Close()
	log.Printf("%s connected", conn.RemoteAddr())

	decoder := json.NewDecoder(conn)

	var msg message.ClientRequest

	for {
		err := decoder.Decode(&msg)
		if err != nil {
			msg := "invalid request"
			sendResponse(log, conn, message.ServerResponse{Error: msg})
			return
		}

		message.LogStruct("request", msg)
		resp, err := processMessage(&msg, cache)
		if err != nil {
			msg := "can't process message"
			sendResponse(log, conn, message.ServerResponse{Error: msg})
			return
		}

		sendResponse(log, conn, resp)
	}
}

func processMessage(msg *message.ClientRequest, cache *memcache.Cache) (message.ServerResponse, error) {
	switch msg.Type {
	case message.GetChallenge:
		task := hashcash.New()
		resp := message.ServerResponse{
			Challenge: &task,
		}
		cache.Set(strconv.Itoa(task.Rand), 1, memcache.DefaultExpiration)
		return resp, nil
	case message.SendSolution:
		_, found := cache.Get(strconv.Itoa(msg.Solution.Rand))
		if !found {
			return message.ServerResponse{Error: "task not found"}, nil
		}
		if !msg.Solution.IsValid() {
			return message.ServerResponse{Error: "invalid solution"}, nil
		}
		resp := message.ServerResponse{
			Quote: quotes[rand.Intn(len(quotes))],
		}
		return resp, nil
	}

	return message.ServerResponse{Error: "unknown request"}, nil
}

func sendResponse(log *log.Logger, conn net.Conn, resp message.ServerResponse) {
	encoder := json.NewEncoder(conn)
	err := encoder.Encode(resp)
	if err != nil {
		log.Printf("can't send response %v", err)
		return
	}

	message.LogStruct("response", resp)
}
