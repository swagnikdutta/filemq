package main

import (
	"sync"

	message_queue "file-mq/fiq"
	"file-mq/producer"
)

const FPATH = "message-queue.txt"

func main() {
	var wg sync.WaitGroup

	mq := message_queue.NewMessageQueue(FPATH)
	producersystem := producer.NewProducer(mq)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		producersystem.Write(mq, wg)
	}(&wg)

	consumer := NewCons

	wg.Wait()
}
