package main

import (
	"context"
	"hashcash/internal/client"
	"log"
	"os"
)

const ADDRESS = "127.0.0.1:8080"

func main() {
	logger := log.New(os.Stdout, "", log.Lshortfile)
	ctx := context.Background()

	err := client.Run(ctx, logger, ADDRESS)
	if err != nil {
		logger.Fatal(err)
	}
}
