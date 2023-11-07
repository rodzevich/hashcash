package server

import (
	"context"
	"encoding/json"
	memcache "github.com/patrickmn/go-cache"
	"hashcash/internal/message"
	"io"
	"log"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	ctx := context.Background()
	loggerMock := log.New(io.Discard, "", log.LstdFlags)
	cache := memcache.New(memcache.NoExpiration, memcache.NoExpiration)

	server, client := net.Pipe()
	defer server.Close()
	defer client.Close()

	go handleConnection(ctx, loggerMock, server, cache)

	buf := make([]byte, 1024)
	_, err := client.Write([]byte("{\"type\": 0}"))
	if err != nil {
		t.Fatal(err)
	}
	size, err := client.Read(buf)
	if err != nil {
		t.Fatal(err)
	}

	response := message.ServerResponse{}
	err = json.Unmarshal(buf[:size], &response)
	if err != nil {
		t.Fatal(err)
	}

	solution, err := response.Challenge.Solve()
	if err != nil {
		t.Fatal(err)
	}

	request := message.ClientRequest{
		Type:     message.SendSolution,
		Solution: &solution,
	}
	buf, err = json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Write(buf)
	if err != nil {
		t.Fatal(err)
	}
	size, err = client.Read(buf)
	if err != nil {
		t.Fatal(err)
	}
}
