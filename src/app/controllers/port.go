package controllers

import (
	"facultad/sistemasConcurrentes/amarrePuerto/src/app/models"
)

type PortInterface interface {
	CheckBollard(id int) bool
	UseBollard(id int)
	FreeBollard(id int)
	GetAvailableBollard(boat *models.Boat) *models.Bollard
}

type Port struct {
	Map map[int]*models.Bollard
}

func CreatePort() *Port {

	m := make(map[int]*models.Bollard)
	// Create Big Bollard
	for i := 1; i <= 4; i++ {
		bollard := NewBollard(models.Big, i)
		m[i] = bollard
	}

	// Create Small Bollard
	for i := 5; i <= 10; i++ {
		bollard := NewBollard(models.Small, i)
		m[i] = bollard
	}

	return &Port{
		Map: m,
	}

}

func NewBollard(t string, n int) *models.Bollard {
	return &models.Bollard{
		Size: t,
		Free: true,
	}
}

func (p Port) CheckBollard(id int) bool {
	return p.Map[id].Free
}

func (p Port) UseBollard(id int) {
	p.Map[id].Free = false
}

func (p Port) FreeBollard(id int) {
	p.Map[id].Free = true
}

func (p Port) GetAvailableBollard(boat *models.Boat) (*models.Bollard, int) {
	if boat.Size == models.Big {
		for i := 1; i <= 4; i++ {
			if p.CheckBollard(i) {
				return p.Map[i], i
			}
		}
	} else {
		for i := 5; i <= 10; i++ {
			if p.CheckBollard(i) {
				return p.Map[i], i
			}
		}
	}
	return nil, 0
}
