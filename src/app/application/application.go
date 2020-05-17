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
	qtyChan := make(chan int, 1)
	waitChan := make(chan *models.Boat, 1)
	// Here we code the execution of the program
	// We create an autoincrement int to id the boats that are created through the execution
	var qty int = 1
	// and we send the value to the channel (channels are used between concurrent functions to communicate)
	qtyChan <- qty


		for {
			// Following infinite loop will wait 2 seconds and create a new boat and send that boat through the boat channel
			time.Sleep(2 * time.Second)
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

			// Wait 2 seconds then run second goroutine
			time.Sleep(2 * time.Second)

			// Following infinite loop will extract a boat from channel, ask if there is an available bollard for its type, and anchor it
			go func(*Application, chan *models.Boat) {
				time.Sleep(2 * time.Second)
				select {
				case boat := <- waitChan:
					switch boat.Size {
					case models.Big:
						app.ConcurrentHelper.BigWaitGroup.Wait()
					case models.Small:
						app.ConcurrentHelper.SmallWaitGroup.Wait()
					}
					controllers.BoatWantsIn(boat)
					bollard, id := app.Port.GetAvailableBollard(boat)
					if bollard == nil {
						// There is no available bollard at the moment
						fmt.Println(fmt.Sprintf("There is no available bollard for boat named ' %s ' at the moment. Boat has to wait", boat.Name))
						waitChan <- boat
					} else {
						// Add to the corresponding semaphore
						switch boat.Size {
						case models.Big:
							app.ConcurrentHelper.BigWaitGroup.Add(1)
						case models.Small:
							app.ConcurrentHelper.SmallWaitGroup.Add(1)
						}
						// Use bollard (anchor)
						app.Port.UseBollard(id)
						controllers.BoatEnters(boat, id)

						// Sleep for 10 seconds then boat leaves port and bollard is available again
						time.Sleep(5 * time.Second)

						// Leave port
						app.Port.FreeBollard(id)
						controllers.BoatLeaves(boat)
						// Signal done to the corresponding semaphore
						switch boat.Size {
						case models.Big:
							app.ConcurrentHelper.BigWaitGroup.Done()
						case models.Small:
							app.ConcurrentHelper.SmallWaitGroup.Done()
						}
					}
				case boat := <-app.ConcurrentHelper.BoatChannel:
					// extract boat from channel
					controllers.BoatWantsIn(boat)
					bollard, id := app.Port.GetAvailableBollard(boat)
					if bollard == nil {
						// There is no available bollard at the moment
						fmt.Println(fmt.Sprintf("There is no available bollard for boat named ' %s ' at the moment. Boat has to wait", boat.Name))
						waitChan <- boat
					} else {
						// Add to the corresponding semaphore
						switch boat.Size {
						case models.Big:
							app.ConcurrentHelper.BigWaitGroup.Add(1)
						case models.Small:
							app.ConcurrentHelper.SmallWaitGroup.Add(1)
						}
						// Use bollard (anchor)
						app.Port.UseBollard(id)
						controllers.BoatEnters(boat, id)

						// Sleep for 10 seconds then boat leaves port and bollard is available again
						time.Sleep(5 * time.Second)

						// Leave port
						app.Port.FreeBollard(id)
						controllers.BoatLeaves(boat)
						// Signal done to the corresponding semaphore
						switch boat.Size {
						case models.Big:
							app.ConcurrentHelper.BigWaitGroup.Done()
						case models.Small:
							app.ConcurrentHelper.SmallWaitGroup.Done()
						}
					}
				}
			}(app,waitChan)
		}
}
