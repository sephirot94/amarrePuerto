package controllers

import (
	"facultad/sistemasConcurrentes/amarrePuerto/src/app/models"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func NewBoat(id int) *models.Boat{
	var sId string = strconv.Itoa(id)
	var size string
	var name string
	// Generate seed
	rand.Seed(time.Now().UnixNano())

	// Generate random integer between 0 and 9
	random := rand.Intn(10)

	if random <=4 {
		size = models.Small
		name = "Small ship #" + sId
	}
	if random >4 {
		size = models.Big
		name = "Big ship #" + sId
	}

	fmt.Printf("Boat named ' %s ' was created. It is a %s boat", name, size)
	return &models.Boat{
		Size: size,
		Name: name,
	}
}

func BoatWantsIn(boat *models.Boat) {
	fmt.Printf("Boat named ' %s ' is waiting to anchor", boat.Name)
}

func BoatLeaves(boat *models.Boat) {
	fmt.Printf("Boat named ' %s ' is leaving port", boat.Name)
}

func BoatEnters(boat *models.Boat, id int) {
	fmt.Printf("Boat named ' %s ' has entered port and anchored at bollard %d", boat.Name, id)
}

