package fiq

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type FIQ struct {
	mu   sync.Mutex
	file *os.File
}

func NewFIQ(fp string) *FIQ {
	f, err := os.OpenFile(fp, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Error opening file: %s", err.Error())
	}

	return &FIQ{file: f}
}

func (f *FIQ) Enqueue(m string) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	_, err := f.file.WriteString(m + "\n")
	if err != nil {
		return fmt.Errorf("error enqueueing: %s", err.Error())
	}
	return err
}

func (f *FIQ) Dequeue() (string, error) {
	// this is where the offset logic comes in
	b := make([]byte, 1024)
	_, err := f.file.Read(b)
	if err != nil {
		return "", fmt.Errorf("error dequeueing: %s", err.Error())
	}

	return string(b), nil
}
