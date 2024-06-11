package consumer_system

import "sync"

type Consumer struct {
	workers int
	wg      sync.WaitGroup
}

func NewConsumer(cf int) *Consumer {
	var wg sync.WaitGroup
	return &Consumer{
		workers: cf,
		wg:      wg,
	}
}

func (c *Consumer) Start() {
	// implement the worker pool pattern
	c.wg.Add(c.workers)
}
