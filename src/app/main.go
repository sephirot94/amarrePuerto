package main

import (
	"facultad/sistemasConcurrentes/amarrePuerto/src/app/application"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("hello")

	// Build Program
	app := application.Build()

	// Execute Program
	application.Execute(app)

	//if err != nil {
	//	fmt.Printf("Unexpected error: %s", err.Error())
	//}

	// Following code will close program after a minute
	time.Sleep(1 * time.Minute)
	fmt.Println("That was enough for today! We are gonna stop working or else this computer might crash. Thank you for watching us work")
	os.Exit(0)
}
