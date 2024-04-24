package consumer

import "sync"

type Worker struct {
	Id     int
	Status string
	Mutex  sync.Mutex
}

func NewWorker(id int) *Worker {
	return &Worker{
		Id:     id,
		Status: "Available",
		Mutex:  sync.Mutex{},
	}
}

type Consumer struct {
	ConcurrencyFactor int
	workers           []*Worker
}

func NewConsumer(cf int) *Consumer {
	c := new(Consumer)
	c.ConcurrencyFactor = cf

	for i := 0; i < cf; i++ {
		c.workers[i] = NewWorker(i)
	}
	return c
}

func (c *Consumer) Read() {

}
