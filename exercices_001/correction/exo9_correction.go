package main

import (
	"fmt"
	"math/rand"
)

func main() {
	secret := rand.Intn(100)
	var guess int

	for {
		fmt.Print("Devinez le nombre: ")
		fmt.Scan(&guess)

		if guess < secret {
			fmt.Println("Plus grand")
		} else if guess > secret {
			fmt.Println("Plus petit")
		} else {
			fmt.Println("Bravo !")
			break
		}
	}
}
