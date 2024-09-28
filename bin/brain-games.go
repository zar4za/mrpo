package main

import (
	"fmt"
	cli "mrpo/src"
	"mrpo/src/games"
)

func main() {
	var name string
	cli.Greet(&name)

	fmt.Println("Let's play LCM game!")

	games.Run(name, 3, games.CreateLcmQuestion)

	fmt.Println("\nNow let's play Progression game!")
	games.Run(name, 3, games.CreateLcmQuestion)
}
