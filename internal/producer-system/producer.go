package producer_system

import (
	"sync"

	"github.com/swagnikdutta/file-mq/internal/fiq"
)

type Producer struct {
	Wg           *sync.WaitGroup
	MessageQueue *fiq.FIQ
	AckCh        chan struct{}
	ErrCh        chan error
}

func NewProducer(fp string) *Producer {
	mq := fiq.NewFIQ(fp)
	wg := new(sync.WaitGroup)

	return &Producer{
		AckCh:        make(chan struct{}),
		ErrCh:        make(chan error),
		MessageQueue: mq,
		Wg:           wg,
	}
}

func (p *Producer) Write(message string) {
	defer p.Wg.Done()

	err := p.MessageQueue.Enqueue(message)
	if err != nil {
		p.ErrCh <- err
		return
	}
	p.AckCh <- struct{}{}
}
