package application

import (
	"facultad/sistemasConcurrentes/amarrePuerto/src/app/controllers"
	"facultad/sistemasConcurrentes/amarrePuerto/src/app/helpers"
	"facultad/sistemasConcurrentes/amarrePuerto/src/app/models"
	"fmt"
	"time"
)

type Application struct {
	Port             *controllers.Port
	ConcurrentHelper *helpers.ConcurrentHelper
}

func Build() *Application {
	port := controllers.CreatePort()
	cHelper := helpers.NewConcurrentHelper()

	return &Application{
		Port:             port,
		ConcurrentHelper: cHelper,
	}
}

func Execute(app *Application) {
	// Create auxiliary channels
	qtyChan := make(chan int)
	waitChan := make(chan *models.Boat)
	// Here we code the execution of the program
	// We create an autoincrement int to id the boats that are created through the execution
	var qty int = 1
	// and we send the value to the channel (channels are used between concurrent functions to communicate)
	qtyChan <- qty

	// Open a wrapper thread
	go func(*Application, chan int) {
		// Following infinite loop will wait 10 seconds and create a new boat and send that boat through the boat channel
		for {
			go func(*Application, chan int) {
				qty := <-qtyChan
				if qty >= 10 {
					time.Sleep(6 * time.Second)
				}
				boat := controllers.NewBoat(qty)
				qty++
				qtyChan <- qty
				app.ConcurrentHelper.BoatChannel <- boat
			}(app, qtyChan)
		}
	}(app, qtyChan)

	// Open a wrapper thread
	go func(*Application, chan *models.Boat) {
		// Following infinite loop will extract a boat from channel, ask if there is an available bollard for its type, and anchor it
		for {
			go func(*Application, chan *models.Boat) {
				time.Sleep(2 * time.Second)
				select {
				case waitingBoat := <-waitChan:
					// extract boat from channel
					controllers.BoatWantsIn(waitingBoat)
					bollard, id := app.Port.GetAvailableBollard(waitingBoat)
					if bollard == nil {
						// There is no available bollard at the moment
						fmt.Printf("There is no available bollard for boat named ' %s ' at the moment. Boat has to wait", waitingBoat.Name)
						waitChan <- waitingBoat
					} else {
						app.Port.UseBollard(id)
						controllers.BoatEnters(waitingBoat, id)

						// Sleep for 10 seconds then boat leaves port and bollard is available again
						time.Sleep(8 * time.Second)

						app.Port.FreeBollard(id)
						controllers.BoatLeaves(waitingBoat)
					}
				case boat := <-app.ConcurrentHelper.BoatChannel:
					// extract boat from channel
					controllers.BoatWantsIn(boat)
					bollard, id := app.Port.GetAvailableBollard(boat)
					if bollard == nil {
						// There is no available bollard at the moment
						fmt.Printf("There is no available bollard for boat named ' %s ' at the moment. Boat has to wait", boat.Name)
						waitChan <- boat
					} else {
						app.Port.UseBollard(id)
						controllers.BoatEnters(boat, id)

						// Sleep for 10 seconds then boat leaves port and bollard is available again
						time.Sleep(8 * time.Second)

						app.Port.FreeBollard(id)
						controllers.BoatLeaves(boat)
					}
				}
			}(app, waitChan)
		}
	}(app, waitChan)
}
