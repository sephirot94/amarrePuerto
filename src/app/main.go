package application

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	// Build Program
	app := Build()

	// Execute Program
	Execute(app)
}