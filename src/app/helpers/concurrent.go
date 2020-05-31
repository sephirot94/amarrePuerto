package helpers

import (
	"facultad/sistemasConcurrentes/amarrePuerto/src/app/models"
	"sync"
)

type ConcurrentInterface interface {

}

type ConcurrentHelper struct {
	// WaitGroup = Semaphore in Golang
	SmallWaitGroup *sync.WaitGroup
	BigWaitGroup *sync.WaitGroup
	// Channels are used by goroutines to communicate
	BoatChannel chan *models.Boat
	BollardChannel chan *models.Bollard
	ErrorChannel chan error
}

func NewConcurrentHelper() *ConcurrentHelper{
	errChan := make(chan error)
	boatChan := make(chan *models.Boat, 10)
	bollardChan := make(chan *models.Bollard, 10)
	var bwg sync.WaitGroup
	var swg sync.WaitGroup
	return &ConcurrentHelper{
		SmallWaitGroup: &swg,
		BigWaitGroup: &bwg,
		BoatChannel:boatChan,
		BollardChannel: bollardChan,
		ErrorChannel: errChan,
	}
}
