package main

import (
	"context"
	"hashcash/internal/server"
	"log"
	"os"
)

const ADDRESS = ":8080"

func main() {
	logger := log.New(os.Stdout, "", log.Lshortfile)
	ctx := context.Background()

	err := server.Run(ctx, logger, ADDRESS)
	if err != nil {
		logger.Fatal(err)
	}
}
