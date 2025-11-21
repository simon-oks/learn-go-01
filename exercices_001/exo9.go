package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var (
		secretNumber = rand.Intn(100)
		userInput    int
	)

	fmt.Print("Devinez le nombre secret: ")
	fmt.Scan(&userInput)

	for userInput != secretNumber {
		if userInput < secretNumber {
			fmt.Println("Plus grand")
		} else {
			fmt.Println("Plus petit")
		}

		fmt.Print("Devinez le nombre secret: ")
		fmt.Scan(&userInput)
	}

	fmt.Println("Bravo, vous avez trouvez !")
}
