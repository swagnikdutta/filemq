package producer

import (
	"fmt"
	"sync"

	"file-mq/common"
	"file-mq/fiq"
)

type Producer struct {
	MQ       *message_queue.MessageQueue
	Messages []common.Message
	ackCh    chan bool
}

func NewProducer(q *message_queue.MessageQueue) *Producer {
	p := new(Producer)
	p.MQ = q
	p.Messages = make([]common.Message, 0)
	messages := []common.Message{
		{Key: "m1", Message: "message 1", ProcessingTime: 2},
		{Key: "m2", Message: "message 2", ProcessingTime: 4},
		{Key: "m3", Message: "message 3", ProcessingTime: 1},
		{Key: "m4", Message: "message 4", ProcessingTime: 6},
	}
	p.Messages = append(p.Messages, messages...)
	return p
}

func (p *Producer) Write(mq *message_queue.MessageQueue, wg *sync.WaitGroup) {
	defer wg.Done()

	n := len(p.Messages)

	for i := 0; i < n; i++ {
		m := &p.Messages[i]

		mq.Write(m, p.ackCh)

		verdict := <-p.ackCh

		switch verdict {
		case true:
			fmt.Println("Writing success")
		case false:
			fmt.Println("Writing failed")
		}
	}
}
