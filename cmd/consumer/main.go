package main

import (
	"log"
	"os"
	"strconv"

	consumersystem "github.com/swagnikdutta/file-mq/internal/consumer-system"
)

// This is the consumer-system(reader) system. The consumers
// will read a message and process them. There can be
// multiple consumer-system threads.
func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage:\nconsumer <concurrency_factor>")
	}

	cf, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Error converting string to integer: %s", err.Error())
	}

	consumer := consumersystem.NewConsumer(cf)
	consumer.Start()
}
