package helpers

import (
	"facultad/sistemasConcurrentes/amarrePuerto/src/app/models"
	"sync"
)

type ConcurrentInterface interface {
	Add()
	Wait()
	Done()
}

type ConcurrentHelper struct {
	WaitGroup *sync.WaitGroup
	BoatChannel chan *models.Boat
	BollardChannel chan *models.Bollard
	ErrorChannel chan error
}

func NewConcurrentHelper() *ConcurrentHelper{
	errChan := make(chan error)
	boatChan := make(chan *models.Boat)
	bollardChan := make(chan *models.Bollard)
	var wg sync.WaitGroup
	return &ConcurrentHelper{
		WaitGroup: &wg,
		BoatChannel:boatChan,
		BollardChannel: bollardChan,
		ErrorChannel: errChan,
	}
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