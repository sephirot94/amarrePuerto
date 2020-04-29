package helpers

import (
	"sync"
)

type ConcurrentInterface interface {
	Add()
	Wait()
	Done()
}

type ConcurrentHelper struct {
	WaitGroup *sync.WaitGroup
}

func NewConcurrentHelper() ConcurrentInterface{
	var wg sync.WaitGroup
	return &ConcurrentHelper{WaitGroup: &wg}
}

func (h ConcurrentHelper) Add() {
	h.WaitGroup.Add(1)
}

func (h ConcurrentHelper) Wait() {
	h.WaitGroup.Wait()
}

func (h ConcurrentHelper) Done() {
	h.WaitGroup.Done()
}