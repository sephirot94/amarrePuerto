package main

import (
	"facultad/sistemasConcurrentes/amarrePuerto/src/app/application"
	"fmt"
	"os"
	"time"
)

func main() {

	// Build Program
	app := application.Build()
	fmt.Println("Port system activating")
	// Execute Program
	go func(*application.Application){
		application.Execute(app)
	}(app)

	// Following code will close program after a minute
	time.Sleep(60 * time.Second)
	fmt.Println("That was enough for today! We are gonna stop working or else this computer might crash. Thank you for watching")
	os.Exit(0)
}
