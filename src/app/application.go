package application

import (
	"facultad/sistemasConcurrentes/amarrePuerto/src/app/controllers"
	"facultad/sistemasConcurrentes/amarrePuerto/src/app/helpers"

)

type Application struct {
	Port *controllers.Port
	ConcurrentHelper *helpers.ConcurrentHelper

}

func Build() *Application {
	port := controllers.CreatePort()
	cHelper := helpers.NewConcurrentHelper()

	return &Application{
		Port: port,
		ConcurrentHelper: cHelper,
	}
}

func Execute(app *Application) {
	// Here we code the execution of the program

}

