package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/swagnikdutta/file-mq"
	producersystem "github.com/swagnikdutta/file-mq/internal/producer-system"
)

// This is the producer-system or the writer CLI. It writes to
// file.
func main() {
	// producer-system needs to have access to the file. how does it do that?
	// do we have a common file path ?
	// this is the entry point, it needs to create an instance of the producer-system.

	producer := producersystem.NewProducer(file_mq.Filepath)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "exit" {
			break
		}

		// TODO: Not sure if this a good design i.e, creating a producer thread for each message
		//  entered in STDIN. As per the problem statement, there is supposed to be just one
		//  producer thread.
		//  Should I use a channel to receive ACK or have the function return an error?
		producer.Wg.Add(1)
		go producer.Write(input)

		select {
		case err := <-producer.ErrCh:
			fmt.Println(err)
		case <-producer.AckCh:
			fmt.Println("Success")
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
