package cli

import "fmt"

func Greet(name *string) {
	fmt.Print("Welcome to the Brain Games!\nMay I have your name? ")
	fmt.Scanln(name)
	fmt.Printf("Hello, %s!\n", *name)
}
