package common

import (
	"fmt"
	"time"
)

type Message struct {
	Key            string
	Message        string
	ProcessingTime time.Duration
}

func (m *Message) CreateMessage() string {
	return fmt.Sprintf("%s:%s:%s\n", m.Key, m.Message, m.ProcessingTime)
}
