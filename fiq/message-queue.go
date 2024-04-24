package message_queue

import (
	"os"

	"file-mq/common"
)

type MessageQueue struct {
	Filepath string
}

func NewMessageQueue(fp string) *MessageQueue {
	return &MessageQueue{
		Filepath: fp,
	}
}

func (mq *MessageQueue) Write(msg *common.Message, ackCh chan bool) {
	m := msg.CreateMessage()

	if err := os.WriteFile(mq.Filepath, []byte(m), 0666); err != nil {
		ackCh <- false
	}
	ackCh <- true
}
